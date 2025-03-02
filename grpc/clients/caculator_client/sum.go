package caculator_client

import (
	"context"
	"log"

	"grpc-funnzy-dev/pbs"
)

func (caculatorClient *CaculatorClient) Sum() {
	log.Println("CaculatorClient.Sum()")
	sumRequest := &pbs.SumRequest{
		Num1: 1,
		Num2: 2,
	}
	log.Println("sumRequest")
	log.Printf("\tNum1: %d\n", sumRequest.GetNum1())
	log.Printf("\tNum2: %d\n", sumRequest.GetNum2())

	sumResponse, err := caculatorClient.client.Sum(context.Background(), sumRequest)
	if err != nil {
		log.Fatalf("CaculatorClient.Sum() error: %v\n", err)
	}
	log.Println("sumResponse")
	log.Printf("\tSum: %d\n", sumResponse.GetSum())
	log.Println("DONE\n")
}
