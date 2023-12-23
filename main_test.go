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

func TestServer(t *testing.T) {
	t.Run("TestHeartbeatRequest", HeartbeatRequest)
	t.Run("TestHeartbeatResponse", HeartbeatResponse)
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
