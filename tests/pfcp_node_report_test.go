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
	pfcpNodeReportRequestMu                     sync.Mutex
	pfcpNodeReportRequesthandlerCalled          bool
	pfcpNodeReportRequestReceivedSequenceNumber uint32
	pfcpNodeReportRequestReceivedNodeID         ie.NodeID
	pfcpNodeReportRequestReceivedNodeReportType ie.NodeReportType
)

var (
	pfcpNodeReportResponseMu                     sync.Mutex
	pfcpNodeReportResponsehandlerCalled          bool
	pfcpNodeReportResponseReceivedSequenceNumber uint32
	pfcpNodeReportResponseReceivedNodeID         ie.NodeID
	pfcpNodeReportResponseReceivedCause          ie.Cause
)

func HandlePFCPNodeReportRequest(client *client.PFCP, sequenceNumber uint32, msg messages.PFCPNodeReportRequest) {
	pfcpNodeReportRequestMu.Lock()
	defer pfcpNodeReportRequestMu.Unlock()
	pfcpNodeReportRequesthandlerCalled = true
	pfcpNodeReportRequestReceivedSequenceNumber = sequenceNumber
	pfcpNodeReportRequestReceivedNodeID = msg.NodeID
	pfcpNodeReportRequestReceivedNodeReportType = msg.NodeReportType
}

func HandlePFCPNodeReportResponse(client *client.PFCP, sequenceNumber uint32, msg messages.PFCPNodeReportResponse) {
	pfcpNodeReportResponseMu.Lock()
	defer pfcpNodeReportResponseMu.Unlock()
	pfcpNodeReportResponsehandlerCalled = true
	pfcpNodeReportResponseReceivedSequenceNumber = sequenceNumber
	pfcpNodeReportResponseReceivedNodeID = msg.NodeID
	pfcpNodeReportResponseReceivedCause = msg.Cause
}

func TestPFCPNodeReport(t *testing.T) {
	t.Run("TestPFCPNodeReportRequest", PFCPNodeReportRequest)
	t.Run("TestPFCPNodeReportResponse", PFCPNodeReportResponse)
}

