package shared

import (
	"crypto/ed25519"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"sync"
)

type SharedServiceContainer struct {
	eth *eth.Ethereum
	bc  *core.BlockChain

	commitmentUpdateLock sync.Mutex // Lock for the forkChoiceUpdated method
	blockExecutionLock   sync.Mutex // Lock for the NewPayload method

	genesisInfoCalled        bool
	getCommitmentStateCalled bool

	bridgeAddresses     map[string]*params.AstriaBridgeAddressConfig // astria bridge addess to config for that bridge account
	bridgeAllowedAssets map[string]struct{}                          // a set of allowed asset IDs structs are left empty

	trustedBuilderPublicKey ed25519.PublicKey

	// TODO: bharath - we could make this an atomic pointer???
	nextFeeRecipient common.Address // Fee recipient for the next block
}

func NewSharedServiceContainer(eth *eth.Ethereum) (*SharedServiceContainer, error) {
	bc := eth.BlockChain()

	if bc.Config().AstriaRollupName == "" {
		return nil, errors.New("rollup name not set")
	}

	if bc.Config().AstriaSequencerInitialHeight == 0 {
		return nil, errors.New("sequencer initial height not set")
	}

	if bc.Config().AstriaCelestiaInitialHeight == 0 {
		return nil, errors.New("celestia initial height not set")
	}

	if bc.Config().AstriaCelestiaHeightVariance == 0 {
		return nil, errors.New("celestia height variance not set")
	}

	bridgeAddresses := make(map[string]*params.AstriaBridgeAddressConfig)
	bridgeAllowedAssets := make(map[string]struct{})
	if bc.Config().AstriaBridgeAddressConfigs == nil {
		log.Warn("bridge addresses not set")
	} else {
		nativeBridgeSeen := false
		for _, cfg := range bc.Config().AstriaBridgeAddressConfigs {
			err := cfg.Validate(bc.Config().AstriaSequencerAddressPrefix)
			if err != nil {
				return nil, fmt.Errorf("invalid bridge address config: %w", err)
			}

			if cfg.Erc20Asset == nil {
				if nativeBridgeSeen {
					return nil, errors.New("only one native bridge address is allowed")
				}
				nativeBridgeSeen = true
			}

			if cfg.Erc20Asset != nil && cfg.SenderAddress == (common.Address{}) {
				return nil, errors.New("astria bridge sender address must be set for bridged ERC20 assets")
			}

			bridgeCfg := cfg
			bridgeAddresses[cfg.BridgeAddress] = &bridgeCfg
			bridgeAllowedAssets[cfg.AssetDenom] = struct{}{}
			if cfg.Erc20Asset == nil {
				log.Info("bridge for sequencer native asset initialized", "bridgeAddress", cfg.BridgeAddress, "assetDenom", cfg.AssetDenom)
			} else {
				log.Info("bridge for ERC20 asset initialized", "bridgeAddress", cfg.BridgeAddress, "assetDenom", cfg.AssetDenom, "contractAddress", cfg.Erc20Asset.ContractAddress)
			}
		}
	}

	// To decrease compute cost, we identify the next fee recipient at the start
	// and update it as we execute blocks.
	nextFeeRecipient := common.Address{}
	if bc.Config().AstriaFeeCollectors == nil {
		log.Warn("fee asset collectors not set, assets will be burned")
	} else {
		maxHeightCollectorMatch := uint32(0)
		nextBlock := uint32(bc.CurrentBlock().Number.Int64()) + 1
		for height, collector := range bc.Config().AstriaFeeCollectors {
			if height <= nextBlock && height > maxHeightCollectorMatch {
				maxHeightCollectorMatch = height
				nextFeeRecipient = collector
			}
		}
	}

	// TODO - is it desirable to not fail if the trusted builder public key is not set?
	if bc.Config().AstriaTrustedBuilderPublicKey == "" {
		return nil, errors.New("trusted builder public key not set")
	}
	// validate if its an ed25519 public key
	if len(bc.Config().AstriaTrustedBuilderPublicKey) != ed25519.PublicKeySize {
		return nil, errors.New("trusted builder public key is not a valid ed25519 public key")
	}

	sharedServiceContainer := &SharedServiceContainer{
		eth:                     eth,
		bc:                      bc,
		bridgeAddresses:         bridgeAddresses,
		bridgeAllowedAssets:     bridgeAllowedAssets,
		nextFeeRecipient:        nextFeeRecipient,
		trustedBuilderPublicKey: ed25519.PublicKey(bc.Config().AstriaTrustedBuilderPublicKey),
	}

	return sharedServiceContainer, nil
}

func (s *SharedServiceContainer) SyncMethodsCalled() bool {
	return s.genesisInfoCalled && s.getCommitmentStateCalled
}

func (s *SharedServiceContainer) Bc() *core.BlockChain {
	return s.bc
}

func (s *SharedServiceContainer) Eth() *eth.Ethereum {
	return s.eth
}

func (s *SharedServiceContainer) SetGenesisInfoCalled(value bool) {
	s.genesisInfoCalled = value
}

func (s *SharedServiceContainer) GenesisInfoCalled() bool {
	return s.genesisInfoCalled
}

func (s *SharedServiceContainer) SetGetCommitmentStateCalled(value bool) {
	s.getCommitmentStateCalled = value
}

func (s *SharedServiceContainer) CommitmentStateCalled() bool {
	return s.getCommitmentStateCalled
}

func (s *SharedServiceContainer) CommitmentUpdateLock() *sync.Mutex {
	return &s.commitmentUpdateLock
}

func (s *SharedServiceContainer) BlockExecutionLock() *sync.Mutex {
	return &s.blockExecutionLock
}

func (s *SharedServiceContainer) NextFeeRecipient() common.Address {
	return s.nextFeeRecipient
}

// assumes that the block execution lock is being held
func (s *SharedServiceContainer) SetNextFeeRecipient(nextFeeRecipient common.Address) {
	s.nextFeeRecipient = nextFeeRecipient
}

func (s *SharedServiceContainer) BridgeAddresses() map[string]*params.AstriaBridgeAddressConfig {
	return s.bridgeAddresses
}

func (s *SharedServiceContainer) BridgeAllowedAssets() map[string]struct{} {
	return s.bridgeAllowedAssets
}

func (s *SharedServiceContainer) TrustedBuilderPublicKey() ed25519.PublicKey {
	return s.trustedBuilderPublicKey
}
