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
type HandlePFCPAssociationReleaseRequest func(sequenceNumber uint32, msg messages.PFCPAssociationReleaseRequest)
type HandlePFCPAssociationReleaseResponse func(sequenceNumber uint32, msg messages.PFCPAssociationReleaseResponse)
type HandlePFCPNodeReportRequest func(sequenceNumber uint32, msg messages.PFCPNodeReportRequest)
type HandlePFCPNodeReportResponse func(sequenceNumber uint32, msg messages.PFCPNodeReportResponse)
type HandlePFCPSessionEstablishmentRequest func(sequenceNumber uint32, msg messages.PFCPSessionEstablishmentRequest)

type Server struct {
	address   string
	udpServer *network.UdpServer

	heartbeatRequestHandler                HandleHeartbeatRequest
	heartbeatResponseHandler               HandleHeartbeatResponse
	pfcpAssociationSetupRequestHandler     HandlePFCPAssociationSetupRequest
	pfcpAssociationSetupResponseHandler    HandlePFCPAssociationSetupResponse
	pfcpAssociationUpdateRequestHandler    HandlePFCPAssociationUpdateRequest
	pfcpAssociationUpdateResponseHandler   HandlePFCPAssociationUpdateResponse
	pfcpAssociationReleaseRequestHandler   HandlePFCPAssociationReleaseRequest
	pfcpAssociationReleaseResponseHandler  HandlePFCPAssociationReleaseResponse
	pfcpNodeReportRequestHandler           HandlePFCPNodeReportRequest
	pfcpNodeReportResponseHandler          HandlePFCPNodeReportResponse
	pfcpSessionEstablishmentRequestHandler HandlePFCPSessionEstablishmentRequest
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

func (server *Server) PFCPAssociationReleaseRequest(handler HandlePFCPAssociationReleaseRequest) {
	server.pfcpAssociationReleaseRequestHandler = handler
}

func (server *Server) PFCPAssociationReleaseResponse(handler HandlePFCPAssociationReleaseResponse) {
	server.pfcpAssociationReleaseResponseHandler = handler
}

func (server *Server) PFCPNodeReportRequest(handler HandlePFCPNodeReportRequest) {
	server.pfcpNodeReportRequestHandler = handler
}

func (server *Server) PFCPNodeReportResponse(handler HandlePFCPNodeReportResponse) {
	server.pfcpNodeReportResponseHandler = handler
}

func (server *Server) PFCPSessionEstablishmentRequest(handler HandlePFCPSessionEstablishmentRequest) {
	server.pfcpSessionEstablishmentRequestHandler = handler
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
	case messages.PFCPAssociationReleaseRequestMessageType:
		if server.pfcpAssociationReleaseRequestHandler == nil {
			log.Printf("No handler for PFCP Association Release Request")
			return
		}
		msg, err := messages.ParsePFCPAssociationReleaseRequest(data[headers.HeaderSize:])
		if err != nil {
			log.Printf("Error parsing PFCP Association Release Request: %v", err)
			return
		}
		server.pfcpAssociationReleaseRequestHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationReleaseResponseMessageType:
		if server.pfcpAssociationReleaseResponseHandler == nil {
			log.Printf("No handler for PFCP Association Release Response")
			return
		}
		msg, err := messages.ParsePFCPAssociationReleaseResponse(data[headers.HeaderSize:])
		if err != nil {
			log.Printf("Error parsing PFCP Association Release Response: %v", err)
			return
		}
		server.pfcpAssociationReleaseResponseHandler(header.SequenceNumber, msg)
	case messages.PFCPNodeReportRequestMessageType:
		if server.pfcpNodeReportRequestHandler == nil {
			log.Printf("No handler for PFCP Node Report Request")
			return
		}
		msg, err := messages.ParsePFCPNodeReportRequest(data[headers.HeaderSize:])
		if err != nil {
			log.Printf("Error parsing PFCP Node Report Request: %v", err)
			return
		}
		server.pfcpNodeReportRequestHandler(header.SequenceNumber, msg)
	case messages.PFCPNodeReportResponseMessageType:
		if server.pfcpNodeReportResponseHandler == nil {
			log.Printf("No handler for PFCP Node Report Response")
			return
		}
		msg, err := messages.ParsePFCPNodeReportResponse(data[headers.HeaderSize:])
		if err != nil {
			log.Printf("Error parsing PFCP Node Report Response: %v", err)
			return
		}
		server.pfcpNodeReportResponseHandler(header.SequenceNumber, msg)
	case messages.PFCPSessionEstablishmentRequestMessageType:
		if server.pfcpSessionEstablishmentRequestHandler == nil {
			log.Printf("No handler for PFCP Session Establishment Request")
			return
		}
		msg, err := messages.ParsePFCPSessionEstablishmentRequest(data[headers.HeaderSize:])
		if err != nil {
			log.Printf("Error parsing PFCP Session Establishment Request: %v", err)
			return
		}
		server.pfcpSessionEstablishmentRequestHandler(header.SequenceNumber, msg)
	default:
		log.Printf("Unknown PFCP message type: %v", header.MessageType)
	}
}
