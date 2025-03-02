package main

import (
	"grpc-funnzy-dev/grpc/server"
)

func main() {
	server := server.NewServer()
	server.RegisterServices()
	server.Serve()
	defer server.Stop()
}
