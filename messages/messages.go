package messages

import (
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
