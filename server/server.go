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
type HandlePFCPSessionDeletionRequest func(sequenceNumber uint32, seid uint64, msg messages.PFCPSessionDeletionRequest)
type HandlePFCPSessionDeletionResponse func(sequenceNumber uint32, seid uint64, msg messages.PFCPSessionDeletionResponse)
type HandlePFCPSessionReportRequest func(sequenceNumber uint32, seid uint64, msg messages.PFCPSessionReportRequest)
type HandlePFCPSessionReportResponse func(sequenceNumber uint32, seid uint64, msg messages.PFCPSessionReportResponse)

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
	pfcpSessionDeletionRequestHandler       HandlePFCPSessionDeletionRequest
	pfcpSessionDeletionResponseHandler      HandlePFCPSessionDeletionResponse
	pfcpSessionReportRequestHandler         HandlePFCPSessionReportRequest
	pfcpSessionReportResponseHandler        HandlePFCPSessionReportResponse
}

func New(address string) *Server {
	server := &Server{
		address:   address,
		udpServer: network.NewUdpServer(),
	}
	return server
}

func (server *Server) Run() error {
	server.udpServer.SetHandler(server.handlePFCPMessage)
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

func (server *Server) PFCPSessionDeletionRequest(handler HandlePFCPSessionDeletionRequest) {
	server.pfcpSessionDeletionRequestHandler = handler
}

func (server *Server) PFCPSessionDeletionResponse(handler HandlePFCPSessionDeletionResponse) {
	server.pfcpSessionDeletionResponseHandler = handler
}

func (server *Server) PFCPSessionReportRequest(handler HandlePFCPSessionReportRequest) {
	server.pfcpSessionReportRequestHandler = handler
}

func (server *Server) PFCPSessionReportResponse(handler HandlePFCPSessionReportResponse) {
	server.pfcpSessionReportResponseHandler = handler
}

func (server *Server) handlePFCPMessage(payload []byte) {
	header, err := messages.DeserializeHeader(payload)
	if err != nil {
		log.Fatalf("Error deserializing header: %v", err)
		return
	}

	payloadOffset := 8
	if header.S {
		payloadOffset = 16
	}

	if len(payload) < payloadOffset {
		log.Fatalf("Error: payload length is less than payload offset")
		return
	}
	payloadMessage := payload[payloadOffset:]

	switch header.MessageType {
	case messages.HeartbeatRequestMessageType:
		if server.heartbeatRequestHandler == nil {
			log.Printf("No handler for Heartbeat Request")
			return
		}
		msg, err := messages.DeserializeHeartbeatRequest(payloadMessage)
		if err != nil {
			log.Printf("Error deserializing Heartbeat Request: %v", err)
			return
		}
		server.heartbeatRequestHandler(header.SequenceNumber, msg)
	case messages.HeartbeatResponseMessageType:
		if server.heartbeatResponseHandler == nil {
			log.Printf("No handler for Heartbeat Response")
			return
		}
		msg, err := messages.DeserializeHeartbeatResponse(payloadMessage)
		if err != nil {
			log.Printf("Error deserializing Heartbeat Response: %v", err)
			return
		}
		server.heartbeatResponseHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationSetupRequestMessageType:
		if server.pfcpAssociationSetupRequestHandler == nil {
			log.Printf("No handler for PFCP Association Setup Request")
			return
		}
		msg, err := messages.DeserializePFCPAssociationSetupRequest(payloadMessage)
		if err != nil {
			log.Printf("Error deserializing PFCP Association Setup Request: %v", err)
			return
		}
		server.pfcpAssociationSetupRequestHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationSetupResponseMessageType:
		if server.pfcpAssociationSetupResponseHandler == nil {
			log.Printf("No handler for PFCP Association Setup Response")
			return
		}
		msg, err := messages.DeserializePFCPAssociationSetupResponse(payloadMessage)
		if err != nil {
			log.Printf("Error deserializing PFCP Association Setup Response: %v", err)
			return
		}
		server.pfcpAssociationSetupResponseHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationUpdateRequestMessageType:
		if server.pfcpAssociationUpdateRequestHandler == nil {
			log.Printf("No handler for PFCP Association Update Request")
			return
		}
		msg, err := messages.DeserializePFCPAssociationUpdateRequest(payloadMessage)
		if err != nil {
			log.Printf("Error deserializing PFCP Association Update Request: %v", err)
			return
		}
		server.pfcpAssociationUpdateRequestHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationUpdateResponseMessageType:
		if server.pfcpAssociationUpdateResponseHandler == nil {
			log.Printf("No handler for PFCP Association Update Response")
			return
		}
		msg, err := messages.DeserializePFCPAssociationUpdateResponse(payloadMessage)
		if err != nil {
			log.Printf("Error deserializing PFCP Association Update Response: %v", err)
			return
		}
		server.pfcpAssociationUpdateResponseHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationReleaseRequestMessageType:
		if server.pfcpAssociationReleaseRequestHandler == nil {
			log.Printf("No handler for PFCP Association Release Request")
			return
		}
		msg, err := messages.DeserializePFCPAssociationReleaseRequest(payloadMessage)
		if err != nil {
			log.Printf("Error deserializing PFCP Association Release Request: %v", err)
			return
		}
		server.pfcpAssociationReleaseRequestHandler(header.SequenceNumber, msg)
	case messages.PFCPAssociationReleaseResponseMessageType:
		if server.pfcpAssociationReleaseResponseHandler == nil {
			log.Printf("No handler for PFCP Association Release Response")
			return
		}
		msg, err := messages.DeserializePFCPAssociationReleaseResponse(payloadMessage)
		if err != nil {
			log.Printf("Error deserializing PFCP Association Release Response: %v", err)
			return
		}
		server.pfcpAssociationReleaseResponseHandler(header.SequenceNumber, msg)
	case messages.PFCPNodeReportRequestMessageType:
		if server.pfcpNodeReportRequestHandler == nil {
			log.Printf("No handler for PFCP Node Report Request")
			return
		}
		msg, err := messages.DeserializePFCPNodeReportRequest(payloadMessage)
		if err != nil {
			log.Printf("Error deserializing PFCP Node Report Request: %v", err)
			return
		}
		server.pfcpNodeReportRequestHandler(header.SequenceNumber, msg)
	case messages.PFCPNodeReportResponseMessageType:
		if server.pfcpNodeReportResponseHandler == nil {
			log.Printf("No handler for PFCP Node Report Response")
			return
		}
		msg, err := messages.DeserializePFCPNodeReportResponse(payloadMessage)
		if err != nil {
			log.Printf("Error deserializing PFCP Node Report Response: %v", err)
			return
		}
		server.pfcpNodeReportResponseHandler(header.SequenceNumber, msg)
	case messages.PFCPSessionEstablishmentRequestMessageType:
		if server.pfcpSessionEstablishmentRequestHandler == nil {
			log.Printf("No handler for PFCP Session Establishment Request")
			return
		}
		msg, err := messages.DeserializePFCPSessionEstablishmentRequest(payloadMessage)
		if err != nil {
			log.Printf("Error deserializing PFCP Session Establishment Request: %v", err)
			return
		}
		server.pfcpSessionEstablishmentRequestHandler(header.SequenceNumber, header.SEID, msg)
	case messages.PFCPSessionEstablishmentResponseMessageType:
		if server.pfcpSessionEstablishmentResponseHandler == nil {
			log.Printf("No handler for PFCP Session Establishment Response")
			return
		}
		msg, err := messages.DeserializePFCPSessionEstablishmentResponse(payloadMessage)
		if err != nil {
			log.Printf("Error deserializing PFCP Session Establishment Response: %v", err)
			return
		}
		server.pfcpSessionEstablishmentResponseHandler(header.SequenceNumber, header.SEID, msg)
	case messages.PFCPSessionDeletionRequestMessageType:
		if server.pfcpSessionDeletionRequestHandler == nil {
			log.Printf("No handler for PFCP Session Deletion Request")
			return
		}
		msg, err := messages.DeserializePFCPSessionDeletionRequest(payloadMessage)
		if err != nil {
			log.Printf("Error deserializing PFCP Session Deletion Request: %v", err)
			return
		}
		server.pfcpSessionDeletionRequestHandler(header.SequenceNumber, header.SEID, msg)
	case messages.PFCPSessionDeletionResponseMessageType:
		if server.pfcpSessionDeletionResponseHandler == nil {
			log.Printf("No handler for PFCP Session Deletion Response")
			return
		}
		msg, err := messages.DeserializePFCPSessionDeletionResponse(payloadMessage)
		if err != nil {
			log.Printf("Error deserializing PFCP Session Deletion Response: %v", err)
			return
		}
		server.pfcpSessionDeletionResponseHandler(header.SequenceNumber, header.SEID, msg)
	case messages.PFCPSessionReportRequestMessageType:
		if server.pfcpSessionReportRequestHandler == nil {
			log.Printf("No handler for PFCP Session Report Request")
			return
		}
		msg, err := messages.DeserializePFCPSessionReportRequest(payloadMessage)
		if err != nil {
			log.Printf("Error deserializing PFCP Session Report Request: %v", err)
			return
		}
		server.pfcpSessionReportRequestHandler(header.SequenceNumber, header.SEID, msg)
	case messages.PFCPSessionReportResponseMessageType:
		if server.pfcpSessionReportResponseHandler == nil {
			log.Printf("No handler for PFCP Session Report Response")
			return
		}
		msg, err := messages.DeserializePFCPSessionReportResponse(payloadMessage)
		if err != nil {
			log.Printf("Error deserializing PFCP Session Report Response: %v", err)
			return
		}
		server.pfcpSessionReportResponseHandler(header.SequenceNumber, header.SEID, msg)
	default:
		log.Printf("Unknown PFCP message type: %v", header.MessageType)
	}
}
