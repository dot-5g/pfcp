package server

import (
	"log"
	"net"

	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/network"
)

const HeaderSize = 8

const (
	HeartbeatRequestType  byte = 1
	HeartbeatResponseType byte = 2
)

type HandleHeartbeatRequest func(*messages.HeartbeatRequest)
type HandleHeartbeatResponse func(*messages.HeartbeatResponse)
type MessageHandler func(*PfcpMessage)

type PfcpMessage struct {
	Header  messages.PFCPHeader
	Message []byte
}

type Server struct {
	address         string
	udpServer       *network.UdpServer
	messageHandlers map[byte]MessageHandler
}

func New(address string) *Server {
	server := &Server{
		address:         address,
		udpServer:       network.NewUdpServer(),
		messageHandlers: make(map[byte]MessageHandler),
	}
	return server
}

func (server *Server) Run() {
	server.udpServer.SetHandler(server.handleUDPMessage)
	server.udpServer.Run(server.address)
}

func (server *Server) Close() {
	server.udpServer.Close()
}

func (server *Server) HeartbeatRequest(handler HandleHeartbeatRequest) {
	server.registerHandler(HeartbeatRequestType, func(msg *PfcpMessage) {
		server.handleHeartbeatRequest(msg, handler)
	})
}

func (server *Server) HeartbeatResponse(handler HandleHeartbeatResponse) {
	server.registerHandler(HeartbeatResponseType, func(msg *PfcpMessage) {
		server.handleHeartbeatResponse(msg, handler)
	})
}
func (server *Server) handleUDPMessage(data []byte, addr net.Addr) {
	header, err := messages.ParsePFCPHeader(data[:HeaderSize])
	if err != nil {
		log.Printf("Error parsing PFCP header: %v", err)
		return
	}
	pfcpMessage := PfcpMessage{Header: header, Message: data[HeaderSize:]}

	if genericHandler, exists := server.messageHandlers[pfcpMessage.Header.MessageType]; exists {
		genericHandler(&pfcpMessage)
	} else {
		log.Printf("No handler registered for message type %d", pfcpMessage.Header.MessageType)
	}
}

func (server *Server) handleHeartbeatRequest(msg *PfcpMessage, handler HandleHeartbeatRequest) {
	recoveryTimeStamp := messages.FromBytes(msg.Message)
	heartbeatRequest := messages.HeartbeatRequest{
		RecoveryTimeStamp: recoveryTimeStamp,
	}
	handler(&heartbeatRequest)
}

func (server *Server) handleHeartbeatResponse(msg *PfcpMessage, handler HandleHeartbeatResponse) {
	recoveryTimeStamp := messages.FromBytes(msg.Message)
	heartbeatResponse := messages.HeartbeatResponse{
		RecoveryTimeStamp: recoveryTimeStamp,
	}
	handler(&heartbeatResponse)
}

func (server *Server) registerHandler(messageType byte, handler MessageHandler) {
	server.messageHandlers[messageType] = handler
}
