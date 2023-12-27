package messages

import "github.com/dot-5g/pfcp/ie"

type PFCPAssociationReleaseRequest struct {
	SequenceNumber uint32
	NodeID         ie.NodeID
}

type PFCPAssociationReleaseResponse struct {
	SequenceNumber uint32
	NodeID         ie.NodeID
	Cause          ie.Cause
}

func ParsePFCPAssociationReleaseRequest(data []byte) (PFCPAssociationReleaseRequest, error) {
	ies, err := ie.ParseInformationElements(data)
	var nodeID ie.NodeID
	for _, elem := range ies {
		if nodeIDIE, ok := elem.(ie.NodeID); ok {
			nodeID = nodeIDIE
			continue
		}
	}

	return PFCPAssociationReleaseRequest{
		NodeID: nodeID,
	}, err
}

func ParsePFCPAssociationReleaseResponse(data []byte) (PFCPAssociationReleaseResponse, error) {
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

	return PFCPAssociationReleaseResponse{
		NodeID: nodeID,
		Cause:  cause,
	}, err
}
