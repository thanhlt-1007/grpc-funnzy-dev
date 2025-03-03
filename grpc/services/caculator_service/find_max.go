package caculator_service

import (
	"io"
	"log"

	"grpc-funnzy-dev/pbs"
)

func (caculatorService *CaculatorService) FindMax(
	stream pbs.CaculatorService_FindMaxServer,
) error {
	log.Println("CaculatorService.FindMax")

	var max int32 = 0

	for {
		findMaxRequest, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("EOF\n``")
				return nil
			}

			log.Fatalf("stream.Recv() error: %v", err)
		}

		num := findMaxRequest.GetNum()
		log.Println("findMaxRequest")
		log.Printf("\tNum: %d\n", num)

		if num > max {
			max = num
			findMaxResponse := &pbs.FindMaxResponse{
				Max: max,
			}
			log.Println("findMaxResponse")
			log.Printf("\tMax: %d\n", max)
			err := stream.Send(findMaxResponse)
			if err != nil {
				log.Fatalf("stream.Send() error: %v", err)
			}
		}
	}

	return nil
}
