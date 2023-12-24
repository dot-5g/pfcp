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
	header.MessageLength = uint16(4 + len(payload))
	headerBytes := messages.SerializePFCPHeader(header)
	return append(headerBytes, payload...)
}

func (pfcp *Pfcp) SendHeartbeatRequest(recoveryTimeStamp ie.RecoveryTimeStamp, sequenceNumber uint32) (ie.RecoveryTimeStamp, error) {
	header := messages.NewPFCPHeader(messages.PFCPHeartbeatRequest, sequenceNumber)
	payload := []ie.InformationElement{recoveryTimeStamp}
	err := pfcp.sendPfcpMessage(header, payload)
	if err != nil {
		return recoveryTimeStamp, fmt.Errorf("error sending PFCP Heartbeat Request: %w", err)
	}
	return recoveryTimeStamp, nil
}

func (pfcp *Pfcp) SendHeartbeatResponse(recoveryTimeStamp ie.RecoveryTimeStamp, sequenceNumber uint32) (ie.RecoveryTimeStamp, error) {
	header := messages.NewPFCPHeader(messages.PFCPHeartbeatResponse, sequenceNumber)
	payload := []ie.InformationElement{recoveryTimeStamp}
	err := pfcp.sendPfcpMessage(header, payload)
	if err != nil {
		return recoveryTimeStamp, fmt.Errorf("error sending PFCP Heartbeat Response: %w", err)
	}
	return recoveryTimeStamp, nil
}

func (pfcp *Pfcp) SendPFCPAssociationSetupRequest(nodeID ie.NodeID, recoveryTimeStamp ie.RecoveryTimeStamp, sequenceNumber uint32) error {
	header := messages.NewPFCPHeader(messages.PFCPAssociationSetupRequest, sequenceNumber)
	payload := []ie.InformationElement{nodeID, recoveryTimeStamp}
	err := pfcp.sendPfcpMessage(header, payload)
	if err != nil {
		return fmt.Errorf("error sending PFCP Association Setup Request: %w", err)
	}
	return nil
}
