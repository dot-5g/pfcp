package messages

import "github.com/dot-5g/pfcp/ie"

type PFCPAssociationUpdateRequest struct {
	SequenceNumber uint32
	NodeID         ie.NodeID
}

type PFCPAssociationUpdateResponse struct {
	SequenceNumber uint32
	NodeID         ie.NodeID
	Cause          ie.Cause
}

func NewPFCPAssociationUpdateRequest(nodeID ie.NodeID) PFCPAssociationUpdateRequest {
	return PFCPAssociationUpdateRequest{
		NodeID: nodeID,
	}
}

func NewPFCPAssociationUpdateResponse(nodeID ie.NodeID, cause ie.Cause) PFCPAssociationUpdateResponse {
	return PFCPAssociationUpdateResponse{
		NodeID: nodeID,
		Cause:  cause,
	}
}

func ParsePFCPAssociationUpdateRequest(data []byte) (PFCPAssociationUpdateRequest, error) {
	ies, err := ie.ParseInformationElements(data)
	var nodeID ie.NodeID
	for _, elem := range ies {
		if nodeIDIE, ok := elem.(ie.NodeID); ok {
			nodeID = nodeIDIE
			continue
		}
	}

	return PFCPAssociationUpdateRequest{
		NodeID: nodeID,
	}, err
}

func ParsePFCPAssociationUpdateResponse(data []byte) (PFCPAssociationUpdateResponse, error) {
	ies, err := ie.ParseInformationElements(data)
	var nodeID ie.NodeID
	var cause ie.Cause
	for _, elem := range ies {
		if nodeIDIE, ok := elem.(ie.NodeID); ok {
			nodeID = nodeIDIE
			continue
		}
		if causeIE, ok := elem.(ie.Cause); ok {
			cause = causeIE
			continue
		}
	}

	return PFCPAssociationUpdateResponse{
		NodeID: nodeID,
		Cause:  cause,
	}, err
}
