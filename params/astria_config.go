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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

type AstriaTransactionType int

const (
	SequencedData AstriaTransactionType = iota
	Deposit
	PriceFeedData
)

var (
	AstriaTransactionTypeMap = map[string]AstriaTransactionType{
		"sequencedData": SequencedData,
		"deposit":       Deposit,
		"priceFeedData": PriceFeedData,
	}
	AstriaTransactionTypeReverseMap = map[AstriaTransactionType]string{
		SequencedData: "sequencedData",
		Deposit:       "deposit",
		PriceFeedData: "priceFeedData",
	}
)

type AstriaForks struct {
	orderedForks []AstriaForkData // sorted in descending order by Height
	forkMap      map[string]AstriaForkConfig
}

type AstriaForkConfig struct {
	Height              uint64                      `json:"height"`
	Halt                bool                        `json:"halt,omitempty"`
	SnapshotChecksum    string                      `json:"snapshotChecksum,omitempty"`
	ExtraDataOverride   hexutil.Bytes               `json:"extraDataOverride,omitempty"`
	FeeCollector        *common.Address             `json:"feeCollector,omitempty"`
	EIP1559Params       *AstriaEIP1559Params        `json:"eip1559Params,omitempty"`
	Sequencer           *AstriaSequencerConfig      `json:"sequencer,omitempty"`
	Celestia            *AstriaCelestiaConfig       `json:"celestia,omitempty"`
	BridgeAddresses     []AstriaBridgeAddressConfig `json:"bridgeAddresses,omitempty"`
	Oracle              *AstriaOracleConfig         `json:"oracle,omitempty"`
	Precompiles         []PrecompileConfig          `json:"precompiles,omitempty"`
	AppSpecificOrdering []string                    `json:"appSpecificOrdering"`
}

type AstriaForkData struct {
	Name                string
	Height              uint64
	StopHeight          uint64
	Halt                bool
	SnapshotChecksum    string
	ExtraDataOverride   hexutil.Bytes
	ExtraData           hexutil.Bytes
	FeeCollector        common.Address
	EIP1559Params       AstriaEIP1559Params
	Sequencer           AstriaSequencerConfig
	Celestia            AstriaCelestiaConfig
	BridgeAddresses     map[string]*AstriaBridgeAddressConfig // astria bridge addess to config for that bridge account
	BridgeAllowedAssets map[string]struct{}                   // a set of allowed asset IDs structs are left empty
	Oracle              AstriaOracleConfig
	Precompiles         map[common.Address]*PrecompileType
	AppSpecificOrdering []AstriaTransactionType
}

type AstriaSequencerConfig struct {
	ChainID       string `json:"chainId"`
	AddressPrefix string `json:"addressPrefix,omitempty"`
	StartHeight   uint64 `json:"startHeight"`
}

type AstriaCelestiaConfig struct {
	ChainID                  string `json:"chainId"`
	StartHeight              uint64 `json:"startHeight"`
	SearchHeightMaxLookAhead uint64 `json:"searchHeightMaxLookAhead"`
}

type AstriaEIP1559Params struct {
	MinBaseFee               uint64 `json:"minBaseFee"`
	ElasticityMultiplier     uint64 `json:"elasticityMultiplier"`
	BaseFeeChangeDenominator uint64 `json:"baseFeeChangeDenominator"`
}

type AstriaOracleConfig struct {
	ContractAddress *common.Address `json:"contractAddress"`
	CallerAddress   *common.Address `json:"callerAddress"`
}

