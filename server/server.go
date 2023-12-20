package server

import (
	"encoding/binary"
	"log"
	"net"
	"time"

	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/network"
)

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

	pfcpMessage := ParseUDPMessage(data)

	if pfcpMessage.Header.MessageType == 1 {
		timestampBytes := pfcpMessage.Message

		if len(timestampBytes) >= 4 {
			timestamp := binary.BigEndian.Uint32(timestampBytes)
			recoveryTime := time.Unix(int64(timestamp), 0)

			heartbeatRequest := messages.HeartbeatRequest{
				RecoveryTimeStamp: messages.RecoveryTimeStamp(recoveryTime),
			}

			if server.heartbeatRequestHandler != nil {
				server.heartbeatRequestHandler(&heartbeatRequest)
			}
		} else {
			log.Printf("Error: timestampBytes slice is too short to contain a valid timestamp.")
		}
	}
}

func (server *Server) HeartbeatRequest(handler HandleHeartbeatRequest) {
	server.heartbeatRequestHandler = handler
}

func (server *Server) HeartbeatResponse(handler HandleHeartbeatResponse) {
	server.heartbeatResponseHandler = handler
}

func ParseUDPMessage(data []byte) PfcpMessage {
	header := messages.ParsePFCPHeader(data)
	return PfcpMessage{Header: header, Message: data[12:]}
}
