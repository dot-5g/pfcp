package messages

import "github.com/dot-5g/pfcp/ie"

type PFCPAssociationSetupRequest struct {
	MessageType       MessageType
	SequenceNumber    uint32
	NodeID            ie.NodeID
	RecoveryTimeStamp ie.RecoveryTimeStamp
}

type PFCPAssociationSetupResponse struct {
	MessageType       MessageType
	SequenceNumber    uint32
	NodeID            ie.NodeID
	Cause             ie.Cause
	RecoveryTimeStamp ie.RecoveryTimeStamp
}

func ParsePFCPAssociationSetupRequest(data []byte) (PFCPAssociationSetupRequest, error) {
	ies, err := ie.ParseInformationElements(data)
	var nodeID ie.NodeID
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
	}

	return PFCPAssociationSetupRequest{
		NodeID:            nodeID,
		RecoveryTimeStamp: recoveryTimeStamp,
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
