package messages

type MessageType byte

const (
	HeartbeatRequestMessageType              MessageType = 1
	HeartbeatResponseMessageType             MessageType = 2
	PFCPAssociationSetupRequestMessageType   MessageType = 5
	PFCPAssociationSetupResponseMessageType  MessageType = 6
	PFCPAssociationUpdateRequestMessageType  MessageType = 7
	PFCPAssociationUpdateResponseMessageType MessageType = 8
)
