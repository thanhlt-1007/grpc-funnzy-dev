package caculator_client

import (
	"context"
	"log"

	"grpc-funnzy-dev/pbs"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (caculatorClient *CaculatorClient) Square(num int32) {
	log.Println("CaculatorClient.Square()")

	squareRequest := &pbs.SquareRequest{
		Num: num,
	}
	log.Println("squareRequest")
	log.Printf("\tNum: %d\n", num)

	squareResponse, err := caculatorClient.client.Square(context.Background(), squareRequest)
	if err != nil {
		stt, ok := status.FromError(err)
		if ok {
			if stt.Code() == codes.InvalidArgument {
				log.Println("InvalidArgument\n")
				return
			}
		}
	}
	log.Println("squareResponse")
	log.Printf("\tSquare: %f\n", squareResponse.GetSquare())
	log.Println("DONE\n")
}
