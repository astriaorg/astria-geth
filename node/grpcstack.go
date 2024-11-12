package node

import (
	optimisticGrpc "buf.build/gen/go/astria/execution-apis/grpc/go/astria/bundle/v1alpha1/bundlev1alpha1grpc"
	"net"
	"sync"

	astriaGrpc "buf.build/gen/go/astria/execution-apis/grpc/go/astria/execution/v1/executionv1grpc"
	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/grpc"
)

// GRPCServerHandler is the gRPC server handler.
// It gives us a way to attach the gRPC server to the node so it can be stopped on shutdown.
type GRPCServerHandler struct {
	mu sync.Mutex

	endpoint                   string
	execServer                 *grpc.Server
	executionServiceServerV1a2 *astriaGrpc.ExecutionServiceServer
	optimisticExecServ         *optimisticGrpc.OptimisticExecutionServiceServer
	streamBundleServ           *optimisticGrpc.BundleServiceServer

	enableAuctioneer bool
}

// NewServer creates a new gRPC server.
// It registers the execution service server.
// It registers the gRPC server with the node so it can be stopped on shutdown.
func NewGRPCServerHandler(node *Node, execServ astriaGrpc.ExecutionServiceServer, optimisticExecServ optimisticGrpc.OptimisticExecutionServiceServer, streamBundleServ optimisticGrpc.BundleServiceServer, cfg *Config) error {
	execServer := grpc.NewServer()

	log.Info("gRPC server enabled", "endpoint", cfg.GRPCEndpoint())

	serverHandler := &GRPCServerHandler{
		endpoint:                   cfg.GRPCEndpoint(),
		execServer:                 execServer,
		executionServiceServerV1a2: &execServ,
		optimisticExecServ:         &optimisticExecServ,
		streamBundleServ:           &streamBundleServ,
		enableAuctioneer:           cfg.EnableAuctioneer,
	}

	astriaGrpc.RegisterExecutionServiceServer(execServer, execServ)
	if cfg.EnableAuctioneer {
		optimisticGrpc.RegisterOptimisticExecutionServiceServer(execServer, optimisticExecServ)
		optimisticGrpc.RegisterBundleServiceServer(execServer, streamBundleServ)
	}

	node.RegisterGRPCServer(serverHandler)
	return nil
}

// Start starts the gRPC server if it is enabled.
func (handler *GRPCServerHandler) Start() error {
	handler.mu.Lock()
	defer handler.mu.Unlock()

	if handler.endpoint == "" {
		return nil
	}

	// Start the gRPC server
	tcpLis, err := net.Listen("tcp", handler.endpoint)
	if err != nil {
		return err
	}

	go handler.execServer.Serve(tcpLis)

	log.Info("gRPC server started", "endpoint", handler.endpoint)
	return nil
}

// Stop stops the gRPC server.
func (handler *GRPCServerHandler) Stop() error {
	handler.mu.Lock()
	defer handler.mu.Unlock()

	handler.execServer.GracefulStop()

	log.Info("gRPC server stopped", "endpoint", handler.endpoint)
	return nil
}
