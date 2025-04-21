package params

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"
)

var precompileBase64Val = PrecompileBase64

func TestAstriaEIP1559Params(t *testing.T) {
	astriaForks, _ := NewAstriaForks(map[string]AstriaForkConfig{
		"genesis": {
			Height: 1,
			Sequencer: &AstriaSequencerConfig{
				ChainID:     "sequencer-test-chain-0",
				StartHeight: 2,
			},
			Celestia: &AstriaCelestiaConfig{
				ChainID:                  "mocha-4",
				StartHeight:              2,
				SearchHeightMaxLookAhead: 10,
			},
			EIP1559Params: &AstriaEIP1559Params{
				MinBaseFee:               45000000000,
				ElasticityMultiplier:     4,
				BaseFeeChangeDenominator: 100,
			},
		},
		"fork1": {
			Height: 15,
			EIP1559Params: &AstriaEIP1559Params{
				MinBaseFee:               15000000000,
				ElasticityMultiplier:     5,
				BaseFeeChangeDenominator: 50,
			},
		},
		"fork2": {
			Height: 101,
			EIP1559Params: &AstriaEIP1559Params{
				MinBaseFee:               120000000,
				ElasticityMultiplier:     11,
				BaseFeeChangeDenominator: 250,
			},
		},
	})

	minBaseTests := map[uint64]*big.Int{
		0:      common.Big0,
		1:      big.NewInt(45000000000),
		2:      big.NewInt(45000000000),
		14:     big.NewInt(45000000000),
		15:     big.NewInt(15000000000),
		16:     big.NewInt(15000000000),
		50:     big.NewInt(15000000000),
		100:    big.NewInt(15000000000),
		101:    big.NewInt(120000000),
		102:    big.NewInt(120000000),
		123456: big.NewInt(120000000),
	}

	for height, expected := range minBaseTests {
		if got := astriaForks.MinBaseFeeAt(height); got.Cmp(expected) != 0 {
			t.Errorf("MinBaseFeeAt(%d): expected %v, got %v", height, expected, got)
		}
	}

	elasticityMultiplierTests := map[uint64]uint64{
		0:      DefaultElasticityMultiplier,
		1:      4,
		2:      4,
		14:     4,
		15:     5,
		16:     5,
		50:     5,
		100:    5,
		101:    11,
		102:    11,
		123456: 11,
	}

	for height, expected := range elasticityMultiplierTests {
		if got := astriaForks.ElasticityMultiplierAt(height); got != expected {
			t.Errorf("ElasticityMultiplierAt(%d): expected %v, got %v", height, expected, got)
		}
	}

	baseFeeChangeDenominatorTests := map[uint64]uint64{
		0:      DefaultBaseFeeChangeDenominator,
		1:      100,
		2:      100,
		14:     100,
		15:     50,
		16:     50,
		50:     50,
		100:    50,
		101:    250,
		102:    250,
		123456: 250,
	}

	for height, expected := range baseFeeChangeDenominatorTests {
		if got := astriaForks.BaseFeeChangeDenominatorAt(height); got != expected {
			t.Errorf("BaseFeeChangeDenominatorAt(%d): expected %v, got %v", height, expected, got)
		}
	}
}

