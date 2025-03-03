package caculator_client

import (
	"context"
	"io"
	"log"
	"time"

	"grpc-funnzy-dev/pbs"
)

func (caculatorClient *CaculatorClient) FindMax() {
	log.Println("CaculatorClient.FindMax()")
	stream, err := caculatorClient.client.FindMax(context.Background())
	if err != nil {
		log.Fatalf("caculatorClient.client.FindMax() error: %v", err)
	}

	wait_chan := make(chan struct{})

	go func() {
		for i := 1; i <= 10; i++ {
			findMaxRequest := &pbs.FindMaxRequest{
				Num: int32(i),
			}
			log.Println("findMaxRequest")
			log.Printf("\tNum: %d\n", findMaxRequest.GetNum())
			stream.Send(findMaxRequest)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			findMaxResponse, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					log.Println("EOF\n")
					break
				}
				log.Fatalf("stream.Recv() error: %v", err)
				break
			}

			log.Println("findMaxResponse")
			log.Printf("\tMax: %d\n", findMaxResponse.GetMax())
		}
		close(wait_chan)
	}()
	<-wait_chan
}
