package caculator_client

import (
	"context"
	"log"
	"time"

	"grpc-funnzy-dev/pbs"
)

func (caculatorClient *CaculatorClient) Average() {
	log.Println("CaculatorClient.Average()")
	stream, err := caculatorClient.client.Average(context.Background())

	for i := 1; i <= 10; i++ {
		averageRequest := &pbs.AverageRequest{
			Num: int32(i),
		}
		log.Println("averageRequest")
		log.Printf("\tNum: %d\n", averageRequest.GetNum())
		stream.Send(averageRequest)
		time.Sleep(1 * time.Second)
	}

	averageResponse, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("CaculatorClient.Average() error: %v\n", err)
	}
	log.Println("averageResponse")
	log.Printf("\tAverage: %f\n", averageResponse.GetAverage())
	log.Println("DONE\n")
}
