package shared

import (
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
	bridgeSenderAddress common.Address                               // address from which AstriaBridgeableERC20 contracts are called

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

			if cfg.Erc20Asset != nil && bc.Config().AstriaBridgeSenderAddress == (common.Address{}) {
				return nil, errors.New("astria bridge sender address must be set for bridged ERC20 assets")
			}

			bridgeAddresses[cfg.BridgeAddress] = &cfg
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

	sharedServiceContainer := &SharedServiceContainer{
		eth:                 eth,
		bc:                  bc,
		bridgeAddresses:     bridgeAddresses,
		bridgeAllowedAssets: bridgeAllowedAssets,
		bridgeSenderAddress: bc.Config().AstriaBridgeSenderAddress,
		nextFeeRecipient:    nextFeeRecipient,
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

func (s *SharedServiceContainer) BridgeSenderAddress() common.Address {
	return s.bridgeSenderAddress
}
