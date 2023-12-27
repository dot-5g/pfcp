package messages

import (
	"github.com/dot-5g/pfcp/ie"
)

type HeartbeatRequest struct {
	MessageType       MessageType
	SequenceNumber    uint32
	RecoveryTimeStamp ie.RecoveryTimeStamp
}

type HeartbeatResponse struct {
	MessageType       MessageType
	SequenceNumber    uint32
	RecoveryTimeStamp ie.RecoveryTimeStamp
}

func ParseHeartbeatRequest(data []byte) (HeartbeatRequest, error) {
	ies, err := ie.ParseInformationElements(data)
	var recoveryTimeStamp ie.RecoveryTimeStamp
	for _, elem := range ies {
		if tsIE, ok := elem.(ie.RecoveryTimeStamp); ok {
			recoveryTimeStamp = tsIE
			continue
		}
	}

	return HeartbeatRequest{
		RecoveryTimeStamp: recoveryTimeStamp,
	}, err
}

func ParseHeartbeatResponse(data []byte) (HeartbeatResponse, error) {
	ies, err := ie.ParseInformationElements(data)
	var recoveryTimeStamp ie.RecoveryTimeStamp
	for _, elem := range ies {
		if tsIE, ok := elem.(ie.RecoveryTimeStamp); ok {
			recoveryTimeStamp = tsIE
			continue
		}
	}

	return HeartbeatResponse{
		RecoveryTimeStamp: recoveryTimeStamp,
	}, err
}
