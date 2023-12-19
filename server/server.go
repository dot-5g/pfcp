package server

import (
	"fmt"
	"log"

	"github.com/dot-5g/pfcp/network"
)

type HandleHeartbeatRequest func(HeartbeatRequest)

type HandleHeartbeatResponse func(HeartbeatResponse)

type HeartbeatRequest struct{}

type HeartbeatResponse struct{}

type Server struct {
	udpServer                *network.UdpServer
	heartbeatRequestHandler  HandleHeartbeatRequest
	heartbeatResponseHandler HandleHeartbeatResponse
}

func New() *Server {
	fmt.Printf("Hello, world.\n")
	return &Server{
		udpServer: network.NewUdpServer(),
	}
}

func (server *Server) Run(address string) {
	server.udpServer.Run(address)
	log.Printf("Running UDP server")
}

func (server *Server) HeartbeatRequest(handler HandleHeartbeatRequest) {
	log.Printf("Handling HeartbeatRequest")
	server.heartbeatRequestHandler = handler
}

func (server *Server) HeartbeatResponse(handler HandleHeartbeatResponse) {
	server.heartbeatResponseHandler = handler
	log.Printf("Handling HeartbeatResponse")
}
