package caculator_service

import (
	"context"
	"log"
	"math"

	"grpc-funnzy-dev/pbs"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (caculatorService *CaculatorService) Square(
	context context.Context,
	squareRequest *pbs.SquareRequest,
) (*pbs.SquareResponse, error) {
	log.Println("CaculatorService.Square()")

	num := squareRequest.GetNum()

	log.Println("squareRequest")
	log.Printf("\tNum: %d\n", num)

	if num < 0 {
		return nil, status.Error(
			codes.InvalidArgument,
			"num is negative",
		)
	}

	squareResponse := &pbs.SquareResponse{
		Square: math.Sqrt(float64(num)),
	}
	log.Println("squareResponse")
	log.Printf("\tSquare: %f\n", squareResponse.GetSquare())
	log.Println("DONE\n")

	return squareResponse, nil
}
