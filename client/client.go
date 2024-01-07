package client

import (
	"fmt"
	"log"

	"github.com/dot-5g/pfcp/ie"
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
		log.Printf("Failed to initialize UDP client: %v\n", err)
		return nil
	}
	return &Pfcp{ServerAddress: ServerAddress, Udp: udpClient}
}

func (pfcp *Pfcp) sendPfcpMessage(header messages.PFCPHeader, elements []ie.InformationElement) error {
	var payload []byte
	for _, element := range elements {
		payload = append(payload, element.Serialize()...)
	}
	message := serializeMessage(header, payload)
	if err := pfcp.Udp.Send(message); err != nil {
		log.Printf("Failed to send PFCP: %v\n", err)
		return err
	}
	log.Printf("PFCP sent successfully to %s.\n", pfcp.ServerAddress)
	return nil
}

func serializeMessage(header messages.PFCPHeader, payload []byte) []byte {
	header.MessageLength = uint16(len(payload))
	headerBytes := header.Serialize()
	return append(headerBytes, payload...)
}

func (pfcp *Pfcp) SendHeartbeatRequest(msg messages.HeartbeatRequest, sequenceNumber uint32) error {
	header := messages.NewNodePFCPHeader(messages.HeartbeatRequestMessageType, sequenceNumber)

	ies := []ie.InformationElement{msg.RecoveryTimeStamp}

	if !msg.SourceIPAddress.IsZeroValue() {
		ies = append(ies, msg.SourceIPAddress)
	}

	err := pfcp.sendPfcpMessage(header, ies)
	if err != nil {
		return fmt.Errorf("error sending PFCP Heartbeat Request: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendHeartbeatResponse(msg messages.HeartbeatResponse, sequenceNumber uint32) error {
	header := messages.NewNodePFCPHeader(messages.HeartbeatResponseMessageType, sequenceNumber)
	ies := []ie.InformationElement{msg.RecoveryTimeStamp}
	err := pfcp.sendPfcpMessage(header, ies)
	if err != nil {
		return fmt.Errorf("error sending PFCP Heartbeat Response: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPAssociationSetupRequest(msg messages.PFCPAssociationSetupRequest, sequenceNumber uint32) error {
	header := messages.NewNodePFCPHeader(messages.PFCPAssociationSetupRequestMessageType, sequenceNumber)
	ies := []ie.InformationElement{msg.NodeID, msg.RecoveryTimeStamp}

	if !msg.UPFunctionFeatures.IsZeroValue() {
		ies = append(ies, msg.UPFunctionFeatures)
	}

	err := pfcp.sendPfcpMessage(header, ies)
	if err != nil {
		return fmt.Errorf("error sending PFCP Association Setup Request: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPAssociationSetupResponse(msg messages.PFCPAssociationSetupResponse, sequenceNumber uint32) error {
	header := messages.NewNodePFCPHeader(messages.PFCPAssociationSetupResponseMessageType, sequenceNumber)
	ies := []ie.InformationElement{msg.NodeID, msg.Cause, msg.RecoveryTimeStamp}
	err := pfcp.sendPfcpMessage(header, ies)
	if err != nil {
		return fmt.Errorf("error sending PFCP Association Setup Response: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPAssociationUpdateRequest(msg messages.PFCPAssociationUpdateRequest, sequenceNumber uint32) error {
	header := messages.NewNodePFCPHeader(messages.PFCPAssociationUpdateRequestMessageType, sequenceNumber)
	ies := []ie.InformationElement{msg.NodeID}
	err := pfcp.sendPfcpMessage(header, ies)
	if err != nil {
		return fmt.Errorf("error sending PFCP Association Update Request: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPAssociationUpdateResponse(msg messages.PFCPAssociationUpdateResponse, sequenceNumber uint32) error {
	header := messages.NewNodePFCPHeader(messages.PFCPAssociationUpdateResponseMessageType, sequenceNumber)
	ies := []ie.InformationElement{msg.NodeID, msg.Cause}
	err := pfcp.sendPfcpMessage(header, ies)
	if err != nil {
		return fmt.Errorf("error sending PFCP Association Update Response: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPAssociationReleaseRequest(msg messages.PFCPAssociationReleaseRequest, sequenceNumber uint32) error {
	header := messages.NewNodePFCPHeader(messages.PFCPAssociationReleaseRequestMessageType, sequenceNumber)
	ies := []ie.InformationElement{msg.NodeID}
	err := pfcp.sendPfcpMessage(header, ies)
	if err != nil {
		return fmt.Errorf("error sending PFCP Association Release Request: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPAssociationReleaseResponse(msg messages.PFCPAssociationReleaseResponse, sequenceNumber uint32) error {
	header := messages.NewNodePFCPHeader(messages.PFCPAssociationReleaseResponseMessageType, sequenceNumber)
	ies := []ie.InformationElement{msg.NodeID, msg.Cause}
	err := pfcp.sendPfcpMessage(header, ies)
	if err != nil {
		return fmt.Errorf("error sending PFCP Association Release Response: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPNodeReportRequest(msg messages.PFCPNodeReportRequest, sequenceNumber uint32) error {
	header := messages.NewNodePFCPHeader(messages.PFCPNodeReportRequestMessageType, sequenceNumber)
	ies := []ie.InformationElement{msg.NodeID, msg.NodeReportType}
	err := pfcp.sendPfcpMessage(header, ies)
	if err != nil {
		return fmt.Errorf("error sending PFCP Node Report Request: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPNodeReportResponse(msg messages.PFCPNodeReportResponse, sequenceNumber uint32) error {
	header := messages.NewNodePFCPHeader(messages.PFCPNodeReportResponseMessageType, sequenceNumber)
	ies := []ie.InformationElement{msg.NodeID, msg.Cause}
	err := pfcp.sendPfcpMessage(header, ies)
	if err != nil {
		return fmt.Errorf("error sending PFCP Node Report Response: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPSessionEstablishmentRequest(msg messages.PFCPSessionEstablishmentRequest, seid uint64, sequenceNumber uint32) error {
	header := messages.NewSessionPFCPHeader(messages.PFCPSessionEstablishmentRequestMessageType, seid, sequenceNumber)
	ies := []ie.InformationElement{msg.NodeID, msg.CPFSEID, msg.CreatePDR, msg.CreateFAR}
	err := pfcp.sendPfcpMessage(header, ies)
	if err != nil {
		return fmt.Errorf("error sending PFCP Session Establishment Request: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPSessionEstablishmentResponse(msg messages.PFCPSessionEstablishmentResponse, seid uint64, sequenceNumber uint32) error {
	header := messages.NewSessionPFCPHeader(messages.PFCPSessionEstablishmentResponseMessageType, seid, sequenceNumber)
	ies := []ie.InformationElement{msg.NodeID, msg.Cause}
	err := pfcp.sendPfcpMessage(header, ies)
	if err != nil {
		return fmt.Errorf("error sending PFCP Session Establishment Response: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPSessionDeletionRequest(msg messages.PFCPSessionDeletionRequest, seid uint64, sequenceNumber uint32) error {
	header := messages.NewSessionPFCPHeader(messages.PFCPSessionDeletionRequestMessageType, seid, sequenceNumber)
	ies := []ie.InformationElement{}
	err := pfcp.sendPfcpMessage(header, ies)
	if err != nil {
		return fmt.Errorf("error sending PFCP Session Deletion Request: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPSessionDeletionResponse(msg messages.PFCPSessionDeletionResponse, seid uint64, sequenceNumber uint32) error {
	header := messages.NewSessionPFCPHeader(messages.PFCPSessionDeletionResponseMessageType, seid, sequenceNumber)
	ies := []ie.InformationElement{msg.Cause}
	err := pfcp.sendPfcpMessage(header, ies)
	if err != nil {
		return fmt.Errorf("error sending PFCP Session Deletion Response: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPSessionReportRequest(msg messages.PFCPSessionReportRequest, seid uint64, sequenceNumber uint32) error {
	header := messages.NewSessionPFCPHeader(messages.PFCPSessionReportRequestMessageType, seid, sequenceNumber)
	ies := []ie.InformationElement{msg.ReportType}
	err := pfcp.sendPfcpMessage(header, ies)
	if err != nil {
		return fmt.Errorf("error sending PFCP Session Report Request: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPSessionReportResponse(msg messages.PFCPSessionReportResponse, seid uint64, sequenceNumber uint32) error {
	header := messages.NewSessionPFCPHeader(messages.PFCPSessionReportResponseMessageType, seid, sequenceNumber)
	ies := []ie.InformationElement{msg.Cause}
	err := pfcp.sendPfcpMessage(header, ies)
	if err != nil {
		return fmt.Errorf("error sending PFCP Session Report Response: %w", err)
	}
	return nil
}