func TestAstriaBridgeConfigValidation(t *testing.T) {
	bridgeAddressKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	bridgeAddress := crypto.PubkeyToAddress(bridgeAddressKey.PublicKey)
	toEncode, _ := bech32.ConvertBits(bridgeAddress.Bytes(), 8, 5, false)
	bridgeAddressBech32, _ := bech32.EncodeM("astria", toEncode)

	erc20AssetKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	erc20Asset := crypto.PubkeyToAddress(erc20AssetKey.PublicKey)

	tests := []struct {
		description string
		config      AstriaBridgeAddressConfig
		wantErr     error
	}{
		{
			description: "invalid bridge address, non bech32m",
			config: AstriaBridgeAddressConfig{
				BridgeAddress:  "rand address",
				AssetDenom:     "nria",
				AssetPrecision: 18,
				Erc20Asset:     nil,
			},
			wantErr: fmt.Errorf("bridge address must be a bech32 encoded string"),
		},
		{
			description: "invalid bridge address, invalid prefix",
			config: AstriaBridgeAddressConfig{
				BridgeAddress:  "badprefix1u54zke43yc2tpaecvjqj4uy7d3mdmkrj4vch35",
				AssetDenom:     "nria",
				AssetPrecision: 18,
				Erc20Asset:     nil,
			},
			wantErr: fmt.Errorf("bridge address must have prefix astria"),
		},
		{
			description: "invalid bridge address",
			config: AstriaBridgeAddressConfig{
				BridgeAddress:  "astria1u54zke43yc2tpaecvjqj4uy7d3mdmkqjjq96x",
				AssetDenom:     "nria",
				AssetPrecision: 18,
				Erc20Asset:     nil,
			},
			wantErr: fmt.Errorf("bridge address must have resolve to 20 byte address, got 19"),
		},
		{
			description: "invalid asset denom",
			config: AstriaBridgeAddressConfig{
				BridgeAddress:  bridgeAddressBech32,
				AssetDenom:     "",
				AssetPrecision: 18,
				Erc20Asset:     nil,
			},
			wantErr: fmt.Errorf("asset denom must be set"),
		},
		{
			description: "invalid asset precision",
			config: AstriaBridgeAddressConfig{
				BridgeAddress:  bridgeAddressBech32,
				AssetDenom:     "nria",
				AssetPrecision: 22,
				Erc20Asset:     nil,
			},
			wantErr: fmt.Errorf("asset precision of native asset must be less than or equal to 18"),
		},
		{
			description: "invalid contract precision",
			config: AstriaBridgeAddressConfig{
				BridgeAddress:  bridgeAddressBech32,
				AssetDenom:     "nria",
				AssetPrecision: 22,
				Erc20Asset: &AstriaErc20AssetConfig{
					ContractAddress:   erc20Asset,
					ContractPrecision: 18,
				},
			},
			wantErr: fmt.Errorf("asset precision must be less than or equal to contract precision"),
		},
		{
			description: "erc20 assets supported",
			config: AstriaBridgeAddressConfig{
				BridgeAddress:  bridgeAddressBech32,
				AssetDenom:     "nria",
				AssetPrecision: 18,
				Erc20Asset: &AstriaErc20AssetConfig{
					ContractAddress:   erc20Asset,
					ContractPrecision: 18,
				},
			},
			wantErr: nil,
		},
		{
			description: "valid config",
			config: AstriaBridgeAddressConfig{
				BridgeAddress:  bridgeAddressBech32,
				AssetDenom:     "nria",
				AssetPrecision: 18,
				Erc20Asset:     nil,
			},
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := test.config.Validate("astria")
			if test.wantErr != nil && err == nil {
				t.Errorf("expected error, got nil")
			}
			if test.wantErr == nil && err != nil {
				t.Errorf("unexpected error %v", err)
			}

			if !reflect.DeepEqual(err, test.wantErr) {
				t.Errorf("error mismatch:\nconfig: %v\nerr: %v\nwant: %v", test.config, err, test.wantErr)
			}
		})
	}
}

func TestGetForkAtHeight(t *testing.T) {
	forkMap := map[string]AstriaForkConfig{
		"fork1": {
			Height: 1,
			Sequencer: &AstriaSequencerConfig{
				ChainID:     "chain1",
				StartHeight: 1,
			},
			Celestia: &AstriaCelestiaConfig{
				ChainID:                  "celestia1",
				StartHeight:              1,
				SearchHeightMaxLookAhead: 100,
			},
		},
		"fork2": {
			Height: 1000,
			Sequencer: &AstriaSequencerConfig{
				ChainID:     "chain2",
				StartHeight: 2,
			},
			Celestia: &AstriaCelestiaConfig{
				ChainID:                  "celestia2",
				StartHeight:              2,
				SearchHeightMaxLookAhead: 200,
			},
		},
		"fork3": {
			Height: 2000,
			Sequencer: &AstriaSequencerConfig{
				ChainID:     "chain3",
				StartHeight: 3,
			},
			Celestia: &AstriaCelestiaConfig{
				ChainID:                  "celestia3",
				StartHeight:              3,
				SearchHeightMaxLookAhead: 300,
			},
		},
	}

	forks, err := NewAstriaForks(forkMap)
	if err != nil {
		t.Fatalf("failed to create forks: %v", err)
	}

	tests := []struct {
		description string
		height      uint64
		wantFork    string
	}{
		{
			description: "height 1 returns first fork",
			height:      1,
			wantFork:    "fork1",
		},
		{
			description: "height 500 returns first fork",
			height:      500,
			wantFork:    "fork1",
		},
		{
			description: "height 1000 returns second fork",
			height:      1000,
			wantFork:    "fork2",
		},
		{
			description: "height 1500 returns second fork",
			height:      1500,
			wantFork:    "fork2",
		},
		{
			description: "height 2000 returns third fork",
			height:      2000,
			wantFork:    "fork3",
		},
		{
			description: "height 3000 returns third fork",
			height:      3000,
			wantFork:    "fork3",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			fork := forks.GetForkAtHeight(test.height)
			if fork.Name != test.wantFork {
				t.Errorf("got fork %s, want %s", fork.Name, test.wantFork)
			}
		})
	}
}

