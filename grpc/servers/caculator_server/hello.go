package caculator_server

import (
	"context"
	"fmt"
	"log"
	"time"

	"grpc-funnzy-dev/pbs"
)

func (caculatorServer *CaculatorServer) Hello(contex context.Context, HelloRequest *pbs.HelloRequest) (*pbs.HelloResponse, error) {
	log.Println("CaculatorServer.Hello")
	log.Println("HelloRequest")
	log.Printf("\tMessage: %s\n", HelloRequest.GetMessage())

	message := fmt.Sprintf("Hello from server at %s", time.Now().String())
	helloResponse := &pbs.HelloResponse{
		Message: message,
	}
	return helloResponse, nil
}
