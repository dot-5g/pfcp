package tests

import (
	"sync"
	"testing"
	"time"

	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/ie"
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

func HandlePFCPAssociationUpdateRequest(sequenceNumber uint32, nodeID ie.NodeID) {
	pfcpAssociationUpdateRequestMu.Lock()
	defer pfcpAssociationUpdateRequestMu.Unlock()
	pfcpAssociationUpdateRequesthandlerCalled = true
	pfcpAssociationUpdateRequestReceivedSequenceNumber = sequenceNumber
	pfcpAssociationUpdateRequestReceivedNodeID = nodeID
}

func HandlePFCPAssociationUpdateResponse(sequenceNumber uint32, nodeID ie.NodeID, cause ie.Cause) {
	pfcpAssociationUpdateResponseMu.Lock()
	defer pfcpAssociationUpdateResponseMu.Unlock()
	pfcpAssociationUpdateResponsehandlerCalled = true
	pfcpAssociationUpdateResponseReceivedSequenceNumber = sequenceNumber
	pfcpAssociationUpdateResponseReceivedNodeID = nodeID
	pfcpAssociationUpdateResponseReceivedCause = cause
}

func TestPFCPAssociationUpdate(t *testing.T) {
	t.Run("TestPFCPAssociationUpdateRequest", PFCPAssociationUpdateRequest)
	t.Run("TestPFCPAssociationUpdateResponse", PFCPAssociationUpdateResponse)
}

func PFCPAssociationUpdateRequest(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPAssociationUpdateRequest(HandlePFCPAssociationUpdateRequest)

	pfcpServer.Run()

	defer pfcpServer.Close()

	time.Sleep(time.Second)
	pfcpClient := client.New("127.0.0.1:8805")
	nodeID := ie.NewNodeID(ie.IPv4, "12.23.34.45")
	sequenceNumber := uint32(32)
	pfcpClient.SendPFCPAssociationUpdateRequest(nodeID, sequenceNumber)

	time.Sleep(time.Second)

	pfcpAssociationUpdateRequestMu.Lock()
	if !pfcpAssociationUpdateRequesthandlerCalled {
		t.Errorf("PFCP Association Update Request handler was not called")
	}

	if pfcpAssociationUpdateRequestReceivedSequenceNumber != sequenceNumber {
		t.Errorf("PFCP Association Update Request handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sequenceNumber, pfcpAssociationUpdateRequestReceivedSequenceNumber)
	}

	if pfcpAssociationUpdateRequestReceivedNodeID.Length != nodeID.Length {
		t.Errorf("PFCP Association Update Request handler was called with wrong node ID length.\n- Sent node ID length: %v\n- Received node ID length %v\n", nodeID.Length, pfcpAssociationUpdateRequestReceivedNodeID.Length)
	}

	if pfcpAssociationUpdateRequestReceivedNodeID.NodeIDType != nodeID.NodeIDType {
		t.Errorf("PFCP Association Update Request handler was called with wrong node ID type.\n- Sent node ID type: %v\n- Received node ID type %v\n", nodeID.NodeIDType, pfcpAssociationUpdateRequestReceivedNodeID.NodeIDType)
	}

	if len(pfcpAssociationUpdateRequestReceivedNodeID.NodeIDValue) != len(nodeID.NodeIDValue) {
		t.Errorf("PFCP Association Update Request handler was called with wrong node ID value length.\n- Sent node ID value length: %v\n- Received node ID value length %v\n", len(nodeID.NodeIDValue), len(pfcpAssociationUpdateRequestReceivedNodeID.NodeIDValue))
	}

	for i := range nodeID.NodeIDValue {
		if pfcpAssociationUpdateRequestReceivedNodeID.NodeIDValue[i] != nodeID.NodeIDValue[i] {
			t.Errorf("PFCP Association Update Request handler was called with wrong node ID value.\n- Sent node ID value: %v\n- Received node ID value %v\n", nodeID.NodeIDValue, pfcpAssociationUpdateRequestReceivedNodeID.NodeIDValue)
		}
	}

	pfcpAssociationUpdateRequestMu.Unlock()
}

func PFCPAssociationUpdateResponse(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPAssociationUpdateResponse(HandlePFCPAssociationUpdateResponse)

	pfcpServer.Run()

	defer pfcpServer.Close()

	time.Sleep(time.Second)
	pfcpClient := client.New("127.0.0.1:8805")
	nodeID := ie.NewNodeID(ie.IPv4, "3.4.5.6")

	sequenceNumber := uint32(32)
	cause := ie.NewCause(2)
	pfcpClient.SendPFCPAssociationUpdateResponse(nodeID, cause, sequenceNumber)

	time.Sleep(time.Second)

	pfcpAssociationUpdateResponseMu.Lock()
	if !pfcpAssociationUpdateResponsehandlerCalled {
		t.Errorf("PFCP Association Update Response handler was not called")
	}

	if pfcpAssociationUpdateResponseReceivedSequenceNumber != sequenceNumber {
		t.Errorf("PFCP Association Update Response handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sequenceNumber, pfcpAssociationUpdateResponseReceivedSequenceNumber)
	}

	if pfcpAssociationUpdateResponseReceivedNodeID.Length != nodeID.Length {
		t.Errorf("PFCP Association Update Response handler was called with wrong node ID length.\n- Sent node ID length: %v\n- Received node ID length %v\n", nodeID.Length, pfcpAssociationUpdateResponseReceivedNodeID.Length)
	}

	if pfcpAssociationUpdateResponseReceivedNodeID.NodeIDType != nodeID.NodeIDType {
		t.Errorf("PFCP Association Update Response handler was called with wrong node ID type.\n- Sent node ID type: %v\n- Received node ID type %v\n", nodeID.NodeIDType, pfcpAssociationUpdateResponseReceivedNodeID.NodeIDType)
	}

	if len(pfcpAssociationUpdateResponseReceivedNodeID.NodeIDValue) != len(nodeID.NodeIDValue) {
		t.Errorf("PFCP Association Update Response handler was called with wrong node ID value length.\n- Sent node ID value length: %v\n- Received node ID value length %v\n", len(nodeID.NodeIDValue), len(pfcpAssociationUpdateResponseReceivedNodeID.NodeIDValue))
	}

	for i := range nodeID.NodeIDValue {
		if pfcpAssociationUpdateResponseReceivedNodeID.NodeIDValue[i] != nodeID.NodeIDValue[i] {
			t.Errorf("PFCP Association Update Response handler was called with wrong node ID value.\n- Sent node ID value: %v\n- Received node ID value %v\n", nodeID.NodeIDValue, pfcpAssociationUpdateResponseReceivedNodeID.NodeIDValue)
		}
	}

	if pfcpAssociationUpdateResponseReceivedCause.Length != cause.Length {
		t.Errorf("PFCP Association Update Response handler was called with wrong cause length.\n- Sent cause length: %v\n- Received cause length %v\n", cause.Length, pfcpAssociationUpdateResponseReceivedCause.Length)
	}

	if pfcpAssociationUpdateResponseReceivedCause.Type != cause.Type {
		t.Errorf("PFCP Association Update Response handler was called with wrong cause type.\n- Sent cause type: %v\n- Received cause type %v\n", cause.Type, pfcpAssociationUpdateResponseReceivedCause.Type)
	}

	if pfcpAssociationUpdateResponseReceivedCause.Value != cause.Value {
		t.Errorf("PFCP Association Update Response handler was called with wrong cause value.\n- Sent cause value: %v\n- Received cause value %v\n", cause.Value, pfcpAssociationUpdateResponseReceivedCause.Value)
	}

}
