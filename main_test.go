package main_test

import (
	"sync"
	"testing"
	"time"

	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/server"
)

var (
	heartbeatRequestMu                        sync.Mutex
	heartbeatRequesthandlerCalled             bool
	heartbeatRequestreceivedRecoveryTimestamp ie.RecoveryTimeStamp
	heartbeatRequestReceivedSequenceNumber    uint32
)

var (
	heartbeatResponseMu                        sync.Mutex
	heartbeatResponsehandlerCalled             bool
	heartbeatResponsereceivedRecoveryTimestamp ie.RecoveryTimeStamp
	heartbeatResponseReceivedSequenceNumber    uint32
)

var (
	pfcpAssociationSetupRequestMu                        sync.Mutex
	pfcpAssociationSetupRequesthandlerCalled             bool
	pfcpAssociationSetupRequestReceivedSequenceNumber    uint32
	pfcpAssociationSetupRequestReceivedRecoveryTimeStamp ie.RecoveryTimeStamp
	pfcpAssociationSetupRequestReceivedNodeID            ie.NodeID
)

func HandleHeartbeatRequest(sequenceNumber uint32, recoveryTimeStamp ie.RecoveryTimeStamp) {
	heartbeatRequestMu.Lock()
	defer heartbeatRequestMu.Unlock()
	heartbeatRequesthandlerCalled = true
	heartbeatRequestreceivedRecoveryTimestamp = recoveryTimeStamp
	heartbeatRequestReceivedSequenceNumber = sequenceNumber
}
func HandleHeartbeatResponse(sequenceNumber uint32, recoveryTimeStamp ie.RecoveryTimeStamp) {
	heartbeatResponseMu.Lock()
	defer heartbeatResponseMu.Unlock()
	heartbeatResponsehandlerCalled = true
	heartbeatResponsereceivedRecoveryTimestamp = recoveryTimeStamp
	heartbeatResponseReceivedSequenceNumber = sequenceNumber
}
func HandlePFCPAssociationSetupRequest(sequenceNumber uint32, nodeID ie.NodeID, recoveryTimeStamp ie.RecoveryTimeStamp) {
	pfcpAssociationSetupRequestMu.Lock()
	defer pfcpAssociationSetupRequestMu.Unlock()
	pfcpAssociationSetupRequesthandlerCalled = true
	pfcpAssociationSetupRequestReceivedSequenceNumber = sequenceNumber
	pfcpAssociationSetupRequestReceivedRecoveryTimeStamp = recoveryTimeStamp
	pfcpAssociationSetupRequestReceivedNodeID = nodeID
}

func TestServer(t *testing.T) {
	// t.Run("TestHeartbeatRequest", HeartbeatRequest)
	// t.Run("TestHeartbeatResponse", HeartbeatResponse)
	t.Run("TestPFCPAssociationSetupRequest", PFCPAssociationSetupRequest)
}

