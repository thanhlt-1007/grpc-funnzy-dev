package caculator_client

import (
	"grpc-funnzy-dev/pbs"

	"google.golang.org/grpc"
)

func NewCaculatorClient(grpcClientConn *grpc.ClientConn) *CaculatorClient {
	client := pbs.NewCaculatorServiceClient(grpcClientConn)
	caculatorClient := CaculatorClient{
		client: client,
	}

	return &caculatorClient
}
