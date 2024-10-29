package execution

import (
	"github.com/ethereum/go-ethereum/grpc/shared"
	"testing"
)

func SetupExecutionService(t *testing.T, sharedService *shared.SharedServiceContainer) *ExecutionServiceServerV1 {
	t.Helper()

	return NewExecutionServiceServerV1(sharedService)
}
