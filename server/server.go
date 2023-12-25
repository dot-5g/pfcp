package server

import (
	"log"

	"github.com/dot-5g/pfcp/headers"
	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/network"
)

const HeaderSize = 8

const (
	HeartbeatRequestType             byte = 1
	HeartbeatResponseType            byte = 2
	PFCPAssociationSetupRequestType  byte = 5
	PFCPAssociationSetupResponseType byte = 6
)

type HandleHeartbeatRequest func(sequenceNumber uint32, recoveryTimeStampIE ie.RecoveryTimeStamp)
type HandleHeartbeatResponse func(sequenceNumber uint32, recoveryTimeStampIE ie.RecoveryTimeStamp)
type HandlePFCPAssociationSetupRequest func(sequenceNumber uint32, nodeID ie.NodeID, recoveryTimeStampIE ie.RecoveryTimeStamp)
type HandlePFCPAssociationSetupResponse func(sequenceNumber uint32, nodeID ie.NodeID, cause ie.Cause, recoveryTimeStampIE ie.RecoveryTimeStamp)

type MessageHandler func(header headers.PFCPHeader, ies []ie.InformationElement)

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

func (server *Server) Run() error {
	server.udpServer.SetHandler(server.handleUDPMessage)
	err := server.udpServer.Run(server.address)
	if err != nil {
		return err
	}
	return nil
}

func (server *Server) Close() {
	server.udpServer.Close()
}

func (server *Server) HeartbeatRequest(handler HandleHeartbeatRequest) {
	server.registerHandler(HeartbeatRequestType, func(header headers.PFCPHeader, ies []ie.InformationElement) {
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
	server.registerHandler(HeartbeatResponseType, func(header headers.PFCPHeader, ies []ie.InformationElement) {
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

func (server *Server) PFCPAssociationSetupRequest(handler HandlePFCPAssociationSetupRequest) {
	server.registerHandler(PFCPAssociationSetupRequestType, func(header headers.PFCPHeader, ies []ie.InformationElement) {
		var recoveryTimeStamp ie.RecoveryTimeStamp
		var nodeID ie.NodeID
		for _, elem := range ies {
			if tsIE, ok := elem.(ie.RecoveryTimeStamp); ok {
				recoveryTimeStamp = tsIE
			}
			if nodeIDIE, ok := elem.(ie.NodeID); ok {
				nodeID = nodeIDIE
			}
		}

		handler(header.SequenceNumber, nodeID, recoveryTimeStamp)
	})
}

func (server *Server) PFCPAssociationSetupResponse(handler HandlePFCPAssociationSetupResponse) {
	server.registerHandler(PFCPAssociationSetupResponseType, func(header headers.PFCPHeader, ies []ie.InformationElement) {
		var recoveryTimeStamp ie.RecoveryTimeStamp
		var nodeID ie.NodeID
		var cause ie.Cause
		for _, elem := range ies {
			if tsIE, ok := elem.(ie.RecoveryTimeStamp); ok {
				recoveryTimeStamp = tsIE
			}
			if nodeIDIE, ok := elem.(ie.NodeID); ok {
				nodeID = nodeIDIE
			}
			if causeIE, ok := elem.(ie.Cause); ok {
				cause = causeIE
			}
		}

		handler(header.SequenceNumber, nodeID, cause, recoveryTimeStamp)
	})
}

func (server *Server) handleUDPMessage(data []byte) {
	header, err := headers.ParsePFCPHeader(data[:HeaderSize])
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
