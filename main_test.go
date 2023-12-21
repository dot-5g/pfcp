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
	mu                        sync.Mutex
	handlerCalled             bool
	receivedRecoveryTimestamp messages.RecoveryTimeStamp
)

func HandleHeartbeatRequest(h *messages.HeartbeatRequest) {
	mu.Lock()
	defer mu.Unlock()
	handlerCalled = true
	receivedRecoveryTimestamp = h.RecoveryTimeStamp
}

func TestGivenHandleHeartbeatRequestWhenRunThenHeartbeatRequestHandled(t *testing.T) {
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

	mu.Lock()
	if !handlerCalled {
		t.Errorf("Heartbeat request handler was not called")
	}
	if receivedRecoveryTimestamp != sentRecoveryTimeStamp {
		t.Errorf("Heartbeat request handler was called with wrong timestamp.\n- Sent timestamp: %v\n- Received timestamp %v\n", sentRecoveryTimeStamp, receivedRecoveryTimestamp)
	}
	mu.Unlock()

}
