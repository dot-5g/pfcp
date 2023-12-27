package server

import (
	"log"

	"github.com/dot-5g/pfcp/headers"
	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/network"
)

type HandleHeartbeatRequest func(sequenceNumber uint32, msg messages.HeartbeatRequest)
type HandleHeartbeatResponse func(sequenceNumber uint32, msg messages.HeartbeatResponse)
type HandlePFCPAssociationSetupRequest func(sequenceNumber uint32, msg messages.PFCPAssociationSetupRequest)
type HandlePFCPAssociationSetupResponse func(sequenceNumber uint32, msg messages.PFCPAssociationSetupResponse)
type HandlePFCPAssociationUpdateRequest func(sequenceNumber uint32, msg messages.PFCPAssociationUpdateRequest)
type HandlePFCPAssociationUpdateResponse func(sequenceNumber uint32, msg messages.PFCPAssociationUpdateResponse)

type Server struct {
	address   string
	udpServer *network.UdpServer

	heartbeatRequestHandler              HandleHeartbeatRequest
	heartbeatResponseHandler             HandleHeartbeatResponse
	pfcpAssociationSetupRequestHandler   HandlePFCPAssociationSetupRequest
	pfcpAssociationSetupResponseHandler  HandlePFCPAssociationSetupResponse
	pfcpAssociationUpdateRequestHandler  HandlePFCPAssociationUpdateRequest
	pfcpAssociationUpdateResponseHandler HandlePFCPAssociationUpdateResponse
}

func New(address string) *Server {
	server := &Server{
		address:   address,
		udpServer: network.NewUdpServer(),
	}
	return server
}

func (server *Server) Run() error {
	server.udpServer.SetHandler(server.handleUDPMessage)
	err := server.udpServer.Run(server.address)
	return err
}

func (server *Server) Close() {
	server.udpServer.Close()
}

func (server *Server) HeartbeatRequest(handler HandleHeartbeatRequest) {
	server.heartbeatRequestHandler = handler
}

func (server *Server) HeartbeatResponse(handler HandleHeartbeatResponse) {
	server.heartbeatResponseHandler = handler
}

func (server *Server) PFCPAssociationSetupRequest(handler HandlePFCPAssociationSetupRequest) {
	server.pfcpAssociationSetupRequestHandler = handler
}

func (server *Server) PFCPAssociationSetupResponse(handler HandlePFCPAssociationSetupResponse) {
	server.pfcpAssociationSetupResponseHandler = handler
}

func (server *Server) PFCPAssociationUpdateRequest(handler HandlePFCPAssociationUpdateRequest) {
	server.pfcpAssociationUpdateRequestHandler = handler
}

func (server *Server) PFCPAssociationUpdateResponse(handler HandlePFCPAssociationUpdateResponse) {
	server.pfcpAssociationUpdateResponseHandler = handler
}

func (server *Server) handleUDPMessage(data []byte) {
	header, err := headers.ParsePFCPHeader(data[:headers.HeaderSize])
	if err != nil {
		log.Printf("Error parsing PFCP header: %v", err)
		return
	}

	switch header.MessageType {
	case messages.HeartbeatRequestMessageType:
		if server.heartbeatRequestHandler == nil {
			log.Printf("No handler for Heartbeat Request")
			return
		}
		msg, err := messages.ParseHeartbeatRequest(data[headers.HeaderSize:])
		if err != nil {
			log.Printf("Error parsing Heartbeat Request: %v", err)
			return
		}
		server.heartbeatRequestHandler(header.SequenceNumber, msg)
	case messages.HeartbeatResponseMessageType:
		if server.heartbeatResponseHandler == nil {
			log.Printf("No handler for Heartbeat Response")
			return
		}
		msg, err := messages.ParseHeartbeatResponse(data[headers.HeaderSize:])
		if err != nil {
			log.Printf("Error parsing Heartbeat Response: %v", err)
			return
		}
		server.heartbeatResponseHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationSetupRequestMessageType:
		if server.pfcpAssociationSetupRequestHandler == nil {
			log.Printf("No handler for PFCP Association Setup Request")
			return
		}
		msg, err := messages.ParsePFCPAssociationSetupRequest(data[headers.HeaderSize:])
		if err != nil {
			log.Printf("Error parsing PFCP Association Setup Request: %v", err)
			return
		}
		server.pfcpAssociationSetupRequestHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationSetupResponseMessageType:
		if server.pfcpAssociationSetupResponseHandler == nil {
			log.Printf("No handler for PFCP Association Setup Response")
			return
		}
		msg, err := messages.ParsePFCPAssociationSetupResponse(data[headers.HeaderSize:])
		if err != nil {
			log.Printf("Error parsing PFCP Association Setup Response: %v", err)
			return
		}
		server.pfcpAssociationSetupResponseHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationUpdateRequestMessageType:
		if server.pfcpAssociationUpdateRequestHandler == nil {
			log.Printf("No handler for PFCP Association Update Request")
			return
		}
		msg, err := messages.ParsePFCPAssociationUpdateRequest(data[headers.HeaderSize:])
		if err != nil {
			log.Printf("Error parsing PFCP Association Update Request: %v", err)
			return
		}
		server.pfcpAssociationUpdateRequestHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationUpdateResponseMessageType:
		if server.pfcpAssociationUpdateResponseHandler == nil {
			log.Printf("No handler for PFCP Association Update Response")
			return
		}
		msg, err := messages.ParsePFCPAssociationUpdateResponse(data[headers.HeaderSize:])
		if err != nil {
			log.Printf("Error parsing PFCP Association Update Response: %v", err)
			return
		}
		server.pfcpAssociationUpdateResponseHandler(header.SequenceNumber, msg)
	default:
		log.Printf("Unknown PFCP message type: %v", header.MessageType)
	}
}
