package main

import (
	"grpc-funnzy-dev/pbs"
	"log"
	"net"

	"context"

	"google.golang.org/grpc"
)

type CaculatorService struct {
}

func (caculatorService *CaculatorService) Hello(context.Context, *pbs.CaculatorRequest) (*pbs.CaculatorResponse, error) {
	return nil, nil
}

func NewCaculatorService() *CaculatorService {
	service := CaculatorService{}
	return &service
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("net.Listen err: %v\n", err)
	}
	log.Println("net.Listen success")
	log.Printf("\tlistener: %s\n", listener.Addr().String())

	grpcServer := grpc.NewServer()
	log.Println("grpc.NewServer success")

	caculatorService := NewCaculatorService()
	pbs.RegisterCaculatorServiceServer(grpcServer, caculatorService)

	log.Printf("start grpcServer.Serve on: %s\n", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v\n", err)
	}
}
