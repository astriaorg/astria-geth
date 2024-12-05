package params

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sort"

	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

type AstriaForks struct {
	orderedForks []AstriaForkData
	forkMap      map[string]AstriaForkConfig
}

type AstriaForkConfig struct {
	Height            uint64                      `json:"height"`
	Halt              bool                        `json:"halt,omitempty"`
	SnapshotChecksum  string                      `json:"snapshotChecksum,omitempty"`
	ExtraDataOverride hexutil.Bytes               `json:"extraDataOverride,omitempty"`
	FeeCollector      *common.Address             `json:"feeCollector,omitempty"`
	EIP1559Params     *AstriaEIP1559Param         `json:"eip1559Params,omitempty"`
	Sequencer         *AstriaSequencerConfig      `json:"sequencer,omitempty"`
	Celestia          *AstriaCelestiaConfig       `json:"celestia,omitempty"`
	BridgeAddresses   []AstriaBridgeAddressConfig `json:"bridgeAddresses,omitempty"`
}

type AstriaForkData struct {
	Name                string
	Height              uint64
	Halt                bool
	SnapshotChecksum    string
	ExtraDataOverride   hexutil.Bytes
	FeeCollector        common.Address
	EIP1559Params       AstriaEIP1559Param
	Sequencer           AstriaSequencerConfig
	Celestia            AstriaCelestiaConfig
	BridgeAddresses     map[string]*AstriaBridgeAddressConfig // astria bridge addess to config for that bridge account
	BridgeAllowedAssets map[string]struct{}                   // a set of allowed asset IDs structs are left empty
}

type AstriaSequencerConfig struct {
	ChainID           string `json:"chainId"`
	AddressPrefix     string `json:"addressPrefix,omitempty"`
	StartHeight       uint32 `json:"startHeight"`
	StopHeight        uint32 `json:"-"`
	RollupStartHeight uint64 `json:"-"`
}

type AstriaCelestiaConfig struct {
	ChainID        string `json:"chainId"`
	StartHeight    uint64 `json:"startHeight"`
	HeightVariance uint64 `json:"heightVariance"`
}

type AstriaEIP1559Param struct {
	MinBaseFee               uint64 `json:"minBaseFee"`
	ElasticityMultiplier     uint64 `json:"elasticityMultiplier"`
	BaseFeeChangeDenominator uint64 `json:"baseFeeChangeDenominator"`
}

func (c *ChainConfig) AstriaExtraData(height uint64) []byte {
	fork := c.GetAstriaForks().GetForkAtHeight(height)
	if fork.ExtraDataOverride != nil {
		return fork.ExtraDataOverride
	}

	// create default extradata
	extra, _ := rlp.EncodeToBytes([]interface{}{
		c.AstriaRollupName,
		fork.Sequencer.StartHeight,
		fork.Celestia.StartHeight,
		fork.Celestia.HeightVariance,
	})
	if uint64(len(extra)) > MaximumExtraDataSize {
		log.Warn("Miner extra data exceed limit", "extra", hexutil.Bytes(extra), "limit", MaximumExtraDataSize)
		extra = nil
	}
	return extra
}

