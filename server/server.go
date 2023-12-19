package server

import (
	"fmt"
	"net"

	"github.com/dot-5g/pfcp/network"
)

type HandleHeartbeatRequest func(HeartbeatRequest)

type HandleHeartbeatResponse func(HeartbeatResponse)

type HeartbeatRequest struct{}

type HeartbeatResponse struct{}

type Server struct {
	address                  string
	udpServer                *network.UdpServer
	heartbeatRequestHandler  HandleHeartbeatRequest
	heartbeatResponseHandler HandleHeartbeatResponse
}

func New(address string) *Server {
	fmt.Printf("Hello, world.\n")
	return &Server{
		address:   address,
		udpServer: network.NewUdpServer(),
	}
}

func (server *Server) Run() {
	server.udpServer.SetHandler(server.handleUDPMessage)
	server.udpServer.Run(server.address)
}

func (server *Server) handleUDPMessage(data []byte, addr net.Addr) {
	if isHeartbeatRequest(data) {
		heartbeatRequest := HeartbeatRequest{}
		if server.heartbeatRequestHandler != nil {
			server.heartbeatRequestHandler(heartbeatRequest)
		}
	}
}

func (server *Server) HeartbeatRequest(handler HandleHeartbeatRequest) {
	server.heartbeatRequestHandler = handler
}

func (server *Server) HeartbeatResponse(handler HandleHeartbeatResponse) {
	server.heartbeatResponseHandler = handler
}

func isHeartbeatRequest(data []byte) bool {
	if len(data) < 12 { // 12 bytes is the size of PFCPHeader
		return false
	}

	messageType := data[1]

	return messageType == 1
}
