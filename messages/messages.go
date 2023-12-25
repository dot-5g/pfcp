package messages

type PFCPMessageType int

const (
	PFCPHeartbeatRequest PFCPMessageType = iota
	PFCPHeartbeatResponse
	PFCPAssociationSetupRequest
	PFCPAssociationSetupResponse
)

func MessageTypeToByte(messageType PFCPMessageType) byte {
	switch messageType {
	case PFCPHeartbeatRequest:
		return 1
	case PFCPHeartbeatResponse:
		return 2
	case PFCPAssociationSetupRequest:
		return 5
	case PFCPAssociationSetupResponse:
		return 6
	default:
		return 0
	}
}
