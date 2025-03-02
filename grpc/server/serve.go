package server

import "log"

func (server *Server) Serve() {
	log.Printf("Server.Serve() at %s\n\n", server.listener.Addr().String())
	err := server.grpcServer.Serve(server.listener)
	if err != nil {
		log.Fatalf("server.grpcServer.Serve error: %v", err)
	}
	defer server.grpcServer.Stop()
	defer server.listener.Close()
}