func TestGetNextForkAtHeight(t *testing.T) {
	forkMap := map[string]AstriaForkConfig{
		"fork1": {
			Height: 1,
			Sequencer: &AstriaSequencerConfig{
				ChainID:     "chain1",
				StartHeight: 1,
			},
			Celestia: &AstriaCelestiaConfig{
				ChainID:                  "celestia1",
				StartHeight:              1,
				SearchHeightMaxLookAhead: 100,
			},
		},
		"fork2": {
			Height: 1000,
		},
		"fork3": {
			Height: 2000,
		},
	}

	forks, err := NewAstriaForks(forkMap)
	if err != nil {
		t.Fatalf("failed to create forks: %v", err)
	}

	tests := []struct {
		description string
		height      uint64
		wantFork    *string
	}{
		{
			description: "height 0 returns first fork",
			height:      0,
			wantFork:    strPtr("fork1"),
		},
		{
			description: "height 1 returns second fork",
			height:      1,
			wantFork:    strPtr("fork2"),
		},
		{
			description: "height 500 returns second fork",
			height:      500,
			wantFork:    strPtr("fork2"),
		},
		{
			description: "height 1500 returns third fork",
			height:      1500,
			wantFork:    strPtr("fork3"),
		},
		{
			description: "height 1999 returns third fork",
			height:      1999,
			wantFork:    strPtr("fork3"),
		},
		{
			description: "height 2000 returns nil",
			height:      2000,
			wantFork:    nil,
		},
		{
			description: "height 3000 returns nil",
			height:      3000,
			wantFork:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			nextFork := forks.GetNextForkAtHeight(test.height)
			if test.wantFork == nil {
				if nextFork != nil {
					t.Errorf("got fork %v, want nil", nextFork)
				}
				return
			}
			if nextFork == nil {
				t.Errorf("got nil fork, want %s", *test.wantFork)
				return
			}
			if nextFork.Name != *test.wantFork {
				t.Errorf("got fork %s, want %s", nextFork.Name, *test.wantFork)
			}
		})
	}
}

func strPtr(s string) *string {
	return &s
}

