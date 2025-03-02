package server

import "log"

func (server *Server) Stop() {
	log.Printf("Server.Stop()")

	server.grpcServer.Stop()
	server.listener.Close()
}
