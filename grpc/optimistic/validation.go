package optimistic

import (
	optimsticPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/bundle/v1alpha1"
	"fmt"
)

func validateStaticExecuteOptimisticBlockRequest(req *optimsticPb.BaseBlock) error {
	if req.Timestamp == nil {
		return fmt.Errorf("Timestamp cannot be nil")
	}
	if len(req.SequencerBlockHash) == 0 {
		return fmt.Errorf("SequencerBlockHash cannot be empty")
	}

	return nil
}
