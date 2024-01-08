package messages

import "github.com/dot-5g/pfcp/ie"

type PFCPSessionReportRequest struct {
	ReportType ie.ReportType // Mandatory
}

type PFCPSessionReportResponse struct {
	Cause ie.Cause // Mandatory
}

func (msg PFCPSessionReportRequest) GetIEs() []ie.InformationElement {
	return []ie.InformationElement{msg.ReportType}
}

func (msg PFCPSessionReportResponse) GetIEs() []ie.InformationElement {
	return []ie.InformationElement{msg.Cause}
}

func (msg PFCPSessionReportRequest) GetMessageType() MessageType {
	return PFCPSessionReportRequestMessageType
}

func (msg PFCPSessionReportResponse) GetMessageType() MessageType {
	return PFCPSessionReportResponseMessageType
}

func (msg PFCPSessionReportRequest) GetMessageTypeString() string {
	return "PFCP Session Report Request"
}

func (msg PFCPSessionReportResponse) GetMessageTypeString() string {
	return "PFCP Session Report Response"
}

func DeserializePFCPSessionReportRequest(data []byte) (PFCPSessionReportRequest, error) {
	ies, err := ie.ParseInformationElements(data)
	var reportType ie.ReportType

	for _, elem := range ies {
		if reportTypeIE, ok := elem.(ie.ReportType); ok {
			reportType = reportTypeIE
			continue
		}

	}

	return PFCPSessionReportRequest{
		ReportType: reportType,
	}, err
}

func DeserializePFCPSessionReportResponse(data []byte) (PFCPSessionReportResponse, error) {
	ies, err := ie.ParseInformationElements(data)
	var cause ie.Cause

	for _, elem := range ies {
		if causeIE, ok := elem.(ie.Cause); ok {
			cause = causeIE
			continue
		}

	}

	return PFCPSessionReportResponse{
		Cause: cause,
	}, err
}
