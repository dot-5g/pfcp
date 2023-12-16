package pfcp

import (
	"encoding/binary"
	"log"
	"time"
)

type RecoveryTimeStamp time.Time

type HeartbeatRequest struct {
	RecoveryTimeStamp RecoveryTimeStamp
}

func (pfcp *Pfcp) sendPfcpMessage(header PFCPHeader, payload []byte, messageType string) error {
	message := serializeMessage(header, payload)
	if err := pfcp.Udp.Send(message); err != nil {
		log.Printf("Failed to send PFCP %s: %v\n", messageType, err)
		return err
	}
	log.Printf("PFCP %s sent successfully.\n", messageType)
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
