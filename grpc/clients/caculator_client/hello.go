package caculator_client

import (
	"context"
	"fmt"
	"log"
	"time"

	"grpc-funnzy-dev/pbs"
)

func (caculatorClient *CaculatorClient) Hello() {
	log.Println("CaculatorClient.Hello()")
	message := fmt.Sprintf("Hello from client at %s", time.Now().String())
	helloRequest := &pbs.HelloRequest{
		Message: message,
	}
	log.Println("helloRequest")
	log.Printf("\tMessage: %s\n", helloRequest.GetMessage())

	helloResponse, err := caculatorClient.client.Hello(context.Background(), helloRequest)
	if err != nil {
		log.Fatalf("CaculatorClient.Hello error: %v\n", err)
	}
	log.Println("helloResponse")
	log.Printf("\tMessage: %s\n", helloResponse.GetMessage())
	log.Println("DONE\n")
}
