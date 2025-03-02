package server

import (
	"log"

	"grpc-funnzy-dev/grpc/services/caculator_service"
	"grpc-funnzy-dev/pbs"
)

func (server *Server) RegisterServices() {
	log.Println("Server.RegisterServices()")
	pbs.RegisterCaculatorServiceServer(server.grpcServer, caculator_service.NewCaculatorService())
	log.Println()
}
