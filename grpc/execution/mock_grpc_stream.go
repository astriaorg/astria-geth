package execution

import (
	optimsticPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/bundle/v1alpha1"
	"context"
	"google.golang.org/grpc/metadata"
	"io"
	"time"
)

type MockStream struct {
	requestStream        []*optimsticPb.StreamExecuteOptimisticBlockRequest
	accumulatedResponses []*optimsticPb.StreamExecuteOptimisticBlockResponse
	requestCounter       uint64
}

func (ms *MockStream) Recv() (*optimsticPb.StreamExecuteOptimisticBlockRequest, error) {
	// add a delay to make it look like an async stream
	time.Sleep(2 * time.Second)
	if ms.requestCounter > uint64(len(ms.requestStream)-1) {
		// end the stream after all the packets have been sent
		return nil, io.EOF
	}

	req := ms.requestStream[ms.requestCounter]
	ms.requestCounter += 1

	return req, nil
}

func (ms *MockStream) Send(res *optimsticPb.StreamExecuteOptimisticBlockResponse) error {
	ms.accumulatedResponses = append(ms.accumulatedResponses, res)
	return nil
}

func (ms *MockStream) SetHeader(md metadata.MD) error {
	panic("implement me")
}

func (ms *MockStream) SendHeader(md metadata.MD) error {
	panic("implement me")
}

func (ms *MockStream) SetTrailer(md metadata.MD) {
	panic("implement me")
}

func (ms *MockStream) Context() context.Context {
	return context.Background()
}

func (ms *MockStream) SendMsg(m any) error {
	panic("implement me")
}

func (ms *MockStream) RecvMsg(m any) error {
	panic("implement me")
}
