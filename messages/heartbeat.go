package messages

import (
	"github.com/dot-5g/pfcp/ie"
)

type HeartbeatRequest struct {
	RecoveryTimeStamp ie.RecoveryTimeStamp // Mandatory
	SourceIPAddress   ie.SourceIPAddress   // Optional
}

type HeartbeatResponse struct {
	RecoveryTimeStamp ie.RecoveryTimeStamp // Mandatory
}

func DeserializeHeartbeatRequest(data []byte) (PFCPMessage, error) {
	ies, err := ie.ParseInformationElements(data)
	var recoveryTimeStamp ie.RecoveryTimeStamp
	var sourceIPAddress ie.SourceIPAddress
	for _, elem := range ies {
		if tsIE, ok := elem.(ie.RecoveryTimeStamp); ok {
			recoveryTimeStamp = tsIE
			continue
		}
		if ipIE, ok := elem.(ie.SourceIPAddress); ok {
			sourceIPAddress = ipIE
			continue
		}
	}

	return HeartbeatRequest{
		RecoveryTimeStamp: recoveryTimeStamp,
		SourceIPAddress:   sourceIPAddress,
	}, err
}

func DeserializeHeartbeatResponse(data []byte) (PFCPMessage, error) {
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
