package main

import (
	"log"

	"google.golang.org/grpc"
	"grpc-funnzy-dev/grpc/clients/caculator_client"
)

func newGrpcClientConn() *grpc.ClientConn {
	log.Println("newGrpcClientConn()")
	grpcClientConn, err := grpc.Dial("0.0.0.0:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial error: %v\n", err)
	}

	log.Println("\tgrpcClientConn")
	log.Printf("\t\tTarget: %s", grpcClientConn.Target())
	log.Printf("\t\tState: %s", grpcClientConn.GetState())

	return grpcClientConn
}

func main() {
	grpcClientConn := newGrpcClientConn()

	grpc_caculator_client := caculator_client.NewCaculatorClient(grpcClientConn)
	grpc_caculator_client.Hello()
	grpc_caculator_client.Sum()
	grpc_caculator_client.ToPirmeNumber()
	grpc_caculator_client.Average()

	defer grpcClientConn.Close()
}
