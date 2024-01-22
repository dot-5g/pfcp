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
	pfcpAssociationReleaseRequestMu                     sync.Mutex
	pfcpAssociationReleaseRequesthandlerCalled          bool
	pfcpAssociationReleaseRequestReceivedSequenceNumber uint32
	pfcpAssociationReleaseRequestReceivedNodeID         ie.NodeID
)

var (
	pfcpAssociationReleaseResponseMu                     sync.Mutex
	pfcpAssociationReleaseResponsehandlerCalled          bool
	pfcpAssociationReleaseResponseReceivedSequenceNumber uint32
	pfcpAssociationReleaseResponseReceivedNodeID         ie.NodeID
	pfcpAssociationReleaseResponseReceivedCause          ie.Cause
)

func HandlePFCPAssociationReleaseRequest(client *client.PFCP, sequenceNumber uint32, msg messages.PFCPAssociationReleaseRequest) {
	pfcpAssociationReleaseRequestMu.Lock()
	defer pfcpAssociationReleaseRequestMu.Unlock()
	pfcpAssociationReleaseRequesthandlerCalled = true
	pfcpAssociationReleaseRequestReceivedSequenceNumber = sequenceNumber
	pfcpAssociationReleaseRequestReceivedNodeID = msg.NodeID
}

func HandlePFCPAssociationReleaseResponse(client *client.PFCP, sequenceNumber uint32, msg messages.PFCPAssociationReleaseResponse) {
	pfcpAssociationReleaseResponseMu.Lock()
	defer pfcpAssociationReleaseResponseMu.Unlock()
	pfcpAssociationReleaseResponsehandlerCalled = true
	pfcpAssociationReleaseResponseReceivedSequenceNumber = sequenceNumber
	pfcpAssociationReleaseResponseReceivedNodeID = msg.NodeID
	pfcpAssociationReleaseResponseReceivedCause = msg.Cause
}

func TestPFCPAssociationRelease(t *testing.T) {
	t.Run("TestPFCPAssociationReleaseRequest", PFCPAssociationReleaseRequest)
	t.Run("TestPFCPAssociationReleaseResponse", PFCPAssociationReleaseResponse)
}

func PFCPAssociationReleaseRequest(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPAssociationReleaseRequest(HandlePFCPAssociationReleaseRequest)

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
		t.Fatalf("Error creating node ID IE: %v", err)
	}

	sequenceNumber := uint32(32)
	PFCPAssociationReleaseRequestMsg := messages.PFCPAssociationReleaseRequest{
		NodeID: nodeID,
	}

	err = pfcpClient.SendPFCPAssociationReleaseRequest(PFCPAssociationReleaseRequestMsg, sequenceNumber)
	if err != nil {
		t.Fatalf("Failed to send PFCP Association Release Request: %v", err)
	}

	time.Sleep(time.Second)

	pfcpAssociationReleaseRequestMu.Lock()
	if !pfcpAssociationReleaseRequesthandlerCalled {
		t.Fatalf("PFCP Association Release Request handler was not called")
	}

	if pfcpAssociationReleaseRequestReceivedSequenceNumber != sequenceNumber {
		t.Errorf("PFCP Association Release Request handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sequenceNumber, pfcpAssociationReleaseRequestReceivedSequenceNumber)
	}

	if pfcpAssociationReleaseRequestReceivedNodeID.Type != nodeID.Type {
		t.Errorf("PFCP Association Release Request handler was called with wrong node ID type.\n- Sent node ID type: %v\n- Received node ID type %v\n", nodeID.Type, pfcpAssociationReleaseRequestReceivedNodeID.Type)
	}

	if len(pfcpAssociationReleaseRequestReceivedNodeID.Value) != len(nodeID.Value) {
		t.Errorf("PFCP Association Release Request handler was called with wrong node ID value length.\n- Sent node ID value length: %v\n- Received node ID value length %v\n", len(nodeID.Value), len(pfcpAssociationReleaseRequestReceivedNodeID.Value))
	}

	for i := range nodeID.Value {
		if pfcpAssociationReleaseRequestReceivedNodeID.Value[i] != nodeID.Value[i] {
			t.Errorf("PFCP Association Release Request handler was called with wrong node ID value.\n- Sent node ID value: %v\n- Received node ID value %v\n", nodeID.Value, pfcpAssociationReleaseRequestReceivedNodeID.Value)
		}
	}

	pfcpAssociationReleaseRequestMu.Unlock()
}

func PFCPAssociationReleaseResponse(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPAssociationReleaseResponse(HandlePFCPAssociationReleaseResponse)

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
		t.Fatalf("Error creating node ID IE: %v", err)
	}

	sequenceNumber := uint32(32)
	cause, err := ie.NewCause(ie.RequestAccepted)

	if err != nil {
		t.Fatalf("Error creating cause IE: %v", err)
	}

	PFCPAssociationReleaseResponseMsg := messages.PFCPAssociationReleaseResponse{
		NodeID: nodeID,
		Cause:  cause,
	}

	err = pfcpClient.SendPFCPAssociationReleaseResponse(PFCPAssociationReleaseResponseMsg, sequenceNumber)
	if err != nil {
		t.Fatalf("Failed to send PFCP Association Release Response: %v", err)
	}

	time.Sleep(time.Second)

	pfcpAssociationReleaseResponseMu.Lock()
	if !pfcpAssociationReleaseResponsehandlerCalled {
		t.Fatalf("PFCP Association Release Response handler was not called")
	}

	if pfcpAssociationReleaseResponseReceivedSequenceNumber != sequenceNumber {
		t.Errorf("PFCP Association Release Response handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sequenceNumber, pfcpAssociationReleaseResponseReceivedSequenceNumber)
	}

	if pfcpAssociationReleaseResponseReceivedNodeID.Type != nodeID.Type {
		t.Errorf("PFCP Association Release Response handler was called with wrong node ID type.\n- Sent node ID type: %v\n- Received node ID type %v\n", nodeID.Type, pfcpAssociationReleaseResponseReceivedNodeID.Type)
	}

	if len(pfcpAssociationReleaseResponseReceivedNodeID.Value) != len(nodeID.Value) {
		t.Errorf("PFCP Association Release Response handler was called with wrong node ID value length.\n- Sent node ID value length: %v\n- Received node ID value length %v\n", len(nodeID.Value), len(pfcpAssociationReleaseResponseReceivedNodeID.Value))
	}

	for i := range nodeID.Value {
		if pfcpAssociationReleaseResponseReceivedNodeID.Value[i] != nodeID.Value[i] {
			t.Errorf("PFCP Association Release Response handler was called with wrong node ID value.\n- Sent node ID value: %v\n- Received node ID value %v\n", nodeID.Value, pfcpAssociationReleaseResponseReceivedNodeID.Value)
		}
	}

	if pfcpAssociationReleaseResponseReceivedCause.Value != cause.Value {
		t.Errorf("PFCP Association Release Response handler was called with wrong cause value.\n- Sent cause value: %v\n- Received cause value %v\n", cause.Value, pfcpAssociationReleaseResponseReceivedCause.Value)
	}
}
