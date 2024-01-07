package messages

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/dot-5g/pfcp/ie"
)

type MessageType byte

const (
	HeartbeatRequestMessageType                 MessageType = 1
	HeartbeatResponseMessageType                MessageType = 2
	PFCPAssociationSetupRequestMessageType      MessageType = 5
	PFCPAssociationSetupResponseMessageType     MessageType = 6
	PFCPAssociationUpdateRequestMessageType     MessageType = 7
	PFCPAssociationUpdateResponseMessageType    MessageType = 8
	PFCPAssociationReleaseRequestMessageType    MessageType = 9
	PFCPAssociationReleaseResponseMessageType   MessageType = 10
	PFCPNodeReportRequestMessageType            MessageType = 12
	PFCPNodeReportResponseMessageType           MessageType = 13
	PFCPSessionEstablishmentRequestMessageType  MessageType = 50
	PFCPSessionEstablishmentResponseMessageType MessageType = 51
	PFCPSessionDeletionRequestMessageType       MessageType = 54
	PFCPSessionDeletionResponseMessageType      MessageType = 55
	PFCPSessionReportRequestMessageType         MessageType = 56
	PFCPSessionReportResponseMessageType        MessageType = 57
)

type PFCPMessage interface {
	GetIEs() []ie.InformationElement
	GetMessageType() MessageType
	GetMessageTypeString() string
}

type DeserializerFunc func([]byte) (PFCPMessage, error)

var messageTypeDeserializers = map[MessageType]DeserializerFunc{
	HeartbeatRequestMessageType:                 DeserializeHeartbeatRequest,
	HeartbeatResponseMessageType:                DeserializeHeartbeatResponse,
	PFCPAssociationSetupRequestMessageType:      DeserializePFCPAssociationSetupRequest,
	PFCPAssociationSetupResponseMessageType:     DeserializePFCPAssociationSetupResponse,
	PFCPAssociationUpdateRequestMessageType:     DeserializePFCPAssociationUpdateRequest,
	PFCPAssociationUpdateResponseMessageType:    DeserializePFCPAssociationUpdateResponse,
	PFCPAssociationReleaseRequestMessageType:    DeserializePFCPAssociationReleaseRequest,
	PFCPAssociationReleaseResponseMessageType:   DeserializePFCPAssociationReleaseResponse,
	PFCPNodeReportRequestMessageType:            DeserializePFCPNodeReportRequest,
	PFCPNodeReportResponseMessageType:           DeserializePFCPNodeReportResponse,
	PFCPSessionEstablishmentRequestMessageType:  DeserializePFCPSessionEstablishmentRequest,
	PFCPSessionEstablishmentResponseMessageType: DeserializePFCPSessionEstablishmentResponse,
	PFCPSessionDeletionRequestMessageType:       DeserializePFCPSessionDeletionRequest,
	PFCPSessionDeletionResponseMessageType:      DeserializePFCPSessionDeletionResponse,
	PFCPSessionReportRequestMessageType:         DeserializePFCPSessionReportRequest,
	PFCPSessionReportResponseMessageType:        DeserializePFCPSessionReportResponse,
}

type PFCPHeader struct {
	Version        byte
	FO             bool
	MP             bool
	S              bool
	MessageType    MessageType
	MessageLength  uint16
	SEID           uint64
	SequenceNumber uint32
}

func NewNodePFCPHeader(messageType MessageType, sequenceNumber uint32) PFCPHeader {
	var version byte = 1
	var fo bool = false
	var mp bool = false
	var s bool = false
	var messageLength uint16 = 0 // To be set later

	return PFCPHeader{
		Version:        version,
		FO:             fo,
		MP:             mp,
		S:              s,
		MessageType:    messageType,
		MessageLength:  messageLength,
		SequenceNumber: sequenceNumber,
	}
}

func NewSessionPFCPHeader(messageType MessageType, seid uint64, sequenceNumber uint32) PFCPHeader {
	var version byte = 1
	var fo bool = false
	var mp bool = false
	var s bool = true
	var messageLength uint16 = 0 // To be set later

	return PFCPHeader{
		Version:        version,
		FO:             fo,
		MP:             mp,
		S:              s,
		MessageType:    messageType,
		MessageLength:  messageLength,
		SEID:           seid,
		SequenceNumber: sequenceNumber,
	}
}

