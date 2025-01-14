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
