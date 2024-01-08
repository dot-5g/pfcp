package messages

import "github.com/dot-5g/pfcp/ie"

type PFCPSessionDeletionRequest struct{}

type PFCPSessionDeletionResponse struct {
	Cause ie.Cause // Mandatory
}

func (msg PFCPSessionDeletionRequest) GetIEs() []ie.InformationElement {
	return []ie.InformationElement{}
}

func (msg PFCPSessionDeletionResponse) GetIEs() []ie.InformationElement {
	return []ie.InformationElement{msg.Cause}
}

func (msg PFCPSessionDeletionRequest) GetMessageType() MessageType {
	return PFCPSessionDeletionRequestMessageType
}

func (msg PFCPSessionDeletionResponse) GetMessageType() MessageType {
	return PFCPSessionDeletionResponseMessageType
}

func (msg PFCPSessionDeletionRequest) GetMessageTypeString() string {
	return "PFCP Session Deletion Request"
}

func (msg PFCPSessionDeletionResponse) GetMessageTypeString() string {
	return "PFCP Session Deletion Response"
}

func DeserializePFCPSessionDeletionRequest(data []byte) (PFCPSessionDeletionRequest, error) {
	return PFCPSessionDeletionRequest{}, nil
}

func DeserializePFCPSessionDeletionResponse(data []byte) (PFCPSessionDeletionResponse, error) {
	ies, err := ie.DeserializeInformationElements(data)
	if err != nil {
		return PFCPSessionDeletionResponse{}, err
	}

	var cause ie.Cause
	for _, elem := range ies {
		if causeIE, ok := elem.(ie.Cause); ok {
			cause = causeIE
			continue
		}
	}

	return PFCPSessionDeletionResponse{
		Cause: cause,
	}, nil
}
