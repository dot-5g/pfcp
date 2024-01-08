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

func (msg HeartbeatRequest) GetIEs() []ie.InformationElement {
	ies := []ie.InformationElement{msg.RecoveryTimeStamp}
	if !msg.SourceIPAddress.IsZeroValue() {
		ies = append(ies, msg.SourceIPAddress)
	}
	return ies
}

func (msg HeartbeatResponse) GetIEs() []ie.InformationElement {
	return []ie.InformationElement{msg.RecoveryTimeStamp}
}

func (msg HeartbeatRequest) GetMessageType() MessageType {
	return HeartbeatRequestMessageType
}

func (msg HeartbeatResponse) GetMessageType() MessageType {
	return HeartbeatResponseMessageType
}

func (msg HeartbeatRequest) GetMessageTypeString() string {
	return "Heartbeat Request"
}

func (msg HeartbeatResponse) GetMessageTypeString() string {
	return "Heartbeat Response"
}

func DeserializeHeartbeatRequest(data []byte) (HeartbeatRequest, error) {
	ies, err := ie.DeserializeInformationElements(data)
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

func DeserializeHeartbeatResponse(data []byte) (HeartbeatResponse, error) {
	ies, err := ie.DeserializeInformationElements(data)
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
