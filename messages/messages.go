package messages

type MessageType byte

const (
	HeartbeatRequestMessageType                MessageType = 1
	HeartbeatResponseMessageType               MessageType = 2
	PFCPAssociationSetupRequestMessageType     MessageType = 5
	PFCPAssociationSetupResponseMessageType    MessageType = 6
	PFCPAssociationUpdateRequestMessageType    MessageType = 7
	PFCPAssociationUpdateResponseMessageType   MessageType = 8
	PFCPAssociationReleaseRequestMessageType   MessageType = 9
	PFCPAssociationReleaseResponseMessageType  MessageType = 10
	PFCPNodeReportRequestMessageType           MessageType = 12
	PFCPNodeReportResponseMessageType          MessageType = 13
	PFCPSessionEstablishmentRequestMessageType MessageType = 50
)
