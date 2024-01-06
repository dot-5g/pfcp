package server

import (
	"log"

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
type HandlePFCPSessionEstablishmentRequest func(sequenceNumber uint32, seid uint64, msg messages.PFCPSessionEstablishmentRequest)
type HandlePFCPSessionEstablishmentResponse func(sequenceNumber uint32, seid uint64, msg messages.PFCPSessionEstablishmentResponse)

type Server struct {
	address   string
	udpServer *network.UdpServer

	heartbeatRequestHandler                 HandleHeartbeatRequest
	heartbeatResponseHandler                HandleHeartbeatResponse
	pfcpAssociationSetupRequestHandler      HandlePFCPAssociationSetupRequest
	pfcpAssociationSetupResponseHandler     HandlePFCPAssociationSetupResponse
	pfcpAssociationUpdateRequestHandler     HandlePFCPAssociationUpdateRequest
	pfcpAssociationUpdateResponseHandler    HandlePFCPAssociationUpdateResponse
	pfcpAssociationReleaseRequestHandler    HandlePFCPAssociationReleaseRequest
	pfcpAssociationReleaseResponseHandler   HandlePFCPAssociationReleaseResponse
	pfcpNodeReportRequestHandler            HandlePFCPNodeReportRequest
	pfcpNodeReportResponseHandler           HandlePFCPNodeReportResponse
	pfcpSessionEstablishmentRequestHandler  HandlePFCPSessionEstablishmentRequest
	pfcpSessionEstablishmentResponseHandler HandlePFCPSessionEstablishmentResponse
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

func (server *Server) PFCPSessionEstablishmentResponse(handler HandlePFCPSessionEstablishmentResponse) {
	server.pfcpSessionEstablishmentResponseHandler = handler
}

func (server *Server) handleUDPMessage(payload []byte) {
	header, genericMessage, err := messages.DeserializePFCPMessage(payload)

	if err != nil {
		log.Printf("Error deserializing PFCP message: %v", err)
		return
	}

	switch header.MessageType {
	case messages.HeartbeatRequestMessageType:
		if server.heartbeatRequestHandler == nil {
			log.Printf("No handler for Heartbeat Request")
			return
		}
		msg, ok := genericMessage.(messages.HeartbeatRequest)
		if !ok {
			log.Printf("Error asserting Heartbeat Request type")
			return
		}
		server.heartbeatRequestHandler(header.SequenceNumber, msg)
	case messages.HeartbeatResponseMessageType:
		if server.heartbeatResponseHandler == nil {
			log.Printf("No handler for Heartbeat Response")
			return
		}
		msg, ok := genericMessage.(messages.HeartbeatResponse)
		if !ok {
			log.Printf("Error asserting Heartbeat Response type")
			return
		}
		server.heartbeatResponseHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationSetupRequestMessageType:
		if server.pfcpAssociationSetupRequestHandler == nil {
			log.Printf("No handler for PFCP Association Setup Request")
			return
		}
		msg, ok := genericMessage.(messages.PFCPAssociationSetupRequest)
		if !ok {
			log.Printf("Error asserting PFCP Association Setup Request type")
			return
		}
		server.pfcpAssociationSetupRequestHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationSetupResponseMessageType:
		if server.pfcpAssociationSetupResponseHandler == nil {
			log.Printf("No handler for PFCP Association Setup Response")
			return
		}
		msg, ok := genericMessage.(messages.PFCPAssociationSetupResponse)
		if !ok {
			log.Printf("Error asserting PFCP Association Setup Response type")
			return
		}
		server.pfcpAssociationSetupResponseHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationUpdateRequestMessageType:
		if server.pfcpAssociationUpdateRequestHandler == nil {
			log.Printf("No handler for PFCP Association Update Request")
			return
		}
		msg, ok := genericMessage.(messages.PFCPAssociationUpdateRequest)
		if !ok {
			log.Printf("Error asserting PFCP Association Update Request type")
			return
		}
		server.pfcpAssociationUpdateRequestHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationUpdateResponseMessageType:
		if server.pfcpAssociationUpdateResponseHandler == nil {
			log.Printf("No handler for PFCP Association Update Response")
			return
		}
		msg, ok := genericMessage.(messages.PFCPAssociationUpdateResponse)
		if !ok {
			log.Printf("Error asserting PFCP Association Update Response type")
			return
		}
		server.pfcpAssociationUpdateResponseHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationReleaseRequestMessageType:
		if server.pfcpAssociationReleaseRequestHandler == nil {
			log.Printf("No handler for PFCP Association Release Request")
			return
		}
		msg, ok := genericMessage.(messages.PFCPAssociationReleaseRequest)
		if !ok {
			log.Printf("Error asserting PFCP Association Release Request type")
			return
		}
		server.pfcpAssociationReleaseRequestHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationReleaseResponseMessageType:
		if server.pfcpAssociationReleaseResponseHandler == nil {
			log.Printf("No handler for PFCP Association Release Response")
			return
		}
		msg, ok := genericMessage.(messages.PFCPAssociationReleaseResponse)
		if !ok {
			log.Printf("Error asserting PFCP Association Release Response type")
			return
		}
		server.pfcpAssociationReleaseResponseHandler(header.SequenceNumber, msg)
	case messages.PFCPNodeReportRequestMessageType:
		if server.pfcpNodeReportRequestHandler == nil {
			log.Printf("No handler for PFCP Node Report Request")
			return
		}
		msg, ok := genericMessage.(messages.PFCPNodeReportRequest)
		if !ok {
			log.Printf("Error asserting PFCP Node Report Request type")
			return
		}
		server.pfcpNodeReportRequestHandler(header.SequenceNumber, msg)
	case messages.PFCPNodeReportResponseMessageType:
		if server.pfcpNodeReportResponseHandler == nil {
			log.Printf("No handler for PFCP Node Report Response")
			return
		}
		msg, ok := genericMessage.(messages.PFCPNodeReportResponse)
		if !ok {
			log.Printf("Error asserting PFCP Node Report Response type")
			return
		}
		server.pfcpNodeReportResponseHandler(header.SequenceNumber, msg)
	case messages.PFCPSessionEstablishmentRequestMessageType:
		if server.pfcpSessionEstablishmentRequestHandler == nil {
			log.Printf("No handler for PFCP Session Establishment Request")
			return
		}
		msg, ok := genericMessage.(messages.PFCPSessionEstablishmentRequest)
		if !ok {
			log.Printf("Error asserting PFCP Session Establishment Request type")
			return
		}
		server.pfcpSessionEstablishmentRequestHandler(header.SequenceNumber, header.SEID, msg)
	case messages.PFCPSessionEstablishmentResponseMessageType:
		if server.pfcpSessionEstablishmentResponseHandler == nil {
			log.Printf("No handler for PFCP Session Establishment Response")
			return
		}
		msg, ok := genericMessage.(messages.PFCPSessionEstablishmentResponse)
		if !ok {
			log.Printf("Error asserting PFCP Session Establishment Response type")
			return
		}
		server.pfcpSessionEstablishmentResponseHandler(header.SequenceNumber, header.SEID, msg)
	default:
		log.Printf("Unknown PFCP message type: %v", header.MessageType)
	}
}
