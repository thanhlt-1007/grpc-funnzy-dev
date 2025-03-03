package caculator_service

import (
	"io"
	"log"

	"grpc-funnzy-dev/pbs"
)

func (caculatorService *CaculatorService) Average(
	stream pbs.CaculatorService_AverageServer,
) error {
	log.Println("CaculatorService.Average")

	var sum int32 = 0
	count := 0
	for {
		averageRequest, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("EOF\n")
				averageResponse := &pbs.AverageResponse{
					Average: float32(sum) / float32(count),
				}
				return stream.SendAndClose(averageResponse)
			}

			log.Fatalf("stream.Recv() error: %v", err)
		}

		num := averageRequest.GetNum()
		log.Println("averageRequest")
		log.Printf("\tNum: %d\n", num)

		sum += num
		count++
		log.Printf("\tsum: %d\n", sum)
		log.Printf("\tcount: %d\n", count)
	}
	log.Println("DONE\n")

	return nil
}
