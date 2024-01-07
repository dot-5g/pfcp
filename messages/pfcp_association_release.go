package messages

import "github.com/dot-5g/pfcp/ie"

type PFCPAssociationReleaseRequest struct {
	NodeID ie.NodeID // Mandatory
}

type PFCPAssociationReleaseResponse struct {
	NodeID ie.NodeID // Mandatory
	Cause  ie.Cause  // Mandatory
}

func (msg PFCPAssociationReleaseRequest) GetIEs() []ie.InformationElement {
	return []ie.InformationElement{msg.NodeID}
}

func (msg PFCPAssociationReleaseResponse) GetIEs() []ie.InformationElement {
	return []ie.InformationElement{msg.NodeID, msg.Cause}
}

func (msg PFCPAssociationReleaseRequest) GetMessageType() MessageType {
	return PFCPAssociationReleaseRequestMessageType
}

func (msg PFCPAssociationReleaseResponse) GetMessageType() MessageType {
	return PFCPAssociationReleaseResponseMessageType
}

func (msg PFCPAssociationReleaseRequest) GetMessageTypeString() string {
	return "PFCP Association Release Request"
}

func (msg PFCPAssociationReleaseResponse) GetMessageTypeString() string {
	return "PFCP Association Release Response"
}

func DeserializePFCPAssociationReleaseRequest(data []byte) (PFCPMessage, error) {
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

func DeserializePFCPAssociationReleaseResponse(data []byte) (PFCPMessage, error) {
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
