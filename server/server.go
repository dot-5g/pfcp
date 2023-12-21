package server

import (
	"log"
	"net"

	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/network"
)

const HeaderSize = 8

type HandleHeartbeatRequest func(*messages.HeartbeatRequest)

type HandleHeartbeatResponse func(*messages.HeartbeatResponse)

type PfcpMessage struct {
	Header  messages.PFCPHeader
	Message []byte
}

type Server struct {
	address                  string
	udpServer                *network.UdpServer
	heartbeatRequestHandler  HandleHeartbeatRequest
	heartbeatResponseHandler HandleHeartbeatResponse
}

func New(address string) *Server {
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

	header, err := messages.ParsePFCPHeader(data[:HeaderSize])
	if err != nil {
		log.Printf("Error parsing PFCP header: %v", err)
	}
	pfcpMessage := PfcpMessage{Header: header, Message: data[HeaderSize:]}

	if pfcpMessage.Header.MessageType == 1 {
		recoveryTimeStamp := messages.FromBytes(pfcpMessage.Message)
		heartbeatRequest := messages.HeartbeatRequest{
			RecoveryTimeStamp: recoveryTimeStamp,
		}
		server.heartbeatRequestHandler(&heartbeatRequest)
	}
}

func (server *Server) HeartbeatRequest(handler HandleHeartbeatRequest) {
	server.heartbeatRequestHandler = handler
}

func (server *Server) HeartbeatResponse(handler HandleHeartbeatResponse) {
	server.heartbeatResponseHandler = handler
}
