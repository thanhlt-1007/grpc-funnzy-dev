package caculator_client

import (
	"grpc-funnzy-dev/pbs"
)

type CaculatorClient struct {
	client pbs.CaculatorServiceClient
}
