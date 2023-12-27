package client

import (
	"fmt"
	"log"

	"github.com/dot-5g/pfcp/headers"
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

func (pfcp *Pfcp) sendPfcpMessage(header headers.PFCPHeader, elements []ie.InformationElement) error {
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

func serializeMessage(header headers.PFCPHeader, payload []byte) []byte {
	header.MessageLength = uint16(4 + len(payload))
	headerBytes := headers.SerializePFCPHeader(header)
	return append(headerBytes, payload...)
}

func (pfcp *Pfcp) SendHeartbeatRequest(msg messages.HeartbeatRequest, sequenceNumber uint32) (ie.RecoveryTimeStamp, error) {
	header := headers.NewPFCPHeader(messages.HeartbeatRequestMessageType, sequenceNumber)
	payload := []ie.InformationElement{msg.RecoveryTimeStamp}
	err := pfcp.sendPfcpMessage(header, payload)
	if err != nil {
		return msg.RecoveryTimeStamp, fmt.Errorf("error sending PFCP Heartbeat Request: %w", err)
	}
	return msg.RecoveryTimeStamp, nil
}

func (pfcp *Pfcp) SendHeartbeatResponse(msg messages.HeartbeatResponse, sequenceNumber uint32) (ie.RecoveryTimeStamp, error) {
	header := headers.NewPFCPHeader(messages.HeartbeatResponseMessageType, sequenceNumber)
	payload := []ie.InformationElement{msg.RecoveryTimeStamp}
	err := pfcp.sendPfcpMessage(header, payload)
	if err != nil {
		return msg.RecoveryTimeStamp, fmt.Errorf("error sending PFCP Heartbeat Response: %w", err)
	}
	return msg.RecoveryTimeStamp, nil
}

func (pfcp *Pfcp) SendPFCPAssociationSetupRequest(msg messages.PFCPAssociationSetupRequest, sequenceNumber uint32) error {
	header := headers.NewPFCPHeader(messages.PFCPAssociationSetupRequestMessageType, sequenceNumber)
	payload := []ie.InformationElement{msg.NodeID, msg.RecoveryTimeStamp}
	err := pfcp.sendPfcpMessage(header, payload)
	if err != nil {
		return fmt.Errorf("error sending PFCP Association Setup Request: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPAssociationSetupResponse(msg messages.PFCPAssociationSetupResponse, sequenceNumber uint32) error {
	header := headers.NewPFCPHeader(messages.PFCPAssociationSetupResponseMessageType, sequenceNumber)
	payload := []ie.InformationElement{msg.NodeID, msg.Cause, msg.RecoveryTimeStamp}
	err := pfcp.sendPfcpMessage(header, payload)
	if err != nil {
		return fmt.Errorf("error sending PFCP Association Setup Response: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPAssociationUpdateRequest(msg messages.PFCPAssociationUpdateRequest, sequenceNumber uint32) error {
	header := headers.NewPFCPHeader(messages.PFCPAssociationUpdateRequestMessageType, sequenceNumber)
	payload := []ie.InformationElement{msg.NodeID}
	err := pfcp.sendPfcpMessage(header, payload)
	if err != nil {
		return fmt.Errorf("error sending PFCP Association Update Request: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPAssociationUpdateResponse(msg messages.PFCPAssociationUpdateResponse, sequenceNumber uint32) error {
	header := headers.NewPFCPHeader(messages.PFCPAssociationUpdateResponseMessageType, sequenceNumber)
	payload := []ie.InformationElement{msg.NodeID, msg.Cause}
	err := pfcp.sendPfcpMessage(header, payload)
	if err != nil {
		return fmt.Errorf("error sending PFCP Association Update Response: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPAssociationReleaseRequest(msg messages.PFCPAssociationReleaseRequest, sequenceNumber uint32) error {
	header := headers.NewPFCPHeader(messages.PFCPAssociationReleaseRequestMessageType, sequenceNumber)
	payload := []ie.InformationElement{msg.NodeID}
	err := pfcp.sendPfcpMessage(header, payload)
	if err != nil {
		return fmt.Errorf("error sending PFCP Association Release Request: %w", err)
	}
	return nil
}

func (pfcp *Pfcp) SendPFCPAssociationReleaseResponse(msg messages.PFCPAssociationReleaseResponse, sequenceNumber uint32) error {
	header := headers.NewPFCPHeader(messages.PFCPAssociationReleaseResponseMessageType, sequenceNumber)
	payload := []ie.InformationElement{msg.NodeID, msg.Cause}
	err := pfcp.sendPfcpMessage(header, payload)
	if err != nil {
		return fmt.Errorf("error sending PFCP Association Release Response: %w", err)
	}
	return nil
}