func PFCPNodeReportRequest(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPNodeReportRequest(HandlePFCPNodeReportRequest)

	go func() {
		err := pfcpServer.Run()
		if err != nil {
			t.Errorf("Expected no error to be returned")
		}
	}()

	defer pfcpServer.Close()

	time.Sleep(time.Second)
	pfcpClient := client.New("127.0.0.1:8805")
	nodeID, err := ie.NewNodeID("12.23.34.45")

	if err != nil {
		t.Fatalf("Error creating NodeID: %v", err)
	}

	gpqr := false
	ckdr := false
	uprr := true
	upfr := false

	nodeReportType, err := ie.NewNodeReportType(gpqr, ckdr, uprr, upfr)

	if err != nil {
		t.Fatalf("Error creating NodeReportType: %v", err)
	}

	sequenceNumber := uint32(32)
	PFCPNodeReportRequestMsg := messages.PFCPNodeReportRequest{
		NodeID:         nodeID,
		NodeReportType: nodeReportType,
	}

	err = pfcpClient.SendPFCPNodeReportRequest(PFCPNodeReportRequestMsg, sequenceNumber)
	if err != nil {
		t.Fatalf("Error sending PFCP Node Report Request: %v", err)
	}

	time.Sleep(time.Second)

	pfcpNodeReportRequestMu.Lock()
	if !pfcpNodeReportRequesthandlerCalled {
		t.Fatalf("PFCP Node Report Request handler was not called")
	}

	if pfcpNodeReportRequestReceivedSequenceNumber != sequenceNumber {
		t.Errorf("PFCP Node Report Request handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sequenceNumber, pfcpNodeReportRequestReceivedSequenceNumber)
	}

	if pfcpNodeReportRequestReceivedNodeID.Type != nodeID.Type {
		t.Errorf("PFCP Node Report Request handler was called with wrong node ID type.\n- Sent node ID type: %v\n- Received node ID type %v\n", nodeID.Type, pfcpNodeReportRequestReceivedNodeID.Type)
	}

	if len(pfcpNodeReportRequestReceivedNodeID.Value) != len(nodeID.Value) {
		t.Errorf("PFCP Node Report Request handler was called with wrong node ID value length.\n- Sent node ID value length: %v\n- Received node ID value length %v\n", len(nodeID.Value), len(pfcpNodeReportRequestReceivedNodeID.Value))
	}

	for i := range nodeID.Value {
		if pfcpNodeReportRequestReceivedNodeID.Value[i] != nodeID.Value[i] {
			t.Errorf("PFCP Node Report Request handler was called with wrong node ID value.\n- Sent node ID value: %v\n- Received node ID value %v\n", nodeID.Value, pfcpNodeReportRequestReceivedNodeID.Value)
		}
	}

	if pfcpNodeReportRequestReceivedNodeReportType.GPQR != nodeReportType.GPQR {
		t.Errorf("PFCP Node Report Request handler was called with wrong node report type GPQR.\n- Sent node report type GPQR: %v\n- Received node report type GPQR %v\n", nodeReportType.GPQR, pfcpNodeReportRequestReceivedNodeReportType.GPQR)
	}

	if pfcpNodeReportRequestReceivedNodeReportType.CKDR != nodeReportType.CKDR {
		t.Errorf("PFCP Node Report Request handler was called with wrong node report type CKDR.\n- Sent node report type CKDR: %v\n- Received node report type CKDR %v\n", nodeReportType.CKDR, pfcpNodeReportRequestReceivedNodeReportType.CKDR)
	}

	if pfcpNodeReportRequestReceivedNodeReportType.UPRR != nodeReportType.UPRR {
		t.Errorf("PFCP Node Report Request handler was called with wrong node report type UPRR.\n- Sent node report type UPRR: %v\n- Received node report type UPRR %v\n", nodeReportType.UPRR, pfcpNodeReportRequestReceivedNodeReportType.UPRR)
	}

	if pfcpNodeReportRequestReceivedNodeReportType.UPFR != nodeReportType.UPFR {
		t.Errorf("PFCP Node Report Request handler was called with wrong node report type UPFR.\n- Sent node report type UPFR: %v\n- Received node report type UPFR %v\n", nodeReportType.UPFR, pfcpNodeReportRequestReceivedNodeReportType.UPFR)
	}

	pfcpNodeReportRequestMu.Unlock()
}

func PFCPNodeReportResponse(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPNodeReportResponse(HandlePFCPNodeReportResponse)

	go func() {
		err := pfcpServer.Run()
		if err != nil {
			t.Errorf("Expected no error to be returned")
		}
	}()

	defer pfcpServer.Close()

	time.Sleep(time.Second)
	pfcpClient := client.New("127.0.0.1:8805")
	nodeID, err := ie.NewNodeID("3.4.5.6")

	if err != nil {
		t.Fatalf("Error creating NodeID: %v", err)
	}

	sequenceNumber := uint32(32)
	cause, err := ie.NewCause(ie.RequestAccepted)

	if err != nil {
		t.Fatalf("Error creating cause IE: %v", err)
	}

	PFCPNodeReportResponseMsg := messages.PFCPNodeReportResponse{
		NodeID: nodeID,
		Cause:  cause,
	}

	err = pfcpClient.SendPFCPNodeReportResponse(PFCPNodeReportResponseMsg, sequenceNumber)
	if err != nil {
		t.Fatalf("Error sending PFCP Node Report Response: %v", err)
	}

	time.Sleep(time.Second)

	pfcpNodeReportResponseMu.Lock()
	if !pfcpNodeReportResponsehandlerCalled {
		t.Fatalf("PFCP Node Report Response handler was not called")
	}

	if pfcpNodeReportResponseReceivedSequenceNumber != sequenceNumber {
		t.Errorf("PFCP Node Report Response handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sequenceNumber, pfcpNodeReportResponseReceivedSequenceNumber)
	}

	if pfcpNodeReportResponseReceivedNodeID.Type != nodeID.Type {
		t.Errorf("PFCP Node Report Response handler was called with wrong node ID type.\n- Sent node ID type: %v\n- Received node ID type %v\n", nodeID.Type, pfcpNodeReportResponseReceivedNodeID.Type)
	}

	if len(pfcpNodeReportResponseReceivedNodeID.Value) != len(nodeID.Value) {
		t.Errorf("PFCP Node Report Response handler was called with wrong node ID value length.\n- Sent node ID value length: %v\n- Received node ID value length %v\n", len(nodeID.Value), len(pfcpNodeReportResponseReceivedNodeID.Value))
	}

	for i := range nodeID.Value {
		if pfcpNodeReportResponseReceivedNodeID.Value[i] != nodeID.Value[i] {
			t.Errorf("PFCP Node Report Response handler was called with wrong node ID value.\n- Sent node ID value: %v\n- Received node ID value %v\n", nodeID.Value, pfcpNodeReportResponseReceivedNodeID.Value)
		}
	}

	if pfcpNodeReportResponseReceivedCause.Value != cause.Value {
		t.Errorf("PFCP Node Report Response handler was called with wrong cause value.\n- Sent cause value: %v\n- Received cause value %v\n", cause.Value, pfcpNodeReportResponseReceivedCause.Value)
	}
}
