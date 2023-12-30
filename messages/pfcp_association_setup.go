package messages

import (
	"github.com/dot-5g/pfcp/ie"
)

type PFCPAssociationSetupRequest struct {
	NodeID             ie.NodeID             // Mandatory
	RecoveryTimeStamp  ie.RecoveryTimeStamp  // Mandatory
	UPFunctionFeatures ie.UPFunctionFeatures // Conditional
}

type PFCPAssociationSetupResponse struct {
	NodeID            ie.NodeID            // Mandatory
	Cause             ie.Cause             // Mandatory
	RecoveryTimeStamp ie.RecoveryTimeStamp // Mandatory
}

func ParsePFCPAssociationSetupRequest(data []byte) (PFCPAssociationSetupRequest, error) {
	ies, err := ie.ParseInformationElements(data)
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

func ParsePFCPAssociationSetupResponse(data []byte) (PFCPAssociationSetupResponse, error) {
	ies, err := ie.ParseInformationElements(data)
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
