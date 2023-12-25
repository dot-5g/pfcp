package headers_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/dot-5g/pfcp/headers"
)

func TestGivenPfcpHeaderWhenSerializePFCPHeaderThenSerializedCorrectly(t *testing.T) {
	pfcpHeader := headers.PFCPHeader{
		Version:        1,
		MessageType:    2,
		MessageLength:  3,
		SequenceNumber: 4,
	}

	headerBytes := headers.SerializePFCPHeader(pfcpHeader)

	if len(headerBytes) != 8 {
		t.Errorf("Expected 8 bytes, got %d", len(headerBytes))
	}

	serializedVersion := headerBytes[0] >> 5
	if serializedVersion != 1 {
		t.Errorf("Expected version 1, got %d", serializedVersion)
	}

	serializedMessageType := headerBytes[1]
	if serializedMessageType != 2 {
		t.Errorf("Expected message type 2, got %d", serializedMessageType)
	}

	serializedMessageLength := binary.BigEndian.Uint16(headerBytes[2:4])
	if serializedMessageLength != 3 {
		t.Errorf("Expected message length 3, got %d", serializedMessageLength)
	}

	expectedSeqNum := []byte{0, 0, 4}
	serializedSequenceNumber := headerBytes[4:7]
	if !bytes.Equal(serializedSequenceNumber, expectedSeqNum) {
		t.Errorf("Expected sequence number %v, got %v", expectedSeqNum, serializedSequenceNumber)
	}
}
