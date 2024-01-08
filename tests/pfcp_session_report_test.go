package tests

import (
	"sync"
	"testing"
	"time"

	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/server"
)

var (
	pfcpSessionReportRequestMu                     sync.Mutex
	pfcpSessionReportRequesthandlerCalled          bool
	pfcpSessionReportRequestReceivedSequenceNumber uint32
	pfcpSessionReportRequestReceivedSEID           uint64
	pfcpSessionReportRequestReceivedReportType     ie.ReportType
)

var (
	pfcpSessionReportResponseMu                     sync.Mutex
	pfcpSessionReportResponsehandlerCalled          bool
	pfcpSessionReportResponseReceivedSequenceNumber uint32
	pfcpSessionReportResponseReceivedSEID           uint64
	pfcpSessionReportResponseReceivedCause          ie.Cause
)

func HandlePFCPSessionReportRequest(sequenceNumber uint32, seid uint64, msg messages.PFCPSessionReportRequest) {
	pfcpSessionReportRequestMu.Lock()
	defer pfcpSessionReportRequestMu.Unlock()
	pfcpSessionReportRequesthandlerCalled = true
	pfcpSessionReportRequestReceivedSequenceNumber = sequenceNumber
	pfcpSessionReportRequestReceivedSEID = seid
	pfcpSessionReportRequestReceivedReportType = msg.ReportType
}

func HandlePFCPSessionReportResponse(sequenceNumber uint32, seid uint64, msg messages.PFCPSessionReportResponse) {
	pfcpSessionReportResponseMu.Lock()
	defer pfcpSessionReportResponseMu.Unlock()
	pfcpSessionReportResponsehandlerCalled = true
	pfcpSessionReportResponseReceivedSequenceNumber = sequenceNumber
	pfcpSessionReportResponseReceivedSEID = seid
	pfcpSessionReportResponseReceivedCause = msg.Cause
}

func TestPFCPSessionReport(t *testing.T) {
	t.Run("TestPFCPSessionReportRequest", PFCPSessionReportRequest)
	t.Run("TestPFCPSessionReportResponse", PFCPSessionReportResponse)
}

func PFCPSessionReportRequest(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPSessionReportRequest(HandlePFCPSessionReportRequest)
	go pfcpServer.Run()
	defer pfcpServer.Close()

	time.Sleep(time.Second)
	pfcpClient := client.New("127.0.0.1:8805")

	reportType, err := ie.NewReportType([]ie.Report{ie.UISR, ie.SESR})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	PFCPSessionReportRequestMsg := messages.PFCPSessionReportRequest{
		ReportType: reportType,
	}
	seid := uint64(12345)
	sequenceNumber := uint32(1)
	err = pfcpClient.SendPFCPSessionReportRequest(&PFCPSessionReportRequestMsg, seid, sequenceNumber)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	time.Sleep(time.Second)

	pfcpSessionReportRequestMu.Lock()

	if !pfcpSessionReportRequesthandlerCalled {
		t.Fatalf("PFCP Session Report Request handler was not called")
	}

	if pfcpSessionReportRequestReceivedSequenceNumber != sequenceNumber {
		t.Errorf("Expected sequence number %d, got %d", sequenceNumber, pfcpSessionReportRequestReceivedSequenceNumber)
	}

	if pfcpSessionReportRequestReceivedSEID != seid {
		t.Errorf("Expected SEID %d, got %d", seid, pfcpSessionReportRequestReceivedSEID)
	}

	if pfcpSessionReportRequestReceivedReportType.Header.Type != ie.ReportTypeIEType {
		t.Errorf("Expected IE type %d, got %d", ie.ReportTypeIEType, pfcpSessionReportRequestReceivedReportType.Header.Type)
	}

	if pfcpSessionReportRequestReceivedReportType.Header.Length != 1 {
		t.Errorf("Expected length 1, got %d", pfcpSessionReportRequestReceivedReportType.Header.Length)
	}

	if len(pfcpSessionReportRequestReceivedReportType.Reports) != 2 {
		t.Errorf("Expected 2 reports, got %d", len(pfcpSessionReportRequestReceivedReportType.Reports))
	}

	if pfcpSessionReportRequestReceivedReportType.Reports[0] != ie.UISR {
		t.Errorf("Expected report %d, got %d", ie.UISR, pfcpSessionReportRequestReceivedReportType.Reports[0])
	}

	if pfcpSessionReportRequestReceivedReportType.Reports[1] != ie.SESR {
		t.Errorf("Expected report %d, got %d", ie.SESR, pfcpSessionReportRequestReceivedReportType.Reports[1])
	}

}

func PFCPSessionReportResponse(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPSessionReportResponse(HandlePFCPSessionReportResponse)
	go pfcpServer.Run()
	defer pfcpServer.Close()

	time.Sleep(time.Second)
	pfcpClient := client.New("127.0.0.1:8805")

	cause, err := ie.NewCause(ie.RequestAccepted)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	PFCPSessionReportResponseMsg := messages.PFCPSessionReportResponse{
		Cause: cause,
	}
	seid := uint64(12345)
	sequenceNumber := uint32(1)
	err = pfcpClient.SendPFCPSessionReportResponse(&PFCPSessionReportResponseMsg, seid, sequenceNumber)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	time.Sleep(time.Second)

	pfcpSessionReportResponseMu.Lock()

	if !pfcpSessionReportResponsehandlerCalled {
		t.Fatalf("PFCP Session Report Response handler was not called")
	}

	if pfcpSessionReportResponseReceivedSequenceNumber != sequenceNumber {
		t.Errorf("Expected sequence number %d, got %d", sequenceNumber, pfcpSessionReportResponseReceivedSequenceNumber)
	}

	if pfcpSessionReportResponseReceivedSEID != seid {
		t.Errorf("Expected SEID %d, got %d", seid, pfcpSessionReportResponseReceivedSEID)
	}

	if pfcpSessionReportResponseReceivedCause.Header.Type != ie.CauseIEType {
		t.Errorf("Expected IE type %d, got %d", ie.CauseIEType, pfcpSessionReportResponseReceivedCause.Header.Type)
	}

	if pfcpSessionReportResponseReceivedCause.Header.Length != 1 {
		t.Errorf("Expected length 1, got %d", pfcpSessionReportResponseReceivedCause.Header.Length)
	}

	if pfcpSessionReportResponseReceivedCause.Value != ie.RequestAccepted {
		t.Errorf("Expected cause value %d, got %d", ie.RequestAccepted, pfcpSessionReportResponseReceivedCause.Value)
	}

}
