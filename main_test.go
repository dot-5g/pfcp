package main_test

import (
	"sync"
	"testing"
	"time"

	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/server"
)

var (
	heartbeatRequestMu                        sync.Mutex
	heartbeatRequesthandlerCalled             bool
	heartbeatRequestreceivedRecoveryTimestamp messages.RecoveryTimeStamp
)

var (
	heartbeatResponseMu                        sync.Mutex
	heartbeatResponsehandlerCalled             bool
	heartbeatResponsereceivedRecoveryTimestamp messages.RecoveryTimeStamp
)

func HandleHeartbeatRequest(h *messages.HeartbeatRequest) {
	heartbeatRequestMu.Lock()
	defer heartbeatRequestMu.Unlock()
	heartbeatRequesthandlerCalled = true
	heartbeatRequestreceivedRecoveryTimestamp = h.RecoveryTimeStamp
}

func HandleHeartbeatResponse(h *messages.HeartbeatResponse) {
	heartbeatResponseMu.Lock()
	defer heartbeatResponseMu.Unlock()
	heartbeatResponsehandlerCalled = true
	heartbeatResponsereceivedRecoveryTimestamp = h.RecoveryTimeStamp
}

func TestServer(t *testing.T) {
	t.Run("TestHeartbeatRequest", HeartbeatRequest)
	t.Run("TestHeartbeatResponse", HeartbeatResponse)
}

func HeartbeatRequest(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.HeartbeatRequest(HandleHeartbeatRequest)

	go pfcpServer.Run()

	time.Sleep(time.Second)

	pfcpClient := client.New("127.0.0.1:8805")
	sentRecoveryTimeStamp, err := pfcpClient.SendHeartbeatRequest(time.Now())
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
	heartbeatRequestMu.Unlock()
	pfcpServer.Close()
}

func HeartbeatResponse(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.HeartbeatResponse(HandleHeartbeatResponse)

	go pfcpServer.Run()

	time.Sleep(time.Second)

	pfcpClient := client.New("127.0.0.1:8805")
	sentRecoveryTimeStamp, err := pfcpClient.SendHeartbeatResponse(time.Now())
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
	heartbeatResponseMu.Unlock()
	pfcpServer.Close()

}