func NewAstriaForks(forks map[string]AstriaForkConfig) (*AstriaForks, error) {
	if forks == nil {
		return &AstriaForks{
			orderedForks: []AstriaForkData{},
			forkMap:      make(map[string]AstriaForkConfig),
		}, nil
	}

	// Create sorted array of fork names and heights
	type nameHeight struct {
		name   string
		height uint64
	}
	sortedNames := make([]nameHeight, 0, len(forks))
	for name, fork := range forks {
		sortedNames = append(sortedNames, nameHeight{name, fork.Height})
	}
	sort.Slice(sortedNames, func(i, j int) bool {
		return sortedNames[i].height < sortedNames[j].height
	})

	nativeBridgeSeen := false
	orderedForks := make([]AstriaForkData, len(sortedNames))

	for i, nh := range sortedNames {
		currentFork := forks[nh.name]

		if i > 0 {
			// Copy previous fork's configuration as the base
			orderedForks[i] = orderedForks[i-1]
		} else {
			// set default values
			orderedForks[i] = GetDefaultAstriaForkData()
		}

		// Set fork-specific fields
		orderedForks[i].Name = nh.name
		orderedForks[i].Height = currentFork.Height
		orderedForks[i].Halt = currentFork.Halt
		orderedForks[i].SnapshotChecksum = ""

		// Override with any new values from current fork
		if currentFork.SnapshotChecksum != "" {
			orderedForks[i].SnapshotChecksum = currentFork.SnapshotChecksum
		}

		if currentFork.ExtraDataOverride != nil {
			orderedForks[i].ExtraDataOverride = currentFork.ExtraDataOverride
		}

		if currentFork.FeeCollector != nil {
			orderedForks[i].FeeCollector = *currentFork.FeeCollector
		}

		if currentFork.EIP1559Params != nil {
			orderedForks[i].EIP1559Params = *currentFork.EIP1559Params
		}

		if currentFork.Sequencer != nil {
			orderedForks[i].Sequencer = *currentFork.Sequencer
			orderedForks[i].Sequencer.RollupStartHeight = currentFork.Height
			// set stop height for previous fork if sequencer data is changed
			if i > 0 {
				orderedForks[i-1].Sequencer.StopHeight = orderedForks[i-1].Sequencer.StartHeight + uint32(currentFork.Height-orderedForks[i-1].Sequencer.RollupStartHeight)
			}
		}

		if currentFork.Celestia != nil {
			orderedForks[i].Celestia = *currentFork.Celestia
		}

		if len(currentFork.BridgeAddresses) > 0 {
			for _, cfg := range currentFork.BridgeAddresses {
				err := cfg.Validate(orderedForks[i].Sequencer.AddressPrefix)
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
				orderedForks[i].BridgeAddresses[cfg.BridgeAddress] = &bridgeCfg
				orderedForks[i].BridgeAllowedAssets[cfg.AssetDenom] = struct{}{}
				if cfg.Erc20Asset == nil {
					log.Info("bridge for sequencer native asset initialized", "bridgeAddress", cfg.BridgeAddress, "assetDenom", cfg.AssetDenom)
				} else {
					log.Info("bridge for ERC20 asset initialized", "bridgeAddress", cfg.BridgeAddress, "assetDenom", cfg.AssetDenom, "contractAddress", cfg.Erc20Asset.ContractAddress)
				}
			}
		}
	}

	if err := validateAstriaForks(orderedForks); err != nil {
		return nil, err
	}

	return &AstriaForks{
		orderedForks: orderedForks,
		forkMap:      forks,
	}, nil
}

func validateAstriaForks(forks []AstriaForkData) error {
	for _, fork := range forks {
		if !fork.Halt {
			if fork.Sequencer.ChainID == "" {
				return fmt.Errorf("fork %s: sequencer chain ID not set", fork.Name)
			}

			if fork.Sequencer.StartHeight == 0 {
				return fmt.Errorf("fork %s: sequencer initial height not set", fork.Name)
			}

			if fork.Celestia.ChainID == "" {
				return fmt.Errorf("fork %s: celestia chain ID not set", fork.Name)
			}

			if fork.Celestia.StartHeight == 0 {
				return fmt.Errorf("fork %s: celestia initial height not set", fork.Name)
			}

			if fork.Celestia.HeightVariance == 0 {
				return fmt.Errorf("fork %s: celestia height variance not set", fork.Name)
			}

			if fork.FeeCollector == (common.Address{}) {
				log.Warn("fee asset collectors not set, assets will be burned", "fork", fork.Name)
			}
		} else {
			log.Warn("fork will halt", "fork", fork.Name, "height", fork.Height)
		}
	}

	return nil
}

func GetDefaultAstriaForkData() AstriaForkData {
	return AstriaForkData{
		Height:       1,
		FeeCollector: common.Address{},
		EIP1559Params: AstriaEIP1559Param{
			MinBaseFee:               0,
			ElasticityMultiplier:     DefaultElasticityMultiplier,
			BaseFeeChangeDenominator: DefaultBaseFeeChangeDenominator,
		},
		BridgeAddresses:     make(map[string]*AstriaBridgeAddressConfig),
		BridgeAllowedAssets: make(map[string]struct{}),
	}
}

