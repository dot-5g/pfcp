package messages

import "github.com/dot-5g/pfcp/ie"

type PFCPAssociationUpdateRequest struct {
	NodeID ie.NodeID // Mandatory
}

type PFCPAssociationUpdateResponse struct {
	NodeID ie.NodeID // Mandatory
	Cause  ie.Cause  // Mandatory
}

func DeserializePFCPAssociationUpdateRequest(data []byte) (PFCPMessage, error) {
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

func DeserializePFCPAssociationUpdateResponse(data []byte) (PFCPMessage, error) {
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
