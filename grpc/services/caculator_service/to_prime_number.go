package caculator_service

import (
	"log"
	"time"

	"grpc-funnzy-dev/pbs"
)

func (caculatorService *CaculatorService) ToPrimeNumber(
	toPrimeNumberRequest *pbs.ToPrimeNumberRequest,
	stream pbs.CaculatorService_ToPrimeNumberServer,
) error {
	log.Println("CaculatorService.ToPrimeNumber")

	log.Println("toPrimeNumberRequest")
	log.Printf("\tNumber: %d\n", toPrimeNumberRequest.GetNumber())

	k := int32(2)
	n := toPrimeNumberRequest.GetNumber()
	for n > 1 {
		if n%k == 0 {
			n = n / k
			toPrimeNumberResponse := &pbs.ToPrimeNumberResponse{
				PrimeNumber: k,
			}
			log.Printf("toPrimeNumberResponse")
			log.Printf("\tPrimeNumber: %d\n", toPrimeNumberResponse.GetPrimeNumber())
			log.Printf("\tremain: %d\n", n)
			time.Sleep(1 * time.Second)
			stream.Send(toPrimeNumberResponse)
		} else {
			k++
		}
	}
	log.Println("DONE\n")

	return nil
}
