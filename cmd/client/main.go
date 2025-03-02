package main

import (
	"context"
	"fmt"
	"grpc-funnzy-dev/pbs"
	"log"
	"time"

	"google.golang.org/grpc"
)

func clientHello(caculatorClient pbs.CaculatorServiceClient) {
	message := fmt.Sprintf("Hello from client at %s", time.Now().String())
	caculatorRequest := &pbs.CaculatorRequest{
		Message: message,
	}

	caculatorResponse, err := caculatorClient.Hello(context.Background(), caculatorRequest)
	if err != nil {
		log.Fatalf("caculatorClient.Hello error: %v\n", err)
	}
	log.Println("caculatorClient.Hello success")
	log.Println("caculatorResponse")
	log.Printf("\tMessage: %s\n", caculatorResponse.GetMessage())
}

func clientSum(caculatorClient pbs.CaculatorServiceClient) {
	sumRequest := &pbs.SumRequest{
		Num1: 1,
		Num2: 2,
	}
	sumResponse, err := caculatorClient.Sum(context.Background(), sumRequest)
	if err != nil {
		log.Fatalf("caculatorClient.Sum error: %v\n", err)
	}
	log.Println("caculatorClient.Sum success")
	log.Println("sumResponse")
	log.Printf("\tSum: %d\n", sumResponse.GetSum())
}

func main() {
	grpcClientConn, err := grpc.Dial("0.0.0.0:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial error: %v\n", err)
	}
	defer grpcClientConn.Close()
	log.Println("grpc.Dial success")
	log.Println("tgrpcClientConn")
	log.Printf("\tTarget: %s", grpcClientConn.Target())
	log.Printf("\tState: %s", grpcClientConn.GetState())

	caculatorClient := pbs.NewCaculatorServiceClient(grpcClientConn)
	clientHello(caculatorClient)
	clientSum(caculatorClient)
}
