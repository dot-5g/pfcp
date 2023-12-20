package messages

import (
	"bytes"
	"encoding/binary"
)

type PFCPHeader struct {
	Version        byte
	MessageType    byte
	MessageLength  uint16
	SequenceNumber uint32
}

func SerializePFCPHeader(header PFCPHeader) []byte {
	buf := new(bytes.Buffer)

	// Octet 1: Version (3 bits), Spare (3 bits), FO (1 bit set to 0), MP (1 bit set to 0), S (1 bit set to 0)
	firstOctet := (header.Version << 5)
	buf.WriteByte(firstOctet)

	// Octet 2: Message Type (1 byte)
	buf.WriteByte(header.MessageType)

	// Octets 3 and 4: Message Length (2 bytes)
	binary.Write(buf, binary.BigEndian, header.MessageLength)

	// Octets 5, 6, and 7: Sequence Number (3 bytes)
	seqNumBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(seqNumBytes, header.SequenceNumber)
	buf.Write(seqNumBytes[0:3]) // Only write the first 3 bytes

	// Octet 8: Spare (1 byte set to 0)
	buf.WriteByte(0)

	return buf.Bytes()
}

// NewPFCPHeader creates a new PFCPHeader with the given message type and sequence number.
func NewPFCPHeader(messageType byte, sequenceNumber uint32) PFCPHeader {
	return PFCPHeader{
		Version:        1, // Assuming the version is 1
		MessageType:    messageType,
		MessageLength:  0, // To be set later
		SequenceNumber: sequenceNumber,
	}
}

func ParsePFCPHeader(data []byte) PFCPHeader {

	header := PFCPHeader{}
	header.Version = data[0] >> 5
	header.MessageType = data[1]
	header.MessageLength = binary.BigEndian.Uint16(data[2:4])

	seqNumBytes := make([]byte, 4)
	copy(seqNumBytes, data[4:7])
	header.SequenceNumber = binary.BigEndian.Uint32(seqNumBytes)

	return header
}
