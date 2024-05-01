package params

import (
	"encoding/json"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestAstriaEIP1559Params(t *testing.T) {
	jsonBuf := []byte(`{
		"1":{ "minBaseFee": 45000000000, "elasticityMultiplier": 4, "baseFeeChangeDenominator": 100 },
		"101":{ "minBaseFee": 120000000, "elasticityMultiplier": 11, "baseFeeChangeDenominator": 250 },
		"15":{ "minBaseFee": 15000000000, "elasticityMultiplier": 5, "baseFeeChangeDenominator": 50 }
	}`)

	var eip1559Params AstriaEIP1559Params
	err := json.Unmarshal(jsonBuf, &eip1559Params)
	if err != nil {
		t.Errorf("unexpected err %v", err)
	}

	expected := AstriaEIP1559Params{
		heights: map[uint64]AstriaEIP1559Param{
			1:   {MinBaseFee: 45000000000, ElasticityMultiplier: 4, BaseFeeChangeDenominator: 100},
			101: {MinBaseFee: 120000000, ElasticityMultiplier: 11, BaseFeeChangeDenominator: 250},
			15:  {MinBaseFee: 15000000000, ElasticityMultiplier: 5, BaseFeeChangeDenominator: 50},
		},
		orderedHeights: []uint64{101, 15, 1},
	}

	if !reflect.DeepEqual(eip1559Params, expected) {
		t.Errorf("expected %v, got %v", expected, eip1559Params)
	}

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
		if got := eip1559Params.MinBaseFeeAt(height); got.Cmp(expected) != 0 {
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
		if got := eip1559Params.ElasticityMultiplierAt(height); got != expected {
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
		if got := eip1559Params.BaseFeeChangeDenominatorAt(height); got != expected {
			t.Errorf("BaseFeeChangeDenominatorAt(%d): expected %v, got %v", height, expected, got)
		}
	}
}
