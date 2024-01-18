package messages

import "github.com/dot-5g/pfcp/ie"

type PFCPSessionEstablishmentRequest struct {
	NodeID  ie.NodeID // Mandatory
	CPFSEID ie.FSEID  // Mandatory
	PDR     ie.PDR    // Mandatory
	FAR     ie.FAR    // Mandatory
}

type PFCPSessionEstablishmentResponse struct {
	NodeID ie.NodeID // Mandatory
	Cause  ie.Cause  // Mandatory
}

func (msg PFCPSessionEstablishmentRequest) GetIEs() []ie.InformationElement {
	return []ie.InformationElement{msg.NodeID, msg.CPFSEID, msg.PDR, msg.FAR}
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

func DeserializePFCPSessionEstablishmentRequest(data []byte) (PFCPSessionEstablishmentRequest, error) {
	ies, err := ie.DeserializeInformationElements(data)
	var nodeID ie.NodeID
	var controlPlaneFSEID ie.FSEID
	var pdr ie.PDR
	var far ie.FAR

	for _, elem := range ies {
		if nodeIDIE, ok := elem.(ie.NodeID); ok {
			nodeID = nodeIDIE
			continue
		}
		if controlPlaneFSEIDIE, ok := elem.(ie.FSEID); ok {
			controlPlaneFSEID = controlPlaneFSEIDIE
			continue
		}
		if pdrIE, ok := elem.(ie.PDR); ok {
			pdr = pdrIE
			continue
		}
		if farIE, ok := elem.(ie.FAR); ok {
			far = farIE
			continue
		}

	}

	return PFCPSessionEstablishmentRequest{
		NodeID:  nodeID,
		CPFSEID: controlPlaneFSEID,
		PDR:     pdr,
		FAR:     far,
	}, err
}

func DeserializePFCPSessionEstablishmentResponse(data []byte) (PFCPSessionEstablishmentResponse, error) {
	ies, err := ie.DeserializeInformationElements(data)
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
