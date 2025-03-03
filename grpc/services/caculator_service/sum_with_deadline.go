package caculator_service

import (
	"context"
	"log"
	"time"

	"grpc-funnzy-dev/pbs"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (caculatorService *CaculatorService) SumWithDeadline(ctx context.Context, sumRequest *pbs.SumRequest) (*pbs.SumResponse, error) {
	log.Println("CaculatorService.Sum()")

	time.Sleep(1 * time.Second)

	num1 := sumRequest.Num1
	num2 := sumRequest.Num2

	time.Sleep(1 * time.Second)

	log.Println("sumRequest")
	log.Printf("\tNum1: %d\n", num1)
	log.Printf("\tNum2: %d\n", num2)

	time.Sleep(1 * time.Second)

	sum := num1 + num2

	time.Sleep(1 * time.Second)

	sumResponse := &pbs.SumResponse{
		Sum: sum,
	}

	time.Sleep(1 * time.Second)

	if ctx.Err() == context.DeadlineExceeded {
		log.Println("DeadlineExceeded\n")
		return nil, status.Errorf(codes.Canceled, "client Canceled")
	}

	log.Println("sumResponse")
	log.Printf("\tSum: %d\n", sumResponse.GetSum())
	log.Println("DONE\n")

	return sumResponse, nil
}
