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
	pfcpAssociationUpdateRequestMu                     sync.Mutex
	pfcpAssociationUpdateRequesthandlerCalled          bool
	pfcpAssociationUpdateRequestReceivedSequenceNumber uint32
	pfcpAssociationUpdateRequestReceivedNodeID         ie.NodeID
)

var (
	pfcpAssociationUpdateResponseMu                     sync.Mutex
	pfcpAssociationUpdateResponsehandlerCalled          bool
	pfcpAssociationUpdateResponseReceivedSequenceNumber uint32
	pfcpAssociationUpdateResponseReceivedNodeID         ie.NodeID
	pfcpAssociationUpdateResponseReceivedCause          ie.Cause
)

func HandlePFCPAssociationUpdateRequest(sequenceNumber uint32, msg messages.PFCPAssociationUpdateRequest) {
	pfcpAssociationUpdateRequestMu.Lock()
	defer pfcpAssociationUpdateRequestMu.Unlock()
	pfcpAssociationUpdateRequesthandlerCalled = true
	pfcpAssociationUpdateRequestReceivedSequenceNumber = sequenceNumber
	pfcpAssociationUpdateRequestReceivedNodeID = msg.NodeID
}

func HandlePFCPAssociationUpdateResponse(sequenceNumber uint32, msg messages.PFCPAssociationUpdateResponse) {
	pfcpAssociationUpdateResponseMu.Lock()
	defer pfcpAssociationUpdateResponseMu.Unlock()
	pfcpAssociationUpdateResponsehandlerCalled = true
	pfcpAssociationUpdateResponseReceivedSequenceNumber = sequenceNumber
	pfcpAssociationUpdateResponseReceivedNodeID = msg.NodeID
	pfcpAssociationUpdateResponseReceivedCause = msg.Cause
}

func TestPFCPAssociationUpdate(t *testing.T) {
	t.Run("TestPFCPAssociationUpdateRequest", PFCPAssociationUpdateRequest)
	t.Run("TestPFCPAssociationUpdateResponse", PFCPAssociationUpdateResponse)
}

