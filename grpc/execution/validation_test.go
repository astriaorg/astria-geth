package execution

import (
	sequencerblockv1alpha1 "buf.build/gen/go/astria/sequencerblock-apis/protocolbuffers/go/astria/sequencerblock/v1alpha1"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
	"math/big"
	"strings"
	"testing"
)

func randomBlobTx() *types.Transaction {
	return types.NewTx(&types.BlobTx{
		Nonce: 1,
		To:    testAddr,
		Value: uint256.NewInt(1000),
		Gas:   1000,
		Data:  []byte("data"),
	})
}

func randomDepositTx() *types.Transaction {
	return types.NewTx(&types.DepositTx{
		From:  testAddr,
		Value: big.NewInt(1000),
		Gas:   1000,
	})
}

func TestSequenceTxValidation(t *testing.T) {

	blobTx, err := randomBlobTx().MarshalBinary()
	if err != nil {
		t.Fatalf("failed to marshal random blob tx: %v", err)
	}

	depositTx, err := randomDepositTx().MarshalBinary()
	if err != nil {
		t.Fatalf("failed to marshal random deposit tx: %v", err)
	}

	tests := []struct {
		description string
		sequencerTx *sequencerblockv1alpha1.RollupData
		// just check if error contains the string since error contains other details
		wantErr string
	}{
		{
			description: "unmarshallable sequencer tx",
			sequencerTx: &sequencerblockv1alpha1.RollupData{
				Value: &sequencerblockv1alpha1.RollupData_SequencedData{
					SequencedData: []byte("blob tx"),
				},
			},
			wantErr: "failed to unmarshal sequenced data into transaction",
		},
		{
			description: "blob type sequence tx",
			sequencerTx: &sequencerblockv1alpha1.RollupData{
				Value: &sequencerblockv1alpha1.RollupData_SequencedData{
					SequencedData: blobTx,
				},
			},
			wantErr: "blob tx not allowed in sequenced data",
		},
		{
			description: "deposit type sequence tx",
			sequencerTx: &sequencerblockv1alpha1.RollupData{
				Value: &sequencerblockv1alpha1.RollupData_SequencedData{
					SequencedData: depositTx,
				},
			},
			wantErr: "deposit tx not allowed in sequenced data",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			_, _, serviceV1Alpha1 := setupExecutionService(t, 10)

			_, err := serviceV1Alpha1.SequencerTxValidation(test.sequencerTx)
			if test.wantErr != "" && err == nil {
				t.Errorf("expected error, got nil")
			}
			if test.wantErr == "" && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			// check if wantErr is in err.Error()
			if test.wantErr != "" && !strings.Contains(err.Error(), test.wantErr) {
				t.Errorf("expected error to contain %q, got %q", test.wantErr, err.Error())
			}
		})
	}
}
