package node

import (
	optimisticGrpc "buf.build/gen/go/astria/execution-apis/grpc/go/astria/bundle/v1alpha1/bundlev1alpha1grpc"
	"net"
	"os"
	"sync"

	astriaGrpc "buf.build/gen/go/astria/execution-apis/grpc/go/astria/execution/v1/executionv1grpc"
	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/grpc"
)

// GRPCServerHandler is the gRPC server handler.
// It gives us a way to attach the gRPC server to the node so it can be stopped on shutdown.
type GRPCServerHandler struct {
	mu sync.Mutex

	tcpEndpoint                string
	udsEndpoint                string
	execServer                 *grpc.Server
	optimisticServer           *grpc.Server
	executionServiceServerV1a2 *astriaGrpc.ExecutionServiceServer
	optimisticExecServ         *optimisticGrpc.OptimisticExecutionServiceServer
	streamBundleServ           *optimisticGrpc.BundleServiceServer
}

// NewServer creates a new gRPC server.
// It registers the execution service server.
// It registers the gRPC server with the node so it can be stopped on shutdown.
func NewGRPCServerHandler(node *Node, execServ astriaGrpc.ExecutionServiceServer, optimisticExecServ optimisticGrpc.OptimisticExecutionServiceServer, streamBundleServ optimisticGrpc.BundleServiceServer, cfg *Config) error {
	execServer, optimisticServer := grpc.NewServer(), grpc.NewServer()

	log.Info("gRPC server enabled", "tcpEndpoint", cfg.GRPCTcpEndpoint(), "udsEndpoint", cfg.GRPCUdsEndpoint())

	serverHandler := &GRPCServerHandler{
		tcpEndpoint:                cfg.GRPCTcpEndpoint(),
		udsEndpoint:                cfg.GRPCUdsEndpoint(),
		execServer:                 execServer,
		optimisticServer:           optimisticServer,
		executionServiceServerV1a2: &execServ,
		optimisticExecServ:         &optimisticExecServ,
		streamBundleServ:           &streamBundleServ,
	}

	astriaGrpc.RegisterExecutionServiceServer(execServer, execServ)
	optimisticGrpc.RegisterOptimisticExecutionServiceServer(optimisticServer, optimisticExecServ)
	optimisticGrpc.RegisterBundleServiceServer(optimisticServer, streamBundleServ)

	node.RegisterGRPCServer(serverHandler)
	return nil
}

// Start starts the gRPC server if it is enabled.
func (handler *GRPCServerHandler) Start() error {
	handler.mu.Lock()
	defer handler.mu.Unlock()

	if handler.tcpEndpoint == "" {
		return nil
	}
	if handler.udsEndpoint == "" {
		return nil
	}

	// Start the gRPC server
	tcpLis, err := net.Listen("tcp", handler.tcpEndpoint)
	if err != nil {
		return err
	}

	// Remove any existing socket file
	if err := os.RemoveAll(handler.udsEndpoint); err != nil {
		return err
	}
	udsLis, err := net.Listen("unix", handler.udsEndpoint)
	if err != nil {
		return err
	}

	go handler.execServer.Serve(tcpLis)
	go handler.optimisticServer.Serve(udsLis)
	log.Info("gRPC server started", "tcpEndpoint", handler.tcpEndpoint, "udsEndpoint", handler.udsEndpoint)
	return nil
}

// Stop stops the gRPC server.
func (handler *GRPCServerHandler) Stop() error {
	handler.mu.Lock()
	defer handler.mu.Unlock()

	handler.execServer.GracefulStop()
	handler.optimisticServer.GracefulStop()
	log.Info("gRPC server stopped", "tcpEndpoint", handler.tcpEndpoint, "udsEndpoint", handler.udsEndpoint)
	return nil
}
