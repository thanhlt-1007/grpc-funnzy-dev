package main

import (
	"context"
	"fmt"
	"grpc-funnzy-dev/pbs"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type CaculatorServer struct {
}

func (caculatorServer *CaculatorServer) Hello(contex context.Context, caculatorRequest *pbs.CaculatorRequest) (*pbs.CaculatorResponse, error) {
	log.Println("CaculatorServer.Hello")
	log.Println("caculatorRequest")
	log.Printf("\tMessage: %s\n", caculatorRequest.GetMessage())

	message := fmt.Sprintf("Hello from server at %s", time.Now().String())
	caculatorResponse := &pbs.CaculatorResponse{
		Message: message,
	}
	return caculatorResponse, nil
}

func NewCaculatorService() *CaculatorServer {
	service := CaculatorServer{}
	return &service
}

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

	caculatorServer := NewCaculatorService()
	pbs.RegisterCaculatorServiceServer(grpcServer, caculatorServer)

	log.Printf("start grpcServer.Serve on: %s\n", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve error: %v\n", err)
	}
}
