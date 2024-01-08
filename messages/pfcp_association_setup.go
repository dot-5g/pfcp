package messages

import (
	"github.com/dot-5g/pfcp/ie"
)

type PFCPAssociationSetupRequest struct {
	Header             Header
	NodeID             ie.NodeID             // Mandatory
	RecoveryTimeStamp  ie.RecoveryTimeStamp  // Mandatory
	UPFunctionFeatures ie.UPFunctionFeatures // Conditional
}

type PFCPAssociationSetupResponse struct {
	Header            Header
	NodeID            ie.NodeID            // Mandatory
	Cause             ie.Cause             // Mandatory
	RecoveryTimeStamp ie.RecoveryTimeStamp // Mandatory
}

func (msg PFCPAssociationSetupRequest) GetIEs() []ie.InformationElement {
	ies := []ie.InformationElement{msg.NodeID, msg.RecoveryTimeStamp}

	if !msg.UPFunctionFeatures.IsZeroValue() {
		ies = append(ies, msg.UPFunctionFeatures)
	}

	return ies
}

func (msg PFCPAssociationSetupResponse) GetIEs() []ie.InformationElement {
	return []ie.InformationElement{msg.NodeID, msg.Cause, msg.RecoveryTimeStamp}
}

func (msg PFCPAssociationSetupRequest) GetMessageType() MessageType {
	return PFCPAssociationSetupRequestMessageType
}

func (msg PFCPAssociationSetupResponse) GetMessageType() MessageType {
	return PFCPAssociationSetupResponseMessageType
}

func (msg PFCPAssociationSetupRequest) GetMessageTypeString() string {
	return "PFCP Association Setup Request"
}

func (msg PFCPAssociationSetupResponse) GetMessageTypeString() string {
	return "PFCP Association Setup Response"
}

func (msg *PFCPAssociationSetupRequest) SetHeader(h Header) {
	msg.Header = h
}

func (msg *PFCPAssociationSetupResponse) SetHeader(h Header) {
	msg.Header = h
}

func (msg PFCPAssociationSetupRequest) GetHeader() Header {
	return msg.Header
}

func (msg PFCPAssociationSetupResponse) GetHeader() Header {
	return msg.Header
}

func DeserializePFCPAssociationSetupRequest(data []byte) (PFCPAssociationSetupRequest, error) {
	ies, err := ie.DeserializeInformationElements(data)
	var nodeID ie.NodeID
	var recoveryTimeStamp ie.RecoveryTimeStamp
	var upfeatures ie.UPFunctionFeatures
	for _, elem := range ies {
		if tsIE, ok := elem.(ie.RecoveryTimeStamp); ok {
			recoveryTimeStamp = tsIE
			continue
		}
		if nodeIDIE, ok := elem.(ie.NodeID); ok {
			nodeID = nodeIDIE
			continue
		}
		if upfeaturesIE, ok := elem.(ie.UPFunctionFeatures); ok {
			upfeatures = upfeaturesIE
			continue
		}
	}

	return PFCPAssociationSetupRequest{
		NodeID:             nodeID,
		RecoveryTimeStamp:  recoveryTimeStamp,
		UPFunctionFeatures: upfeatures,
	}, err
}

func DeserializePFCPAssociationSetupResponse(data []byte) (PFCPAssociationSetupResponse, error) {
	ies, err := ie.DeserializeInformationElements(data)
	var nodeID ie.NodeID
	var cause ie.Cause
	var recoveryTimeStamp ie.RecoveryTimeStamp
	for _, elem := range ies {
		if tsIE, ok := elem.(ie.RecoveryTimeStamp); ok {
			recoveryTimeStamp = tsIE
			continue
		}
		if nodeIDIE, ok := elem.(ie.NodeID); ok {
			nodeID = nodeIDIE
			continue
		}
		if causeIE, ok := elem.(ie.Cause); ok {
			cause = causeIE
			continue
		}
	}

	return PFCPAssociationSetupResponse{
		NodeID:            nodeID,
		Cause:             cause,
		RecoveryTimeStamp: recoveryTimeStamp,
	}, err
}
