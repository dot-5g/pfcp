package pfcp

import (
	"bytes"
	"encoding/binary"
)

type PFCPHeader struct {
	Version         byte
	MessageType     byte
	MessageLength   uint16
	SEID            uint64
	SequenceNumber  uint32
	MessagePriority byte
}

func SerializePFCPHeader(header PFCPHeader) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, header)
	return buf.Bytes()
}

func NewPFCPHeader(messageType byte, sequenceNumber uint32) PFCPHeader {
	return PFCPHeader{
		Version:         1,
		MessageType:     messageType,
		MessageLength:   0,
		SEID:            0,
		SequenceNumber:  sequenceNumber,
		MessagePriority: 0,
	}
}