func HeartbeatRequest(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.HeartbeatRequest(HandleHeartbeatRequest)
	sentSequenceNumber := uint32(32)
	recoveryTimeStamp := ie.NewRecoveryTimeStamp(time.Now())

	go pfcpServer.Run()

	time.Sleep(time.Second)

	pfcpClient := client.New("127.0.0.1:8805")
	sentRecoveryTimeStamp, err := pfcpClient.SendHeartbeatRequest(recoveryTimeStamp, sentSequenceNumber)
	if err != nil {
		t.Fatalf("Failed to send Heartbeat request: %v", err)
	}

	time.Sleep(time.Second)

	heartbeatRequestMu.Lock()
	if !heartbeatRequesthandlerCalled {
		t.Errorf("Heartbeat request handler was not called")
	}
	if heartbeatRequestreceivedRecoveryTimestamp != sentRecoveryTimeStamp {
		t.Errorf("Heartbeat request handler was called with wrong timestamp.\n- Sent timestamp: %v\n- Received timestamp %v\n", sentRecoveryTimeStamp, heartbeatRequestreceivedRecoveryTimestamp)
	}
	if heartbeatRequestReceivedSequenceNumber != sentSequenceNumber {
		t.Errorf("Heartbeat request handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sentSequenceNumber, heartbeatRequestReceivedSequenceNumber)
	}
	heartbeatRequestMu.Unlock()
	pfcpServer.Close()
}

func HeartbeatResponse(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.HeartbeatResponse(HandleHeartbeatResponse)
	sentSequenceNumber := uint32(971)
	recoveryTimeStamp := ie.NewRecoveryTimeStamp(time.Now())

	go pfcpServer.Run()

	time.Sleep(time.Second)

	pfcpClient := client.New("127.0.0.1:8805")
	sentRecoveryTimeStamp, err := pfcpClient.SendHeartbeatResponse(recoveryTimeStamp, sentSequenceNumber)
	if err != nil {
		t.Fatalf("Failed to send Heartbeat response: %v", err)
	}

	time.Sleep(time.Second)

	heartbeatResponseMu.Lock()
	if !heartbeatResponsehandlerCalled {
		t.Errorf("Heartbeat response handler was not called")
	}
	if heartbeatResponsereceivedRecoveryTimestamp != sentRecoveryTimeStamp {
		t.Errorf("Heartbeat response handler was called with wrong timestamp.\n- Sent timestamp: %v\n- Received timestamp %v\n", sentRecoveryTimeStamp, heartbeatResponsereceivedRecoveryTimestamp)
	}
	if heartbeatResponseReceivedSequenceNumber != sentSequenceNumber {
		t.Errorf("Heartbeat response handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sentSequenceNumber, heartbeatResponseReceivedSequenceNumber)
	}

	heartbeatResponseMu.Unlock()
	pfcpServer.Close()

}

func PFCPAssociationSetupRequest(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPAssociationSetupRequest(HandlePFCPAssociationSetupRequest)

	go pfcpServer.Run()

	time.Sleep(time.Second)
	pfcpClient := client.New("127.0.0.1:8805")
	nodeID := ie.NewNodeID(ie.IPv4, "12.23.34.45")
	recoveryTimeStamp := ie.NewRecoveryTimeStamp(time.Now())
	sequenceNumber := uint32(32)
	pfcpClient.SendPFCPAssociationSetupRequest(nodeID, recoveryTimeStamp, sequenceNumber)

	time.Sleep(time.Second)

	pfcpAssociationSetupRequestMu.Lock()
	if !pfcpAssociationSetupRequesthandlerCalled {
		t.Errorf("PFCP Association Setup Request handler was not called")
	}

	if pfcpAssociationSetupRequestReceivedSequenceNumber != sequenceNumber {
		t.Errorf("PFCP Association Setup Request handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sequenceNumber, pfcpAssociationSetupRequestReceivedSequenceNumber)
	}

	if pfcpAssociationSetupRequestReceivedRecoveryTimeStamp != recoveryTimeStamp {
		t.Errorf("PFCP Association Setup Request handler was called with wrong recovery timestamp.\n- Sent recovery timestamp: %v\n- Received recovery timestamp %v\n", recoveryTimeStamp, pfcpAssociationSetupRequestReceivedRecoveryTimeStamp)
	}

	if pfcpAssociationSetupRequestReceivedNodeID.Length != nodeID.Length {
		t.Errorf("PFCP Association Setup Request handler was called with wrong node ID length.\n- Sent node ID length: %v\n- Received node ID length %v\n", nodeID.Length, pfcpAssociationSetupRequestReceivedNodeID.Length)
	}

	if pfcpAssociationSetupRequestReceivedNodeID.NodeIDType != nodeID.NodeIDType {
		t.Errorf("PFCP Association Setup Request handler was called with wrong node ID type.\n- Sent node ID type: %v\n- Received node ID type %v\n", nodeID.NodeIDType, pfcpAssociationSetupRequestReceivedNodeID.NodeIDType)
	}

	if len(pfcpAssociationSetupRequestReceivedNodeID.NodeIDValue) != len(nodeID.NodeIDValue) {
		t.Errorf("PFCP Association Setup Request handler was called with wrong node ID value length.\n- Sent node ID value length: %v\n- Received node ID value length %v\n", len(nodeID.NodeIDValue), len(pfcpAssociationSetupRequestReceivedNodeID.NodeIDValue))
	}

	for i := range nodeID.NodeIDValue {
		if pfcpAssociationSetupRequestReceivedNodeID.NodeIDValue[i] != nodeID.NodeIDValue[i] {
			t.Errorf("PFCP Association Setup Request handler was called with wrong node ID value.\n- Sent node ID value: %v\n- Received node ID value %v\n", nodeID.NodeIDValue, pfcpAssociationSetupRequestReceivedNodeID.NodeIDValue)
		}
	}

	pfcpAssociationSetupRequestMu.Unlock()
}
