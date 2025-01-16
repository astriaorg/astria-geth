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

func TestAstriaEIP1559Params(t *testing.T) {
	astriaForks, _ := NewAstriaForks(map[string]AstriaForkConfig{
		"genesis": {
			Height: 1,
			Sequencer: &AstriaSequencerConfig{
				ChainID:     "sequencer-test-chain-0",
				StartHeight: 2,
			},
			Celestia: &AstriaCelestiaConfig{
				ChainID:        "mocha-4",
				StartHeight:    2,
				HeightVariance: 10,
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
				StartHeight:    2,
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
				StartHeight:    2,
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
				StartHeight:    2,
				AssetDenom:     "nria",
				AssetPrecision: 18,
				Erc20Asset:     nil,
			},
			wantErr: fmt.Errorf("bridge address must have resolve to 20 byte address, got 19"),
		},
		{
			description: "invalid start height",
			config: AstriaBridgeAddressConfig{
				BridgeAddress:  bridgeAddressBech32,
				StartHeight:    0,
				AssetDenom:     "nria",
				AssetPrecision: 18,
				Erc20Asset:     nil,
			},
			wantErr: fmt.Errorf("start height must be greater than 0"),
		},
		{
			description: "invalid asset denom",
			config: AstriaBridgeAddressConfig{
				BridgeAddress:  bridgeAddressBech32,
				StartHeight:    2,
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
				StartHeight:    2,
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
				StartHeight:    2,
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
				StartHeight:    2,
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
				StartHeight:    2,
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
				ChainID:        "celestia1",
				StartHeight:    1,
				HeightVariance: 100,
			},
		},
		"fork2": {
			Height: 1000,
			Sequencer: &AstriaSequencerConfig{
				ChainID:     "chain2",
				StartHeight: 2,
			},
			Celestia: &AstriaCelestiaConfig{
				ChainID:        "celestia2",
				StartHeight:    2,
				HeightVariance: 200,
			},
		},
		"fork3": {
			Height: 2000,
			Sequencer: &AstriaSequencerConfig{
				ChainID:     "chain3",
				StartHeight: 3,
			},
			Celestia: &AstriaCelestiaConfig{
				ChainID:        "celestia3",
				StartHeight:    3,
				HeightVariance: 300,
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
				ChainID:        "celestia1",
				StartHeight:    1,
				HeightVariance: 100,
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
				StartHeight: 1,
			},
			Celestia: &AstriaCelestiaConfig{
				ChainID:        "celestia1",
				StartHeight:    1,
				HeightVariance: 100,
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
		},
		"fork3": {
			Height: 300,
			// override sequencer config
			Sequencer: &AstriaSequencerConfig{
				ChainID:     "chain3",
				StartHeight: 3,
			},
			// EIP1559Params should be inherited from fork2
		},
	}

	forks, err := NewAstriaForks(forkMap)
	if err != nil {
		t.Fatalf("failed to create forks: %v", err)
	}

	type testCheck struct {
		minBaseFee               uint64
		elasticityMultiplier     uint64
		baseFeeChangeDenominator uint64
		sequencerChainID         string
		sequencerStartHeight     uint32
		celestiaChainID          string
		celestiaStartHeight      uint64
		celestiaHeightVariance   uint64
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
				minBaseFee:               1000,
				elasticityMultiplier:     2,
				baseFeeChangeDenominator: 8,
				sequencerChainID:         "chain1",
				sequencerStartHeight:     1,
				celestiaChainID:          "celestia1",
				celestiaStartHeight:      1,
				celestiaHeightVariance:   100,
			},
		},
		{
			description: "fork2 inherits everything but EIP1559Params",
			height:      250,
			checks: testCheck{
				minBaseFee:               2000,
				elasticityMultiplier:     4,
				baseFeeChangeDenominator: 16,
				sequencerChainID:         "chain1",
				sequencerStartHeight:     1,
				celestiaChainID:          "celestia1",
				celestiaStartHeight:      1,
				celestiaHeightVariance:   100,
			},
		},
		{
			description: "fork3 inherits EIP1559Params but changes sequencer",
			height:      350,
			checks: testCheck{
				minBaseFee:               2000,
				elasticityMultiplier:     4,
				baseFeeChangeDenominator: 16,
				sequencerChainID:         "chain3",
				sequencerStartHeight:     3,
				celestiaChainID:          "celestia1",
				celestiaStartHeight:      1,
				celestiaHeightVariance:   100,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			fork := forks.GetForkAtHeight(test.height)

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
			if got := fork.Celestia.HeightVariance; got != test.checks.celestiaHeightVariance {
				t.Errorf("Celestia.HeightVariance = %v, want %v", got, test.checks.celestiaHeightVariance)
			}
		})
	}
}