func TestAstriaForksInheritance(t *testing.T) {
	forkMap := map[string]AstriaForkConfig{
		"fork1": {
			Height: 1,
			EIP1559Params: &AstriaEIP1559Params{
				MinBaseFee:               1000,
				ElasticityMultiplier:     2,
				BaseFeeChangeDenominator: 8,
			},
			Sequencer: &AstriaSequencerConfig{
				ChainID:     "chain1",
				StartHeight: 100,
			},
			Celestia: &AstriaCelestiaConfig{
				ChainID:                  "celestia1",
				StartHeight:              1,
				SearchHeightMaxLookAhead: 100,
			},
		},
		"fork2": {
			Height: 200,
			// override EIP1559Params
			EIP1559Params: &AstriaEIP1559Params{
				MinBaseFee:               2000,
				ElasticityMultiplier:     4,
				BaseFeeChangeDenominator: 16,
			},
			Precompiles: []PrecompileConfig{
				{
					Address: common.HexToAddress("0x01"),
					Type:    PrecompileBase64,
				},
			},
		},
		"fork3": {
			Height: 251,
			// override EIP1559Params
			EIP1559Params: &AstriaEIP1559Params{
				MinBaseFee:               1000,
				ElasticityMultiplier:     4,
				BaseFeeChangeDenominator: 16,
			},
		},
		"fork4": {
			Height: 300,
			// override sequencer config
			Sequencer: &AstriaSequencerConfig{
				ChainID:     "chain3",
				StartHeight: 325,
			},
			Precompiles: []PrecompileConfig{
				{
					Address: common.HexToAddress("0x02"),
					Type:    PrecompileBase64,
				},
			},
		},
	}

	forks, err := NewAstriaForks(forkMap)
	if err != nil {
		t.Fatalf("failed to create forks: %v", err)
	}

	t.Logf("Forks configuration:")
	for i := len(forks.orderedForks) - 1; i >= 0; i-- {
		fork := forks.orderedForks[i]
		t.Logf("  Fork %s:", fork.Name)
		t.Logf("    Height: %d", fork.Height)
		t.Logf("    StopHeight: %d", fork.StopHeight)
		t.Logf("    Sequencer:")
		t.Logf("      ChainID: %s", fork.Sequencer.ChainID)
		t.Logf("      StartHeight: %d", fork.Sequencer.StartHeight)
		t.Logf("    Precompiles: %v", fork.Precompiles)
	}

	type testCheck struct {
		stopHeight                       uint64
		minBaseFee                       uint64
		elasticityMultiplier             uint64
		baseFeeChangeDenominator         uint64
		sequencerChainID                 string
		sequencerStartHeight             uint64
		celestiaChainID                  string
		celestiaStartHeight              uint64
		celestiaSearchHeightMaxLookAhead uint64
		precompiles                      map[common.Address]*PrecompileType
	}

	tests := []struct {
		description string
		height      uint64
		checks      testCheck
	}{
		{
			description: "fork1 sets initial values",
			height:      150,
			checks: testCheck{
				stopHeight:                       199,
				minBaseFee:                       1000,
				elasticityMultiplier:             2,
				baseFeeChangeDenominator:         8,
				sequencerChainID:                 "chain1",
				sequencerStartHeight:             100,
				celestiaChainID:                  "celestia1",
				celestiaStartHeight:              1,
				celestiaSearchHeightMaxLookAhead: 100,
				precompiles:                      map[common.Address]*PrecompileType{},
			},
		},
		{
			description: "fork2 inherits everything but EIP1559Params",
			height:      250,
			checks: testCheck{
				stopHeight:                       250,
				minBaseFee:                       2000,
				elasticityMultiplier:             4,
				baseFeeChangeDenominator:         16,
				sequencerChainID:                 "chain1",
				sequencerStartHeight:             299,
				celestiaChainID:                  "celestia1",
				celestiaStartHeight:              1,
				celestiaSearchHeightMaxLookAhead: 100,
				precompiles: map[common.Address]*PrecompileType{
					common.HexToAddress("0x01"): &precompileBase64Val,
				},
			},
		},
		{
			description: "fork3 inherits everything but EIP1559Params",
			height:      251,
			checks: testCheck{
				stopHeight:                       299,
				minBaseFee:                       1000,
				elasticityMultiplier:             4,
				baseFeeChangeDenominator:         16,
				sequencerChainID:                 "chain1",
				sequencerStartHeight:             350,
				celestiaChainID:                  "celestia1",
				celestiaStartHeight:              1,
				celestiaSearchHeightMaxLookAhead: 100,
				precompiles: map[common.Address]*PrecompileType{
					common.HexToAddress("0x01"): &precompileBase64Val,
				},
			},
		},
		{
			description: "fork4 inherits EIP1559Params but changes sequencer",
			height:      350,
			checks: testCheck{
				stopHeight:                       0,
				minBaseFee:                       1000,
				elasticityMultiplier:             4,
				baseFeeChangeDenominator:         16,
				sequencerChainID:                 "chain3",
				sequencerStartHeight:             325,
				celestiaChainID:                  "celestia1",
				celestiaStartHeight:              1,
				celestiaSearchHeightMaxLookAhead: 100,
				precompiles: map[common.Address]*PrecompileType{
					common.HexToAddress("0x01"): &precompileBase64Val,
					common.HexToAddress("0x02"): &precompileBase64Val,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			fork := forks.GetForkAtHeight(test.height)

			if got := fork.StopHeight; got != test.checks.stopHeight {
				t.Errorf("StopHeight = %v, want %v", got, test.checks.stopHeight)
			}
			if got := fork.EIP1559Params.MinBaseFee; got != test.checks.minBaseFee {
				t.Errorf("MinBaseFee = %v, want %v", got, test.checks.minBaseFee)
			}
			if got := fork.EIP1559Params.ElasticityMultiplier; got != test.checks.elasticityMultiplier {
				t.Errorf("ElasticityMultiplier = %v, want %v", got, test.checks.elasticityMultiplier)
			}
			if got := fork.EIP1559Params.BaseFeeChangeDenominator; got != test.checks.baseFeeChangeDenominator {
				t.Errorf("BaseFeeChangeDenominator = %v, want %v", got, test.checks.baseFeeChangeDenominator)
			}
			if got := fork.Sequencer.ChainID; got != test.checks.sequencerChainID {
				t.Errorf("Sequencer.ChainID = %v, want %v", got, test.checks.sequencerChainID)
			}
			if got := fork.Sequencer.StartHeight; got != test.checks.sequencerStartHeight {
				t.Errorf("Sequencer.StartHeight = %v, want %v", got, test.checks.sequencerStartHeight)
			}
			if got := fork.Celestia.ChainID; got != test.checks.celestiaChainID {
				t.Errorf("Celestia.ChainID = %v, want %v", got, test.checks.celestiaChainID)
			}
			if got := fork.Celestia.StartHeight; got != test.checks.celestiaStartHeight {
				t.Errorf("Celestia.StartHeight = %v, want %v", got, test.checks.celestiaStartHeight)
			}
			if got := fork.Celestia.SearchHeightMaxLookAhead; got != test.checks.celestiaSearchHeightMaxLookAhead {
				t.Errorf("Celestia.SearchHeightMaxLookAhead = %v, want %v", got, test.checks.celestiaSearchHeightMaxLookAhead)
			}
			if !reflect.DeepEqual(fork.Precompiles, test.checks.precompiles) {
				t.Errorf("Precompiles = %v, want %v", fork.Precompiles, test.checks.precompiles)
			}
		})
	}
}
