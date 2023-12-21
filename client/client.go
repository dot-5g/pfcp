package client

import (
	"log"
	"time"

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

func (pfcp *Pfcp) sendPfcpMessage(header messages.PFCPHeader, payload []byte, messageType string) error {
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

func (pfcp *Pfcp) SendHeartbeatRequest(time time.Time) (messages.RecoveryTimeStamp, error) {
	// Create a RecoveryTimeStamp with the current time
	recoveryTimeStamp := messages.NewRecoveryTimeStamp(time)
	timeBytes := recoveryTimeStamp.ToBytes()
	header := messages.NewPFCPHeader(1, 1)
	err := pfcp.sendPfcpMessage(header, timeBytes, "Heartbeat Request")
	return recoveryTimeStamp, err
}
