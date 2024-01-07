package messages

import "github.com/dot-5g/pfcp/ie"

type PFCPSessionReportRequest struct {
	ReportType ie.ReportType // Mandatory
}

type PFCPSessionReportResponse struct {
	Cause ie.Cause // Mandatory
}

func DeserializePFCPSessionReportRequest(data []byte) (PFCPMessage, error) {
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

func DeserializePFCPSessionReportResponse(data []byte) (PFCPMessage, error) {
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
