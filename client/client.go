package client

import (
	"log"

	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/network"
)

type PfcpClienter interface {
	SendHeartbeatRequest(msg messages.HeartbeatRequest, sequenceNumber uint32) error
	SendHeartbeatResponse(msg messages.HeartbeatResponse, sequenceNumber uint32) error
	SendPFCPAssociationSetupRequest(msg messages.PFCPAssociationSetupRequest, sequenceNumber uint32) error
	SendPFCPAssociationSetupResponse(msg messages.PFCPAssociationSetupResponse, sequenceNumber uint32) error
	SendPFCPAssociationUpdateRequest(msg messages.PFCPAssociationUpdateRequest, sequenceNumber uint32) error
	SendPFCPAssociationUpdateResponse(msg messages.PFCPAssociationUpdateResponse, sequenceNumber uint32) error
	SendPFCPAssociationReleaseRequest(msg messages.PFCPAssociationReleaseRequest, sequenceNumber uint32) error
	SendPFCPAssociationReleaseResponse(msg messages.PFCPAssociationReleaseResponse, sequenceNumber uint32) error
	SendPFCPNodeReportRequest(msg messages.PFCPNodeReportRequest, sequenceNumber uint32) error
	SendPFCPNodeReportResponse(msg messages.PFCPNodeReportResponse, sequenceNumber uint32) error
	SendPFCPSessionEstablishmentRequest(msg messages.PFCPSessionEstablishmentRequest, seid uint64, sequenceNumber uint32) error
	SendPFCPSessionEstablishmentResponse(msg messages.PFCPSessionEstablishmentResponse, seid uint64, sequenceNumber uint32) error
	SendPFCPSessionDeletionRequest(msg messages.PFCPSessionDeletionRequest, seid uint64, sequenceNumber uint32) error
	SendPFCPSessionDeletionResponse(msg messages.PFCPSessionDeletionResponse, seid uint64, sequenceNumber uint32) error
	SendPFCPSessionReportRequest(msg messages.PFCPSessionReportRequest, seid uint64, sequenceNumber uint32) error
	SendPFCPSessionReportResponse(msg messages.PFCPSessionReportResponse, seid uint64, sequenceNumber uint32) error
}

var _ PfcpClienter = (*PFCP)(nil)

type PFCP struct {
	ServerAddress string
	Udp           network.UDPSender
}

func New(ServerAddress string) *PFCP {
	udpClient, err := network.NewUDP(ServerAddress)
	if err != nil {
		log.Printf("Failed to initialize PFCP client: %v\n", err)
		return nil
	}
	return &PFCP{ServerAddress: ServerAddress, Udp: udpClient}
}

func (pfcp *PFCP) sendNodePfcpMessage(message messages.PFCPMessage, sequenceNumber uint32) error {
	messageType := message.GetMessageType()
	header := messages.NewNodeHeader(messageType, sequenceNumber)
	return pfcp.sendPfcpMessage(message, header)
}

func (pfcp *PFCP) sendSessionPfcpMessage(message messages.PFCPMessage, seid uint64, sequenceNumber uint32) error {
	messageType := message.GetMessageType()
	header := messages.NewSessionHeader(messageType, seid, sequenceNumber)
	return pfcp.sendPfcpMessage(message, header)
}

func (pfcp *PFCP) sendPfcpMessage(message messages.PFCPMessage, header messages.Header) error {
	messageName := message.GetMessageTypeString()
	payload := messages.Serialize(message, header)
	if err := pfcp.Udp.Send(payload); err != nil {
		log.Printf("Failed to send %s message to %v: %v\n", messageName, pfcp.ServerAddress, err)
		return err
	}
	log.Printf("%s message sent successfully to %s.\n", messageName, pfcp.ServerAddress)
	return nil
}

func (pfcp *PFCP) SendHeartbeatRequest(msg messages.HeartbeatRequest, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *PFCP) SendHeartbeatResponse(msg messages.HeartbeatResponse, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *PFCP) SendPFCPAssociationSetupRequest(msg messages.PFCPAssociationSetupRequest, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *PFCP) SendPFCPAssociationSetupResponse(msg messages.PFCPAssociationSetupResponse, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *PFCP) SendPFCPAssociationUpdateRequest(msg messages.PFCPAssociationUpdateRequest, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *PFCP) SendPFCPAssociationUpdateResponse(msg messages.PFCPAssociationUpdateResponse, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *PFCP) SendPFCPAssociationReleaseRequest(msg messages.PFCPAssociationReleaseRequest, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *PFCP) SendPFCPAssociationReleaseResponse(msg messages.PFCPAssociationReleaseResponse, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *PFCP) SendPFCPNodeReportRequest(msg messages.PFCPNodeReportRequest, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *PFCP) SendPFCPNodeReportResponse(msg messages.PFCPNodeReportResponse, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *PFCP) SendPFCPSessionEstablishmentRequest(msg messages.PFCPSessionEstablishmentRequest, seid uint64, sequenceNumber uint32) error {
	return pfcp.sendSessionPfcpMessage(msg, seid, sequenceNumber)
}

func (pfcp *PFCP) SendPFCPSessionEstablishmentResponse(msg messages.PFCPSessionEstablishmentResponse, seid uint64, sequenceNumber uint32) error {
	return pfcp.sendSessionPfcpMessage(msg, seid, sequenceNumber)
}

func (pfcp *PFCP) SendPFCPSessionDeletionRequest(msg messages.PFCPSessionDeletionRequest, seid uint64, sequenceNumber uint32) error {
	return pfcp.sendSessionPfcpMessage(msg, seid, sequenceNumber)
}

func (pfcp *PFCP) SendPFCPSessionDeletionResponse(msg messages.PFCPSessionDeletionResponse, seid uint64, sequenceNumber uint32) error {
	return pfcp.sendSessionPfcpMessage(msg, seid, sequenceNumber)
}

func (pfcp *PFCP) SendPFCPSessionReportRequest(msg messages.PFCPSessionReportRequest, seid uint64, sequenceNumber uint32) error {
	return pfcp.sendSessionPfcpMessage(msg, seid, sequenceNumber)
}

func (pfcp *PFCP) SendPFCPSessionReportResponse(msg messages.PFCPSessionReportResponse, seid uint64, sequenceNumber uint32) error {
	return pfcp.sendSessionPfcpMessage(msg, seid, sequenceNumber)
}
