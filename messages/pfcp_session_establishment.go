package messages

import "github.com/dot-5g/pfcp/ie"

type PFCPSessionEstablishmentRequest struct {
	NodeID    ie.NodeID    // Mandatory
	CPFSEID   ie.FSEID     // Mandatory
	CreatePDR ie.CreatePDR // Mandatory
	CreateFAR ie.CreateFAR // Mandatory
}

type PFCPSessionEstablishmentResponse struct {
	NodeID ie.NodeID // Mandatory
	Cause  ie.Cause  // Mandatory
}

func (msg PFCPSessionEstablishmentRequest) GetIEs() []ie.InformationElement {
	return []ie.InformationElement{msg.NodeID, msg.CPFSEID, msg.CreatePDR, msg.CreateFAR}
}

func (msg PFCPSessionEstablishmentResponse) GetIEs() []ie.InformationElement {
	return []ie.InformationElement{msg.NodeID, msg.Cause}
}

func (msg PFCPSessionEstablishmentRequest) GetMessageType() MessageType {
	return PFCPSessionEstablishmentRequestMessageType
}

func (msg PFCPSessionEstablishmentResponse) GetMessageType() MessageType {
	return PFCPSessionEstablishmentResponseMessageType
}

func (msg PFCPSessionEstablishmentRequest) GetMessageTypeString() string {
	return "PFCP Session Establishment Request"
}

func (msg PFCPSessionEstablishmentResponse) GetMessageTypeString() string {
	return "PFCP Session Establishment Response"
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

func DeserializePFCPSessionEstablishmentResponse(data []byte) (PFCPMessage, error) {
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

	return PFCPSessionEstablishmentResponse{
		NodeID: nodeID,
		Cause:  cause,
	}, err
}
