package caculator_client

import (
	"context"
	"io"
	"log"

	"grpc-funnzy-dev/pbs"
)

func (caculatorClient *CaculatorClient) ToPirmeNumber() {
	log.Println("CaculatorClient.ToPirmeNumber()")
	toPrimeNumberRequest := &pbs.ToPrimeNumberRequest{
		Number: 120,
	}
	log.Println("toPrimeNumberRequest")
	log.Printf("\tNumber: %d\n", toPrimeNumberRequest.GetNumber())

	stream, err := caculatorClient.client.ToPrimeNumber(context.Background(), toPrimeNumberRequest)
	if err != nil {
		log.Fatalf("CaculatorClient.ToPirmeNumber() error: %v\n", err)
	}

	for {
		toPrimeNumberResponse, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("EOF\n")
				return
			} else {
				log.Fatalf("stream.Recv() error: %v\n", err)
			}
		}

		log.Println("toPrimeNumberResponse")
		log.Printf("\tPrimeNumber: %d\n", toPrimeNumberResponse.GetPrimeNumber())
	}
}
