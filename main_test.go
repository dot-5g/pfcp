package main_test

import (
	"sync"
	"testing"
	"time"

	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/server"
)

var (
	mu            sync.Mutex
	handlerCalled bool
)

func HandleHeartbeatRequest(h server.HeartbeatRequest) {
	mu.Lock()
	defer mu.Unlock()
	handlerCalled = true
}

func TestGivenHandleHeartbeatRequestWhenRunThenHeartbeatRequestHandled(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.HeartbeatRequest(HandleHeartbeatRequest)

	go pfcpServer.Run()

	time.Sleep(time.Second)

	// Setup PFCP client
	pfcpClient := client.New("127.0.0.1:8805")
	err := pfcpClient.SendHeartbeatRequest()
	if err != nil {
		t.Fatalf("Failed to send Heartbeat request: %v", err)
	}

	time.Sleep(time.Second)

	// Check if handler was called
	mu.Lock()
	if !handlerCalled {
		t.Errorf("Heartbeat request handler was not called")
	}
	mu.Unlock()

}
