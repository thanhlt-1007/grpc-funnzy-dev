package main

import (
	"log"
	"net"

	"grpc-funnzy-dev/grpc/services/caculator_service"
	"grpc-funnzy-dev/pbs"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("net.Listen error: %v\n", err)
	}
	log.Println("net.Listen success")
	log.Println("listener")
	log.Printf("\tAddr: %s\n", listener.Addr().String())

	grpcServer := grpc.NewServer()
	log.Println("grpc.NewServer success")

	pbs.RegisterCaculatorServiceServer(grpcServer, caculator_service.NewCaculatorService())

	log.Printf("start grpcServer.Serve on: %s\n\n", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve error: %v\n", err)
	}
}