func (c *ChainConfig) AstriaExtraData(height uint64) []byte {
	fork := c.GetAstriaForks().GetForkAtHeight(height)
	return fork.ExtraData
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

			// maps are pointers, so we need to create new maps for each fork
			orderedForks[i].BridgeAddresses = make(map[string]*AstriaBridgeAddressConfig)
			orderedForks[i].BridgeAllowedAssets = make(map[string]struct{})
			orderedForks[i].Precompiles = make(map[common.Address]*PrecompileType)

			// Copy previous fork's maps if needed
			for k, v := range orderedForks[i-1].BridgeAddresses {
				orderedForks[i].BridgeAddresses[k] = v
			}
			for k, v := range orderedForks[i-1].BridgeAllowedAssets {
				orderedForks[i].BridgeAllowedAssets[k] = v
			}
			for k, v := range orderedForks[i-1].Precompiles {
				orderedForks[i].Precompiles[k] = v
			}
		} else {
			orderedForks[i] = GetDefaultAstriaForkData()
		}

		// Set fork-specific fields
		orderedForks[i].Name = nh.name
		orderedForks[i].Height = currentFork.Height
		orderedForks[i].StopHeight = 0
		orderedForks[i].Halt = currentFork.Halt
		orderedForks[i].SnapshotChecksum = ""

		// set stop height of previous fork
		if i > 0 {
			orderedForks[i-1].StopHeight = currentFork.Height - 1
		}

		// Override with any new values from current fork
		if currentFork.SnapshotChecksum != "" {
			orderedForks[i].SnapshotChecksum = currentFork.SnapshotChecksum
		}

		if currentFork.AppSpecificOrdering != nil {
			log.Debug("currentFork.AppSpecificOrdering before", "ordering", fmt.Sprintf("%v", currentFork.AppSpecificOrdering))
		} else {
			log.Debug("currentFork.AppSpecificOrdering before", "ordering", "nil")
		}
		// Apply ordering rules to the current fork
		if currentFork.AppSpecificOrdering == nil && i > 0 {
			// Carry over ordering from previous fork
			orderedForks[i].AppSpecificOrdering = orderedForks[i-1].AppSpecificOrdering
		} else if currentFork.AppSpecificOrdering != nil && len(currentFork.AppSpecificOrdering) == 0 {
			// Explicitly reset ordering to empty (not nil)
			orderedForks[i].AppSpecificOrdering = []AstriaTransactionType{}
		} else if currentFork.AppSpecificOrdering != nil {
			// Apply set ordering rules to the current fork
			orderedForks[i].AppSpecificOrdering = make([]AstriaTransactionType, len(currentFork.AppSpecificOrdering))
			for j, transactionType := range currentFork.AppSpecificOrdering {
				if _, ok := AstriaTransactionTypeMap[transactionType]; !ok {
					return nil, fmt.Errorf("invalid transaction type: %s", transactionType)
				}
				orderedForks[i].AppSpecificOrdering[j] = AstriaTransactionTypeMap[transactionType]
			}
		}
		if orderedForks[i].AppSpecificOrdering != nil {
			log.Debug("orderedForks.AppSpecificOrdering after", "index", i, "ordering", fmt.Sprintf("%v", orderedForks[i].AppSpecificOrdering))
		} else {
			log.Debug("orderedForks.AppSpecificOrdering after", "index", i, "ordering", "nil")
		}

		if currentFork.FeeCollector != nil {
			orderedForks[i].FeeCollector = *currentFork.FeeCollector
		}

		if currentFork.EIP1559Params != nil {
			orderedForks[i].EIP1559Params = *currentFork.EIP1559Params
		}

		if currentFork.Oracle != nil {
			orderedForks[i].Oracle = *currentFork.Oracle
		}

		if currentFork.Sequencer != nil {
			orderedForks[i].Sequencer = *currentFork.Sequencer
		} else {
			orderedForks[i].Sequencer.StartHeight = orderedForks[i-1].Sequencer.StartHeight + currentFork.Height - orderedForks[i-1].Height
		}

		if currentFork.Celestia != nil {
			orderedForks[i].Celestia = *currentFork.Celestia
		}

		if len(currentFork.BridgeAddresses) > 0 {
			for _, cfg := range currentFork.BridgeAddresses {
				err := cfg.Validate(orderedForks[i].Sequencer.AddressPrefix)
				if orderedForks[i].BridgeAddresses[cfg.BridgeAddress] != nil {
					return nil, fmt.Errorf("duplicate bridge address config: %s", cfg.BridgeAddress)
				}
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

		// add precompiles
		if len(currentFork.Precompiles) > 0 {
			for _, precompile := range currentFork.Precompiles {
				if orderedForks[i].Precompiles[precompile.Address] != nil {
					return nil, fmt.Errorf("duplicate precompile address config: %s", precompile.Address)
				}
				orderedForks[i].Precompiles[precompile.Address] = &precompile.Type
			}
		}

		orderedForks[i].ExtraDataOverride = currentFork.ExtraDataOverride
		if currentFork.ExtraDataOverride == nil {
			data, err := json.Marshal(orderedForks[i].ToConfig())
			if err != nil {
				return nil, fmt.Errorf("failed to marshal astria forks: %w", err)
			}
			orderedForks[i].ExtraData = crypto.Keccak256(data)
		} else {
			orderedForks[i].ExtraData = currentFork.ExtraDataOverride
		}
	}

	if err := validateAstriaForks(orderedForks); err != nil {
		return nil, err
	}

	// Sort orderedForks in descending order by Height
	sort.Slice(orderedForks, func(i, j int) bool {
		return orderedForks[i].Height > orderedForks[j].Height
	})

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

			if fork.Celestia.SearchHeightMaxLookAhead == 0 {
				return fmt.Errorf("fork %s: celestia search height max lookahead not set", fork.Name)
			}

			if fork.FeeCollector == (common.Address{}) {
				log.Warn("fee asset collectors not set, assets will be burned", "fork", fork.Name)
			}

			for _, pType := range fork.Precompiles {
				if err := pType.Validate(); err != nil {
					return fmt.Errorf("fork %s: invalid precompile %s", fork.Name, *pType)
				}
			}

			if err := validateAppSpecificOrdering(fork.AppSpecificOrdering); err != nil {
				return fmt.Errorf("fork %s: %w", fork.Name, err)
			}
		} else {
			log.Warn("fork will halt", "fork", fork.Name, "height", fork.Height)
		}
	}

	return nil
}

