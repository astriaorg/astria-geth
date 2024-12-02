package execution

import (
	astriaPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/execution/v1alpha2"
	"fmt"
)

// `validateStaticExecuteBlockRequest` validates the given execute block request without regard
// to the current state of the system. This is useful for validating the request before any
// state changes or reads are made as a basic guard.
func validateStaticExecuteBlockRequest(req *astriaPb.ExecuteBlockRequest) error {
	if req.PrevBlockHash == nil {
		return fmt.Errorf("PrevBlockHash cannot be nil")
	}
	if req.Timestamp == nil {
		return fmt.Errorf("Timestamp cannot be nil")
	}

	return nil
}

// `validateStaticCommitment` validates the given commitment without regard to the current state of the system.
func validateStaticCommitmentState(commitmentState *astriaPb.CommitmentState) error {
	if commitmentState == nil {
		return fmt.Errorf("commitment state is nil")
	}
	if commitmentState.Soft == nil {
		return fmt.Errorf("soft block is nil")
	}
	if commitmentState.Firm == nil {
		return fmt.Errorf("firm block is nil")
	}
	if commitmentState.BaseCelestiaHeight == 0 {
		return fmt.Errorf("base celestia height of 0 is not valid")
	}

	if err := validateStaticBlock(commitmentState.Soft); err != nil {
		return fmt.Errorf("soft block invalid: %w", err)
	}
	if err := validateStaticBlock(commitmentState.Firm); err != nil {
		return fmt.Errorf("firm block invalid: %w", err)
	}

	return nil
}

// `validateStaticBlock` validates the given block as a  without regard to the current state of the system.
func validateStaticBlock(block *astriaPb.Block) error {
	if block.ParentBlockHash == nil {
		return fmt.Errorf("parent block hash is nil")
	}
	if block.Hash == nil {
		return fmt.Errorf("block hash is nil")
	}
	if block.Timestamp == nil {
		return fmt.Errorf("timestamp is 0")
	}

	return nil
}
