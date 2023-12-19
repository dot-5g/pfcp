package client

import (
	"encoding/binary"
	"log"
	"time"

	"github.com/dot-5g/pfcp/network"
)

type Pfcp struct {
	ServerAddress string
	Udp           network.UdpSender
}

type RecoveryTimeStamp time.Time

type HeartbeatRequest struct {
	RecoveryTimeStamp RecoveryTimeStamp
}

type HeartbeatResponse struct {
	RecoveryTimeStamp RecoveryTimeStamp
}

func New(ServerAddress string) *Pfcp {

	udpClient, err := network.NewUdp(ServerAddress)
	if err != nil {
		log.Printf("Failed to initialize UDP client: %v\n", err)
		return nil
	}

	return &Pfcp{ServerAddress: ServerAddress, Udp: udpClient}
}

func (pfcp *Pfcp) sendPfcpMessage(header PFCPHeader, payload []byte, messageType string) error {
	message := serializeMessage(header, payload)
	if err := pfcp.Udp.Send(message); err != nil {
		log.Printf("Failed to send PFCP %s: %v\n", messageType, err)
		return err
	}
	log.Printf("PFCP %s sent successfully to %s.\n", messageType, pfcp.ServerAddress)
	return nil
}

func serializeMessage(header PFCPHeader, payload []byte) []byte {
	headerBytes := SerializePFCPHeader(header)
	header.MessageLength = uint16(len(headerBytes) + len(payload))
	return append(headerBytes, payload...)
}

func (pfcp *Pfcp) SendHeartbeatRequest() error {
	request := HeartbeatRequest{RecoveryTimeStamp: RecoveryTimeStamp(time.Now())}
	requestBytes := request.ToBytes()
	header := NewPFCPHeader(1, 1)
	return pfcp.sendPfcpMessage(header, requestBytes, "Heartbeat Request")
}

func (p *Pfcp) SendHeartbeatResponse() error {
	response := HeartbeatRequest{RecoveryTimeStamp: RecoveryTimeStamp(time.Now())}
	responseBytes := response.ToBytes()
	header := NewPFCPHeader(2, 1)
	return p.sendPfcpMessage(header, responseBytes, "Heartbeat Response")
}

func (h HeartbeatRequest) ToBytes() []byte {
	timestamp := time.Time(h.RecoveryTimeStamp).Unix()
	timeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(timeBytes, uint64(timestamp))
	return timeBytes
}
