package caculator_client

import (
	"context"
	"log"
	"time"

	"grpc-funnzy-dev/pbs"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (caculatorClient *CaculatorClient) SumWithDeadline() {
	log.Println("CaculatorClient.SumWithDeadline()")
	sumRequest := &pbs.SumRequest{
		Num1: 1,
		Num2: 2,
	}
	log.Println("sumRequest")
	log.Printf("\tNum1: %d\n", sumRequest.GetNum1())
	log.Printf("\tNum2: %d\n", sumRequest.GetNum2())

	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()

	sumResponse, err := caculatorClient.client.SumWithDeadline(ctx, sumRequest)
	if err != nil {
		errStatus, ok := status.FromError(err)
		if ok {
			if errStatus.Code() == codes.DeadlineExceeded {
				log.Fatal("DeadlineExceeded\n")
			}
		}
		log.Fatalf("CaculatorClient.SumWithDeadline() error: %v\n", err)
	}
	log.Println("sumResponse")
	log.Printf("\tSum: %d\n", sumResponse.GetSum())
	log.Println("DONE\n")
}
