package client

import (
	"log"

	"github.com/dot-5g/pfcp/information_elements"
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

func (pfcp *Pfcp) sendPfcpMessage(header messages.PFCPHeader, elements []information_elements.InformationElement, messageType string) error {
	var payload []byte
	for _, element := range elements {
		payload = append(payload, element.Serialize()...)
	}
	message := serializeMessage(header, payload)
	if err := pfcp.Udp.Send(message); err != nil {
		log.Printf("Failed to send PFCP %s: %v\n", messageType, err)
		return err
	}
	log.Printf("PFCP %s sent successfully to %s.\n", messageType, pfcp.ServerAddress)
	return nil
}

func serializeMessage(header messages.PFCPHeader, payload []byte) []byte {
	header.MessageLength = uint16(4 + len(payload))
	headerBytes := messages.SerializePFCPHeader(header)
	return append(headerBytes, payload...)
}

func (pfcp *Pfcp) SendHeartbeatRequest(recoveryTimeStamp information_elements.RecoveryTimeStamp, sequenceNumber uint32) (information_elements.RecoveryTimeStamp, error) {
	header := messages.NewPFCPHeader(messages.PFCPHeartbeatRequest, sequenceNumber)
	err := pfcp.sendPfcpMessage(header, []information_elements.InformationElement{recoveryTimeStamp}, "Heartbeat Request")
	return recoveryTimeStamp, err
}

func (pfcp *Pfcp) SendHeartbeatResponse(recoveryTimeStamp information_elements.RecoveryTimeStamp, sequenceNumber uint32) (information_elements.RecoveryTimeStamp, error) {
	header := messages.NewPFCPHeader(messages.PFCPHeartbeatResponse, sequenceNumber)
	err := pfcp.sendPfcpMessage(header, []information_elements.InformationElement{recoveryTimeStamp}, "Heartbeat Response")
	return recoveryTimeStamp, err
}
