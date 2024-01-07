package client

import (
	"log"

	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/network"
)

type Pfcp struct {
	ServerAddress string
	Udp           network.UdpSender
}

func New(ServerAddress string) *Pfcp {
	udpClient, err := network.NewUdp(ServerAddress)
	if err != nil {
		log.Printf("Failed to initialize PFCP client: %v\n", err)
		return nil
	}
	return &Pfcp{ServerAddress: ServerAddress, Udp: udpClient}
}

func (pfcp *Pfcp) sendNodePfcpMessage(message messages.PFCPMessage, sequenceNumber uint32) error {
	messageType := message.GetMessageType()
	header := messages.NewNodePFCPHeader(messageType, sequenceNumber)
	return pfcp.sendPfcpMessage(message, header)
}

func (pfcp *Pfcp) sendSessionPfcpMessage(message messages.PFCPMessage, seid uint64, sequenceNumber uint32) error {
	messageType := message.GetMessageType()
	header := messages.NewSessionPFCPHeader(messageType, seid, sequenceNumber)
	return pfcp.sendPfcpMessage(message, header)
}

func (pfcp *Pfcp) sendPfcpMessage(message messages.PFCPMessage, header messages.PFCPHeader) error {
	payload := messages.Serialize(message, header)
	if err := pfcp.Udp.Send(payload); err != nil {
		log.Printf("Failed to send PFCP: %v\n", err)
		return err
	}
	log.Printf("PFCP message sent successfully to %s.\n", pfcp.ServerAddress)
	return nil
}

func (pfcp *Pfcp) SendHeartbeatRequest(msg messages.HeartbeatRequest, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *Pfcp) SendHeartbeatResponse(msg messages.HeartbeatResponse, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *Pfcp) SendPFCPAssociationSetupRequest(msg messages.PFCPAssociationSetupRequest, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *Pfcp) SendPFCPAssociationSetupResponse(msg messages.PFCPAssociationSetupResponse, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *Pfcp) SendPFCPAssociationUpdateRequest(msg messages.PFCPAssociationUpdateRequest, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *Pfcp) SendPFCPAssociationUpdateResponse(msg messages.PFCPAssociationUpdateResponse, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *Pfcp) SendPFCPAssociationReleaseRequest(msg messages.PFCPAssociationReleaseRequest, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *Pfcp) SendPFCPAssociationReleaseResponse(msg messages.PFCPAssociationReleaseResponse, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *Pfcp) SendPFCPNodeReportRequest(msg messages.PFCPNodeReportRequest, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *Pfcp) SendPFCPNodeReportResponse(msg messages.PFCPNodeReportResponse, sequenceNumber uint32) error {
	return pfcp.sendNodePfcpMessage(msg, sequenceNumber)
}

func (pfcp *Pfcp) SendPFCPSessionEstablishmentRequest(msg messages.PFCPSessionEstablishmentRequest, seid uint64, sequenceNumber uint32) error {
	return pfcp.sendSessionPfcpMessage(msg, seid, sequenceNumber)
}

func (pfcp *Pfcp) SendPFCPSessionEstablishmentResponse(msg messages.PFCPSessionEstablishmentResponse, seid uint64, sequenceNumber uint32) error {
	return pfcp.sendSessionPfcpMessage(msg, seid, sequenceNumber)
}

func (pfcp *Pfcp) SendPFCPSessionDeletionRequest(msg messages.PFCPSessionDeletionRequest, seid uint64, sequenceNumber uint32) error {
	return pfcp.sendSessionPfcpMessage(msg, seid, sequenceNumber)
}

func (pfcp *Pfcp) SendPFCPSessionDeletionResponse(msg messages.PFCPSessionDeletionResponse, seid uint64, sequenceNumber uint32) error {
	return pfcp.sendSessionPfcpMessage(msg, seid, sequenceNumber)
}

func (pfcp *Pfcp) SendPFCPSessionReportRequest(msg messages.PFCPSessionReportRequest, seid uint64, sequenceNumber uint32) error {
	return pfcp.sendSessionPfcpMessage(msg, seid, sequenceNumber)
}

func (pfcp *Pfcp) SendPFCPSessionReportResponse(msg messages.PFCPSessionReportResponse, seid uint64, sequenceNumber uint32) error {
	return pfcp.sendSessionPfcpMessage(msg, seid, sequenceNumber)
}
