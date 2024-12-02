package optimistic

import (
	"github.com/ethereum/go-ethereum/grpc/shared"
	"testing"
)

func SetupOptimisticService(t *testing.T, sharedService *shared.SharedServiceContainer) *OptimisticServiceV1Alpha1 {
	t.Helper()

	return NewOptimisticServiceV1Alpha(sharedService)
}
