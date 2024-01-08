package messages

import (
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

func DeserializePFCPMessage(payload []byte) (MessageHeader, PFCPMessage, error) {
	header, err := DeserializeMessageHeader(payload)
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
