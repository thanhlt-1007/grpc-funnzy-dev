package caculator_server

import (
	"context"
	"log"

	"grpc-funnzy-dev/pbs"
)

func (caculatorServer *CaculatorServer) Sum(context context.Context, sumRequest *pbs.SumRequest) (*pbs.SumResponse, error) {
	log.Println("CaculatorServer.Sum")

	num1 := sumRequest.Num1
	num2 := sumRequest.Num2

	log.Println("sumRequest")
	log.Printf("\tNum1: %d\n", num1)
	log.Printf("\tNum2: %d\n", num2)

	sum := num1 + num2
	sumResponse := &pbs.SumResponse{
		Sum: sum,
	}

	log.Println("sumResponse")
	log.Printf("\tSum: %d\n", sumResponse.GetSum())
	log.Println("DONE\n")

	return sumResponse, nil
}
