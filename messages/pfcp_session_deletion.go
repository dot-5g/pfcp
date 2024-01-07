package messages

import "github.com/dot-5g/pfcp/ie"

type PFCPSessionDeletionRequest struct{}

type PFCPSessionDeletionResponse struct {
	Cause ie.Cause // Mandatory
}

func DeserializePFCPSessionDeletionRequest(data []byte) (PFCPMessage, error) {
	return PFCPSessionDeletionRequest{}, nil
}

func DeserializePFCPSessionDeletionResponse(data []byte) (PFCPMessage, error) {
	ies, err := ie.ParseInformationElements(data)
	var cause ie.Cause
	for _, elem := range ies {
		if causeIE, ok := elem.(ie.Cause); ok {
			cause = causeIE
			continue
		}
	}

	return PFCPSessionDeletionResponse{
		Cause: cause,
	}, err
}
