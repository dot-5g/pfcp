package messages

import "github.com/dot-5g/pfcp/ie"

type PFCPSessionEstablishmentRequest struct {
	NodeID    ie.NodeID    // Mandatory
	CPFSEID   ie.FSEID     // Mandatory
	CreatePDR ie.CreatePDR // Mandatory
	CreateFAR ie.CreateFAR // Mandatory
}

func DeserializePFCPSessionEstablishmentRequest(data []byte) (PFCPMessage, error) {
	ies, err := ie.ParseInformationElements(data)
	var nodeID ie.NodeID
	var controlPlaneFSEID ie.FSEID
	var createPDR ie.CreatePDR
	var createFAR ie.CreateFAR

	for _, elem := range ies {
		if nodeIDIE, ok := elem.(ie.NodeID); ok {
			nodeID = nodeIDIE
			continue
		}
		if controlPlaneFSEIDIE, ok := elem.(ie.FSEID); ok {
			controlPlaneFSEID = controlPlaneFSEIDIE
			continue
		}
		if createPDRIE, ok := elem.(ie.CreatePDR); ok {
			createPDR = createPDRIE
			continue
		}
		if createFARIE, ok := elem.(ie.CreateFAR); ok {
			createFAR = createFARIE
			continue
		}

	}

	return PFCPSessionEstablishmentRequest{
		NodeID:    nodeID,
		CPFSEID:   controlPlaneFSEID,
		CreatePDR: createPDR,
		CreateFAR: createFAR,
	}, err
}
