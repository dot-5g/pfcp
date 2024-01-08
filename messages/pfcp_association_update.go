package messages

import "github.com/dot-5g/pfcp/ie"

type PFCPAssociationUpdateRequest struct {
	Header Header
	NodeID ie.NodeID // Mandatory
}

type PFCPAssociationUpdateResponse struct {
	Header Header
	NodeID ie.NodeID // Mandatory
	Cause  ie.Cause  // Mandatory
}

func (msg PFCPAssociationUpdateRequest) GetIEs() []ie.InformationElement {
	return []ie.InformationElement{msg.NodeID}
}

func (msg PFCPAssociationUpdateResponse) GetIEs() []ie.InformationElement {
	return []ie.InformationElement{msg.NodeID, msg.Cause}
}

func (msg PFCPAssociationUpdateRequest) GetMessageType() MessageType {
	return PFCPAssociationUpdateRequestMessageType
}

func (msg PFCPAssociationUpdateResponse) GetMessageType() MessageType {
	return PFCPAssociationUpdateResponseMessageType
}

func (msg PFCPAssociationUpdateRequest) GetMessageTypeString() string {
	return "PFCP Association Update Request"
}

func (msg PFCPAssociationUpdateResponse) GetMessageTypeString() string {
	return "PFCP Association Update Response"
}

func (msg *PFCPAssociationUpdateRequest) SetHeader(h Header) {
	msg.Header = h
}

func (msg *PFCPAssociationUpdateResponse) SetHeader(h Header) {
	msg.Header = h
}

func (msg PFCPAssociationUpdateRequest) GetHeader() Header {
	return msg.Header
}

func (msg PFCPAssociationUpdateResponse) GetHeader() Header {
	return msg.Header
}

func DeserializePFCPAssociationUpdateRequest(data []byte) (PFCPAssociationUpdateRequest, error) {
	ies, err := ie.DeserializeInformationElements(data)
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

func DeserializePFCPAssociationUpdateResponse(data []byte) (PFCPAssociationUpdateResponse, error) {
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

	return PFCPAssociationUpdateResponse{
		NodeID: nodeID,
		Cause:  cause,
	}, err
}
