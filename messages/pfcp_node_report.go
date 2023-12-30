package messages

import (
	"github.com/dot-5g/pfcp/ie"
)

type PFCPNodeReportRequest struct {
	NodeID         ie.NodeID         // Mandatory
	NodeReportType ie.NodeReportType // Mandatory
}

type PFCPNodeReportResponse struct {
	NodeID ie.NodeID // Mandatory
	Cause  ie.Cause  // Mandatory
}

func ParsePFCPNodeReportRequest(data []byte) (PFCPNodeReportRequest, error) {
	ies, err := ie.ParseInformationElements(data)
	var nodeID ie.NodeID
	var nodeReportType ie.NodeReportType
	for _, elem := range ies {
		if nodeIDIE, ok := elem.(ie.NodeID); ok {
			nodeID = nodeIDIE
			continue
		}
		if nodeReportTypeIE, ok := elem.(ie.NodeReportType); ok {
			nodeReportType = nodeReportTypeIE
			continue
		}
	}

	return PFCPNodeReportRequest{
		NodeID:         nodeID,
		NodeReportType: nodeReportType,
	}, err
}

func ParsePFCPNodeReportResponse(data []byte) (PFCPNodeReportResponse, error) {
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

	return PFCPNodeReportResponse{
		NodeID: nodeID,
		Cause:  cause,
	}, err
}