func validateAppSpecificOrdering(appSpecificOrdering []AstriaTransactionType) error {
	// ordering not set
	if appSpecificOrdering == nil {
		return nil
	}
	// ordering explicitly set to [] (reset ordering)
	if appSpecificOrdering != nil && len(appSpecificOrdering) == 0 {
		return nil
	}
	transactionTypeSet := make(map[AstriaTransactionType]struct{})
	for _, transactionType := range appSpecificOrdering {
		if _, ok := transactionTypeSet[transactionType]; ok {
			return fmt.Errorf("app specific ordering contains duplicate transaction types")
		}
		transactionTypeSet[transactionType] = struct{}{}
	}
	if len(appSpecificOrdering) != len(AstriaTransactionTypeMap) {
		return fmt.Errorf("app specific ordering must contain all transaction types")
	}
	return nil
}

func GetDefaultAstriaForkData() AstriaForkData {
	return AstriaForkData{
		Name:         "default",
		Height:       1,
		FeeCollector: common.Address{},
		EIP1559Params: AstriaEIP1559Params{
			MinBaseFee:               0,
			ElasticityMultiplier:     DefaultElasticityMultiplier,
			BaseFeeChangeDenominator: DefaultBaseFeeChangeDenominator,
		},
		BridgeAddresses:     make(map[string]*AstriaBridgeAddressConfig),
		BridgeAllowedAssets: make(map[string]struct{}),
		Precompiles:         make(map[common.Address]*PrecompileType),
	}
}

func (c *AstriaForks) GetForkAtHeight(height uint64) AstriaForkData {
	idx := sort.Search(len(c.orderedForks), func(i int) bool {
		return c.orderedForks[i].Height <= height
	})
	// no named fork at this height
	if idx == len(c.orderedForks) {
		return GetDefaultAstriaForkData()
	}
	return c.orderedForks[idx]
}

func (c *AstriaForks) GetNextForkAtHeight(height uint64) *AstriaForkData {
	// orderedForks are sorted in descending order; iterate from the end
	for i := len(c.orderedForks) - 1; i >= 0; i-- {
		if c.orderedForks[i].Height > height {
			return &c.orderedForks[i]
		}
	}
	return nil
}

func (c *AstriaForks) MinBaseFeeAt(height uint64) *big.Int {
	return big.NewInt(0).SetUint64(c.GetForkAtHeight(height).EIP1559Params.MinBaseFee)
}

// BaseFeeChangeDenominator bounds the amount the base fee can change between blocks.
func (c *AstriaForks) BaseFeeChangeDenominatorAt(height uint64) uint64 {
	return c.GetForkAtHeight(height).EIP1559Params.BaseFeeChangeDenominator
}