func (header PFCPHeader) Serialize() []byte {
	// if S = 0, SEID field is not present, k = 0, m = 0 and n = 5;
	// if S = 1, SEID field is present, k = 1, m = 5 and n = 13.
	buf := new(bytes.Buffer)

	// Octet 1: Version (3 bits), Spare (2 bits), FO (1 bit), MP (1 bit), S (1 bit)
	firstOctet := (header.Version << 5)
	if header.FO {
		firstOctet |= 1 << 2 // Set the FO bit
	}
	if header.MP {
		firstOctet |= 1 << 1 // Set the MP bit
	}
	if header.S {
		firstOctet |= 1 // Set the S bit
	}
	buf.WriteByte(firstOctet)

	// Octet 2: Message Type (1 byte)
	buf.WriteByte(byte(header.MessageType))

	// Octets 3 to 4: Message Length (2 bytes)
	binary.Write(buf, binary.BigEndian, header.MessageLength)

	// Octets m to k(m+7): SEID (8 bytes)
	if header.S {
		binary.Write(buf, binary.BigEndian, header.SEID)
	}

	// Octets n to (n+2): Sequence Number (4 bytes)
	seqNumBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(seqNumBytes, header.SequenceNumber)
	buf.Write(seqNumBytes[1:])

	// Octet 8: Spare (1 byte set to 0)
	buf.WriteByte(0)

	return buf.Bytes()
}

func Serialize(message PFCPMessage, header PFCPHeader) []byte {
	var payload []byte
	ies := message.GetIEs()
	fmt.Printf("ies: %v\n", ies)
	for _, element := range ies {
		fmt.Printf("IE: %v\n", element)
		// fmt.Printf("IE type: %v\n", element.Type)
		payload = append(payload, element.Serialize()...)
	}
	header.MessageLength = uint16(len(payload))
	headerBytes := header.Serialize()
	return append(headerBytes, payload...)
}

func DeserializePFCPHeader(data []byte) (PFCPHeader, error) {
	const baseHeaderSize = 8                            // Base size for node-related messages
	const seidSize = 8                                  // Size of SEID field
	const sessionHeaderSize = baseHeaderSize + seidSize // Total size for session-related messages

	if len(data) < baseHeaderSize {
		return PFCPHeader{}, fmt.Errorf("expected at least %d bytes, got %d", baseHeaderSize, len(data))
	}

	header := PFCPHeader{}
	header.Version = data[0] >> 5
	header.FO = (data[0] & 0x04) != 0 // Extract the FO bit
	header.MP = (data[0] & 0x02) != 0 // Extract the MP bit
	header.S = (data[0] & 0x01) != 0  // Extract the S bit
	header.MessageType = MessageType(data[1])
	header.MessageLength = binary.BigEndian.Uint16(data[2:4])

	// For node-related messages, sequence number starts at offset 4
	// For session-related messages, SEID is between offsets 4-11, and sequence number starts at offset 12
	var seqNumOffset int
	if header.S {
		if len(data) < sessionHeaderSize {
			return PFCPHeader{}, fmt.Errorf("expected %d bytes for session message, got %d", sessionHeaderSize, len(data))
		}
		header.SEID = binary.BigEndian.Uint64(data[4:12])
		seqNumOffset = 12
	} else {
		seqNumOffset = 4
	}

	// Extract the sequence number
	if len(data) < seqNumOffset+3 {
		return PFCPHeader{}, fmt.Errorf("insufficient data for sequence number")
	}
	seqNumBytes := []byte{0, data[seqNumOffset], data[seqNumOffset+1], data[seqNumOffset+2]}
	header.SequenceNumber = binary.BigEndian.Uint32(seqNumBytes)

	return header, nil
}

func DeserializePFCPMessage(payload []byte) (PFCPHeader, PFCPMessage, error) {
	header, err := DeserializePFCPHeader(payload)
	if err != nil {
		return header, nil, err
	}

	payloadOffset := 8
	if header.S {
		payloadOffset = 16
	}

	if len(payload) < payloadOffset {
		return header, nil, fmt.Errorf("insufficient data for payload message")
	}
	payloadMessage := payload[payloadOffset:]
	if deserializer, exists := messageTypeDeserializers[header.MessageType]; exists {
		msg, err := deserializer(payloadMessage)
		if err != nil {
			return header, nil, fmt.Errorf("error deserializing payload message: %v", err)
		}
		return header, msg, nil
	}

	return header, nil, fmt.Errorf("unsupported message type: %d", header.MessageType)
}
