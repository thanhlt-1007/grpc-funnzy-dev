package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func NewServer() *Server {
	log.Println("server.NewServer()")
	server := Server{
		listener:   newListener(),
		grpcServer: newGrpcServer(),
	}
	return &server
}

func newListener() net.Listener {
	log.Println("server.newListener()")
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("ERROR %s\n", err.Error())
	}
	log.Println("\tlistener")
	log.Printf("\t\tAdd: %s", listener.Addr().String())
	log.Printf("\t\tNetwork: %s", listener.Addr().Network())

	return listener
}

func newGrpcServer() *grpc.Server {
	log.Println("server.newGrpcServer()")
	grpcServer := grpc.NewServer()

	return grpcServer
}