// ElasticityMultiplier bounds the maximum gas limit an EIP-1559 block may have.
func (c *AstriaForks) ElasticityMultiplierAt(height uint64) uint64 {
	return c.GetForkAtHeight(height).EIP1559Params.ElasticityMultiplier
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
	AssetDenom     string                  `json:"assetDenom"`
	AssetPrecision uint16                  `json:"assetPrecision"`
	Erc20Asset     *AstriaErc20AssetConfig `json:"erc20Asset,omitempty"`
}

type AstriaErc20AssetConfig struct {
	ContractAddress   common.Address `json:"contractAddress"`
	ContractPrecision uint16         `json:"contractPrecision"`
}

func (abc *AstriaBridgeAddressConfig) Validate(expectedPrefix string) error {
	prefix, byteAddress, err := bech32.Decode(abc.BridgeAddress)
	if err != nil {
		return fmt.Errorf("bridge address must be a bech32 encoded string")
	}
	byteAddress, err = bech32.ConvertBits(byteAddress, 5, 8, false)
	if err != nil {
		return fmt.Errorf("failed to convert address to 8 bit")
	}
	if prefix != expectedPrefix {
		return fmt.Errorf("bridge address must have prefix %s", expectedPrefix)
	}
	if len(byteAddress) != 20 {
		return fmt.Errorf("bridge address must have resolve to 20 byte address, got %d", len(byteAddress))
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

type PrecompileConfig struct {
	Address common.Address `json:"address"`
	Type    PrecompileType `json:"precompileType"`
}

// PrecompileType represents the type of precompile to enable
type PrecompileType string

const (
	PrecompileBase64 PrecompileType = "base64"
	// Add other precompile types here as needed
)

// Validate ensures the PrecompileType is a known value
func (p PrecompileType) Validate() error {
	switch p {
	case PrecompileBase64:
		return nil
	default:
		return fmt.Errorf("unknown precompile type: %s", p)
	}
}

// converts an AstriaForkData to AstriaForkConfig
func (fd AstriaForkData) ToConfig() AstriaForkConfig {
	// Convert bridge addresses from map to slice
	bridgeAddrs := make([]string, 0, len(fd.BridgeAddresses))
	for addr := range fd.BridgeAddresses {
		bridgeAddrs = append(bridgeAddrs, addr)
	}
	bridgeAddressConfigs := make([]AstriaBridgeAddressConfig, len(fd.BridgeAddresses))
	sort.Strings(bridgeAddrs)
	for _, addr := range bridgeAddrs {
		bridgeAddressConfigs = append(bridgeAddressConfigs, *fd.BridgeAddresses[addr])
	}

	precompileAddresses := make([]common.Address, 0, len(fd.Precompiles))
	for addr := range fd.Precompiles {
		precompileAddresses = append(precompileAddresses, addr)
	}
	precompiles := make([]PrecompileConfig, len(fd.Precompiles))
	sort.Slice(precompileAddresses, func(i, j int) bool {
		return precompileAddresses[i].Hex() < precompileAddresses[j].Hex()
	})
	for _, addr := range precompileAddresses {
		precompile := PrecompileConfig{
			Address: addr,
			Type:    *fd.Precompiles[addr],
		}
		precompiles = append(precompiles, precompile)
	}

	var appSpecificOrdering []string
	if fd.AppSpecificOrdering != nil {
		for _, transactionType := range fd.AppSpecificOrdering {
			appSpecificOrdering = append(appSpecificOrdering, AstriaTransactionTypeReverseMap[transactionType])
		}
		if len(appSpecificOrdering) == len(AstriaTransactionTypeMap) {
			// TODO: should setting the ordering rules wrong halt the chain?
			appSpecificOrdering = nil
		}
	}

	config := AstriaForkConfig{
		Height:              fd.Height,
		Halt:                fd.Halt,
		SnapshotChecksum:    fd.SnapshotChecksum,
		ExtraDataOverride:   fd.ExtraDataOverride,
		FeeCollector:        &fd.FeeCollector,
		EIP1559Params:       &fd.EIP1559Params,
		Sequencer:           &fd.Sequencer,
		Celestia:            &fd.Celestia,
		BridgeAddresses:     bridgeAddressConfigs,
		Oracle:              &fd.Oracle,
		Precompiles:         precompiles,
		AppSpecificOrdering: appSpecificOrdering,
	}

	return config
}

func (fd AstriaForkData) MarshalJSON() ([]byte, error) {
	return json.Marshal(fd.ToConfig())
}
