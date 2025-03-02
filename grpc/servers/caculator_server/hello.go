package caculator_server

import (
	"context"
	"fmt"
	"log"
	"time"

	"grpc-funnzy-dev/pbs"
)

func (caculatorServer *CaculatorServer) Hello(contex context.Context, helloRequest *pbs.HelloRequest) (*pbs.HelloResponse, error) {
	log.Println("\nCaculatorServer.Hello")

	message := helloRequest.GetMessage()

	log.Println("helloRequest")
	log.Printf("\tMessage: %s\n", message)

	message = fmt.Sprintf("Hello from server at %s", time.Now().String())
	helloResponse := &pbs.HelloResponse{
		Message: message,
	}

	log.Println("helloResponse")
	log.Printf("\tMessage: %s\n", helloResponse.GetMessage())
	log.Println("DONE\n")

	return helloResponse, nil
}
