package headers

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/dot-5g/pfcp/messages"
)

const HeaderSize = 8

type PFCPHeader struct {
	Version        byte
	MessageType    messages.MessageType
	MessageLength  uint16
	SequenceNumber uint32
}

func SerializePFCPHeader(header PFCPHeader) []byte {
	buf := new(bytes.Buffer)

	// Octet 1: Version (3 bits), Spare (2 bits), FO (1 bit set to 0), MP (1 bit set to 0), S (1 bit set to 0)
	firstOctet := (header.Version << 5)
	buf.WriteByte(firstOctet)

	// Octet 2: Message Type (1 byte)
	buf.WriteByte(byte(header.MessageType))

	// Octets 3 and 4: Message Length (2 bytes)
	binary.Write(buf, binary.BigEndian, header.MessageLength)

	// Octets 5, 6, and 7: Sequence Number (3 bytes)
	seqNumBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(seqNumBytes, header.SequenceNumber)
	buf.Write(seqNumBytes[1:])

	// Octet 8: Spare (1 byte set to 0)
	buf.WriteByte(0)

	return buf.Bytes()
}

func NewPFCPHeader(messageType messages.MessageType, sequenceNumber uint32) PFCPHeader {
	return PFCPHeader{
		Version:        1,
		MessageType:    messageType,
		MessageLength:  0, // To be set later
		SequenceNumber: sequenceNumber,
	}
}

func ParsePFCPHeader(data []byte) (PFCPHeader, error) {
	if len(data) != HeaderSize {
		return PFCPHeader{}, fmt.Errorf("expected %d bytes, got %d", HeaderSize, len(data))
	}

	header := PFCPHeader{}
	header.Version = data[0] >> 5
	header.MessageType = messages.MessageType(data[1])
	header.MessageLength = binary.BigEndian.Uint16(data[2:4])

	seqNumBytes := []byte{0, data[4], data[5], data[6]}
	header.SequenceNumber = binary.BigEndian.Uint32(seqNumBytes)

	return header, nil
}