func (c *AstriaForks) GetForkAtHeight(height uint64) AstriaForkData {
	if len(c.orderedForks) == 0 {
		return GetDefaultAstriaForkData()
	}

	if height < c.orderedForks[0].Height {
		return GetDefaultAstriaForkData()
	}

	idx := sort.Search(len(c.orderedForks), func(i int) bool {
		return c.orderedForks[i].Height > height
	}) - 1

	if idx < 0 {
		return GetDefaultAstriaForkData()
	}

	return c.orderedForks[idx]
}

func (c *AstriaForks) GetNextForkAtHeight(height uint64) *AstriaForkData {
	idx := sort.Search(len(c.orderedForks), func(i int) bool {
		return c.orderedForks[i].Height > height
	})
	if idx < 0 {
		return nil
	}
	return &c.orderedForks[idx]
}

func (c *AstriaForks) MinBaseFeeAt(height uint64) *big.Int {
	return big.NewInt(0).SetUint64(c.GetForkAtHeight(height).EIP1559Params.MinBaseFee)
}

func (c *AstriaForks) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.forkMap)
}

func (c *AstriaForks) UnmarshalJSON(data []byte) error {
	var forkMap map[string]AstriaForkConfig
	if err := json.Unmarshal(data, &forkMap); err != nil {
		return err
	}

	newForks, err := NewAstriaForks(forkMap)
	if err != nil {
		return err
	}

	*c = *newForks
	return nil
}

func (c *ChainConfig) GetAstriaForks() *AstriaForks {
	if c.AstriaForks == nil {
		forks, _ := NewAstriaForks(nil)
		return forks
	}
	return c.AstriaForks
}

type AstriaBridgeAddressConfig struct {
	BridgeAddress  string                  `json:"bridgeAddress"`
	SenderAddress  common.Address          `json:"senderAddress,omitempty"`
	StartHeight    uint32                  `json:"startHeight"`
	AssetDenom     string                  `json:"assetDenom"`
	AssetPrecision uint16                  `json:"assetPrecision"`
	Erc20Asset     *AstriaErc20AssetConfig `json:"erc20Asset,omitempty"`
}

type AstriaErc20AssetConfig struct {
	ContractAddress   common.Address `json:"contractAddress"`
	ContractPrecision uint16         `json:"contractPrecision"`
}

func (abc *AstriaBridgeAddressConfig) Validate(genesisPrefix string) error {
	prefix, byteAddress, err := bech32.Decode(abc.BridgeAddress)
	if err != nil {
		return fmt.Errorf("bridge address must be a bech32 encoded string")
	}
	byteAddress, err = bech32.ConvertBits(byteAddress, 5, 8, false)
	if err != nil {
		return fmt.Errorf("failed to convert address to 8 bit")
	}
	if prefix != genesisPrefix {
		return fmt.Errorf("bridge address must have prefix %s", genesisPrefix)
	}
	if len(byteAddress) != 20 {
		return fmt.Errorf("bridge address must have resolve to 20 byte address, got %d", len(byteAddress))
	}
	if abc.StartHeight == 0 {
		return fmt.Errorf("start height must be greater than 0")
	}
	if abc.AssetDenom == "" {
		return fmt.Errorf("asset denom must be set")
	}
	if abc.Erc20Asset == nil && abc.AssetPrecision > 18 {
		return fmt.Errorf("asset precision of native asset must be less than or equal to 18")
	}
	if abc.Erc20Asset != nil && abc.AssetPrecision > abc.Erc20Asset.ContractPrecision {
		return fmt.Errorf("asset precision must be less than or equal to contract precision")
	}

	return nil
}

func (abc *AstriaBridgeAddressConfig) ScaledDepositAmount(deposit *big.Int) *big.Int {
	var exponent uint16
	if abc.Erc20Asset != nil {
		exponent = abc.Erc20Asset.ContractPrecision - abc.AssetPrecision
	} else {
		exponent = 18 - abc.AssetPrecision
	}
	multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(exponent)), nil)

	return new(big.Int).Mul(deposit, multiplier)
}