func PFCPAssociationUpdateRequest(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPAssociationUpdateRequest(HandlePFCPAssociationUpdateRequest)

	go pfcpServer.Run()

	defer pfcpServer.Close()

	time.Sleep(time.Second)
	pfcpClient := client.New("127.0.0.1:8805")
	nodeID, err := ie.NewNodeID("12.23.34.45")

	if err != nil {
		t.Fatalf("Error creating node ID IE: %v", err)
	}

	sequenceNumber := uint32(32)
	PFCPAssociationUpdateRequestMsg := messages.PFCPAssociationUpdateRequest{
		NodeID: nodeID,
	}

	pfcpClient.SendPFCPAssociationUpdateRequest(PFCPAssociationUpdateRequestMsg, sequenceNumber)

	time.Sleep(time.Second)

	pfcpAssociationUpdateRequestMu.Lock()
	if !pfcpAssociationUpdateRequesthandlerCalled {
		t.Fatalf("PFCP Association Update Request handler was not called")
	}

	if pfcpAssociationUpdateRequestReceivedSequenceNumber != sequenceNumber {
		t.Errorf("PFCP Association Update Request handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sequenceNumber, pfcpAssociationUpdateRequestReceivedSequenceNumber)
	}

	if pfcpAssociationUpdateRequestReceivedNodeID.Header.Length != nodeID.Header.Length {
		t.Errorf("PFCP Association Update Request handler was called with wrong node ID length.\n- Sent node ID length: %v\n- Received node ID length %v\n", nodeID.Header.Length, pfcpAssociationUpdateRequestReceivedNodeID.Header.Length)
	}

	if pfcpAssociationUpdateRequestReceivedNodeID.Type != nodeID.Type {
		t.Errorf("PFCP Association Update Request handler was called with wrong node ID type.\n- Sent node ID type: %v\n- Received node ID type %v\n", nodeID.Type, pfcpAssociationUpdateRequestReceivedNodeID.Type)
	}

	if len(pfcpAssociationUpdateRequestReceivedNodeID.Value) != len(nodeID.Value) {
		t.Errorf("PFCP Association Update Request handler was called with wrong node ID value length.\n- Sent node ID value length: %v\n- Received node ID value length %v\n", len(nodeID.Value), len(pfcpAssociationUpdateRequestReceivedNodeID.Value))
	}

	for i := range nodeID.Value {
		if pfcpAssociationUpdateRequestReceivedNodeID.Value[i] != nodeID.Value[i] {
			t.Errorf("PFCP Association Update Request handler was called with wrong node ID value.\n- Sent node ID value: %v\n- Received node ID value %v\n", nodeID.Value, pfcpAssociationUpdateRequestReceivedNodeID.Value)
		}
	}

	pfcpAssociationUpdateRequestMu.Unlock()
}

func PFCPAssociationUpdateResponse(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPAssociationUpdateResponse(HandlePFCPAssociationUpdateResponse)

	go pfcpServer.Run()

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

	PFCPAssociationUpdateResponseMsg := messages.PFCPAssociationUpdateResponse{
		NodeID: nodeID,
		Cause:  cause,
	}

	pfcpClient.SendPFCPAssociationUpdateResponse(PFCPAssociationUpdateResponseMsg, sequenceNumber)

	time.Sleep(time.Second)

	pfcpAssociationUpdateResponseMu.Lock()
	if !pfcpAssociationUpdateResponsehandlerCalled {
		t.Fatalf("PFCP Association Update Response handler was not called")
	}

	if pfcpAssociationUpdateResponseReceivedSequenceNumber != sequenceNumber {
		t.Errorf("PFCP Association Update Response handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sequenceNumber, pfcpAssociationUpdateResponseReceivedSequenceNumber)
	}

	if pfcpAssociationUpdateResponseReceivedNodeID.Header.Length != nodeID.Header.Length {
		t.Errorf("PFCP Association Update Response handler was called with wrong node ID length.\n- Sent node ID length: %v\n- Received node ID length %v\n", nodeID.Header.Length, pfcpAssociationUpdateResponseReceivedNodeID.Header.Length)
	}

	if pfcpAssociationUpdateResponseReceivedNodeID.Type != nodeID.Type {
		t.Errorf("PFCP Association Update Response handler was called with wrong node ID type.\n- Sent node ID type: %v\n- Received node ID type %v\n", nodeID.Type, pfcpAssociationUpdateResponseReceivedNodeID.Type)
	}

	if len(pfcpAssociationUpdateResponseReceivedNodeID.Value) != len(nodeID.Value) {
		t.Errorf("PFCP Association Update Response handler was called with wrong node ID value length.\n- Sent node ID value length: %v\n- Received node ID value length %v\n", len(nodeID.Value), len(pfcpAssociationUpdateResponseReceivedNodeID.Value))
	}

	for i := range nodeID.Value {
		if pfcpAssociationUpdateResponseReceivedNodeID.Value[i] != nodeID.Value[i] {
			t.Errorf("PFCP Association Update Response handler was called with wrong node ID value.\n- Sent node ID value: %v\n- Received node ID value %v\n", nodeID.Value, pfcpAssociationUpdateResponseReceivedNodeID.Value)
		}
	}

	if pfcpAssociationUpdateResponseReceivedCause.Header.Length != cause.Header.Length {
		t.Errorf("PFCP Association Update Response handler was called with wrong cause length.\n- Sent cause length: %v\n- Received cause length %v\n", cause.Header.Length, pfcpAssociationUpdateResponseReceivedCause.Header.Length)
	}

	if pfcpAssociationUpdateResponseReceivedCause.Header.Type != cause.Header.Type {
		t.Errorf("PFCP Association Update Response handler was called with wrong cause type.\n- Sent cause type: %v\n- Received cause type %v\n", cause.Header.Type, pfcpAssociationUpdateResponseReceivedCause.Header.Type)
	}

	if pfcpAssociationUpdateResponseReceivedCause.Value != cause.Value {
		t.Errorf("PFCP Association Update Response handler was called with wrong cause value.\n- Sent cause value: %v\n- Received cause value %v\n", cause.Value, pfcpAssociationUpdateResponseReceivedCause.Value)
	}

}
