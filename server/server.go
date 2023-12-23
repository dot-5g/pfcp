package server

import (
	"log"
	"net"

	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/network"
)

const HeaderSize = 8

const (
	HeartbeatRequestType  byte = 1
	HeartbeatResponseType byte = 2
)

// Define a new handler type that specifically accepts RecoveryTimeStampIE
type HandleHeartbeatRequest func(sequenceNumber uint32, recoveryTimeStampIE ie.RecoveryTimeStamp)
type HandleHeartbeatResponse func(sequenceNumber uint32, recoveryTimeStampIE ie.RecoveryTimeStamp)
type MessageHandler func(header messages.PFCPHeader, ies []ie.InformationElement)

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
	server.registerHandler(HeartbeatRequestType, func(header messages.PFCPHeader, ies []ie.InformationElement) {
		var recoveryTimeStamp ie.RecoveryTimeStamp
		for _, elem := range ies {
			if tsIE, ok := elem.(ie.RecoveryTimeStamp); ok {
				recoveryTimeStamp = tsIE
				break
			}
		}

		handler(header.SequenceNumber, recoveryTimeStamp)
	})
}

func (server *Server) HeartbeatResponse(handler HandleHeartbeatResponse) {
	server.registerHandler(HeartbeatResponseType, func(header messages.PFCPHeader, ies []ie.InformationElement) {
		var recoveryTimeStamp ie.RecoveryTimeStamp
		for _, elem := range ies {
			if tsIE, ok := elem.(ie.RecoveryTimeStamp); ok {
				recoveryTimeStamp = tsIE
				break
			}
		}

		handler(header.SequenceNumber, recoveryTimeStamp)
	})
}

func (server *Server) handleUDPMessage(data []byte, addr net.Addr) {
	header, err := messages.ParsePFCPHeader(data[:HeaderSize])
	if err != nil {
		log.Printf("Error parsing PFCP header: %v", err)
		return
	}
	ies, err := ie.ParseInformationElements(data[HeaderSize:])
	if err != nil {
		log.Printf("Error parsing Information Elements: %v", err)
		return
	}

	if handler, exists := server.messageHandlers[header.MessageType]; exists {
		handler(header, ies)
	} else {
		log.Printf("No handler registered for message type %d", header.MessageType)
	}
}

func (server *Server) registerHandler(messageType byte, handler MessageHandler) {
	server.messageHandlers[messageType] = handler
}
