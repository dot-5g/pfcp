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
	heartbeatRequestMu                        sync.Mutex
	heartbeatRequesthandlerCalled             bool
	heartbeatRequestreceivedRecoveryTimestamp ie.RecoveryTimeStamp
	heartbeatRequestReceivedSequenceNumber    uint32
	heartbeatRequestReceivedPFCPClient        *client.PFCP
)

var (
	heartbeatRequestWithSourceIPMu                        sync.Mutex
	heartbeatRequestWithSourceIPhandlerCalled             bool
	heartbeatRequestWithSourceIPreceivedRecoveryTimestamp ie.RecoveryTimeStamp
	heartbeatRequestWithSourceIPreceivedSourceIPAddress   ie.SourceIPAddress
	heartbeatRequestWithSourceIPReceivedSequenceNumber    uint32
)

var (
	heartbeatResponseMu                        sync.Mutex
	heartbeatResponsehandlerCalled             bool
	heartbeatResponsereceivedRecoveryTimestamp ie.RecoveryTimeStamp
	heartbeatResponseReceivedSequenceNumber    uint32
)

func HandleHeartbeatRequest(pfcpClient *client.PFCP, sequenceNumber uint32, msg messages.HeartbeatRequest) {
	heartbeatRequestMu.Lock()
	defer heartbeatRequestMu.Unlock()
	heartbeatRequesthandlerCalled = true
	heartbeatRequestreceivedRecoveryTimestamp = msg.RecoveryTimeStamp
	heartbeatRequestReceivedSequenceNumber = sequenceNumber
	heartbeatRequestReceivedPFCPClient = pfcpClient
}

func HandleHeartbeatRequestWithSourceIP(pfcpClient *client.PFCP, sequenceNumber uint32, msg messages.HeartbeatRequest) {
	heartbeatRequestWithSourceIPMu.Lock()
	defer heartbeatRequestWithSourceIPMu.Unlock()
	heartbeatRequestWithSourceIPhandlerCalled = true
	heartbeatRequestWithSourceIPreceivedRecoveryTimestamp = msg.RecoveryTimeStamp
	heartbeatRequestWithSourceIPreceivedSourceIPAddress = msg.SourceIPAddress
	heartbeatRequestWithSourceIPReceivedSequenceNumber = sequenceNumber
}

func HandleHeartbeatResponse(pfcpClient *client.PFCP, sequenceNumber uint32, msg messages.HeartbeatResponse) {
	heartbeatResponseMu.Lock()
	defer heartbeatResponseMu.Unlock()
	heartbeatResponsehandlerCalled = true
	heartbeatResponsereceivedRecoveryTimestamp = msg.RecoveryTimeStamp
	heartbeatResponseReceivedSequenceNumber = sequenceNumber
}

func TestHeartbeat(t *testing.T) {
	t.Run("TestHeartbeatRequest", HeartbeatRequest)
	t.Run("TestHeartbeatRequestWithSourceIPAddress", HeartbeatRequestWithSourceIPAddress)
	t.Run("TestHeartbeatResponse", HeartbeatResponse)
}

func HeartbeatRequest(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.HeartbeatRequest(HandleHeartbeatRequest)
	sentSequenceNumber := uint32(32)
	recoveryTimeStamp, err := ie.NewRecoveryTimeStamp(time.Now())

	if err != nil {
		t.Fatalf("Error creating Recovery Time Stamp IE: %v", err)
	}

	heartbeatRequestMsg := messages.HeartbeatRequest{
		RecoveryTimeStamp: recoveryTimeStamp,
	}

	go pfcpServer.Run()

	defer pfcpServer.Close()

	time.Sleep(time.Second)

	pfcpClient := client.New("127.0.0.1:8805")
	err = pfcpClient.SendHeartbeatRequest(heartbeatRequestMsg, sentSequenceNumber)
	if err != nil {
		t.Fatalf("Failed to send Heartbeat request: %v", err)
	}

	time.Sleep(time.Second)

	heartbeatRequestMu.Lock()
	if !heartbeatRequesthandlerCalled {
		t.Fatalf("Heartbeat request handler was not called")
	}
	if heartbeatRequestreceivedRecoveryTimestamp != recoveryTimeStamp {
		t.Errorf("Heartbeat request handler was called with wrong timestamp.\n- Sent timestamp: %v\n- Received timestamp %v\n", recoveryTimeStamp, heartbeatRequestreceivedRecoveryTimestamp)
	}
	if heartbeatRequestReceivedSequenceNumber != sentSequenceNumber {
		t.Errorf("Heartbeat request handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sentSequenceNumber, heartbeatRequestReceivedSequenceNumber)
	}

	if heartbeatRequestReceivedPFCPClient == nil {
		t.Errorf("Heartbeat request handler was called with wrong PFCP client.\n- Sent PFCP client: %v\n- Received PFCP client %v\n", pfcpClient, heartbeatRequestReceivedPFCPClient)
	}

	heartbeatRequestMu.Unlock()
}

func HeartbeatRequestWithSourceIPAddress(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.HeartbeatRequest(HandleHeartbeatRequestWithSourceIP)
	sentSequenceNumber := uint32(32)
	recoveryTimeStamp, err := ie.NewRecoveryTimeStamp(time.Now())

	if err != nil {
		t.Fatalf("Error creating Recovery Time Stamp IE: %v", err)
	}

	sourceIPAddress, err := ie.NewSourceIPAddress("2.3.2.3/24", "")

	if err != nil {
		t.Fatalf("Error creating Source IP Address IE: %v", err)
	}

	heartbeatRequestMsg := messages.HeartbeatRequest{
		RecoveryTimeStamp: recoveryTimeStamp,
		SourceIPAddress:   sourceIPAddress,
	}

	go pfcpServer.Run()

	defer pfcpServer.Close()

	time.Sleep(time.Second)

	pfcpClient := client.New("127.0.0.1:8805")
	err = pfcpClient.SendHeartbeatRequest(heartbeatRequestMsg, sentSequenceNumber)
	if err != nil {
		t.Fatalf("Failed to send Heartbeat request: %v", err)
	}

	time.Sleep(time.Second)

	heartbeatRequestWithSourceIPMu.Lock()
	if !heartbeatRequestWithSourceIPhandlerCalled {
		t.Fatalf("Heartbeat request handler was not called")
	}
	if heartbeatRequestWithSourceIPreceivedRecoveryTimestamp != recoveryTimeStamp {
		t.Errorf("Heartbeat request handler was called with wrong timestamp.\n- Sent timestamp: %v\n- Received timestamp %v\n", recoveryTimeStamp, heartbeatRequestWithSourceIPreceivedRecoveryTimestamp)
	}

	if len(heartbeatRequestWithSourceIPreceivedSourceIPAddress.IPv4Address) != len(sourceIPAddress.IPv4Address) {
		t.Errorf("Heartbeat request handler was called with wrong source IP address.\n- Sent source IP address: %v\n- Received source IP address %v\n", sourceIPAddress.IPv4Address, heartbeatRequestWithSourceIPreceivedSourceIPAddress.IPv4Address)
	}

	for i := range sourceIPAddress.IPv4Address {
		if heartbeatRequestWithSourceIPreceivedSourceIPAddress.IPv4Address[i] != sourceIPAddress.IPv4Address[i] {
			t.Errorf("Heartbeat request handler was called with wrong source IP address.\n- Sent source IP address: %v\n- Received source IP address %v\n", sourceIPAddress.IPv4Address, heartbeatRequestWithSourceIPreceivedSourceIPAddress.IPv4Address)
		}
	}

	if heartbeatRequestWithSourceIPReceivedSequenceNumber != sentSequenceNumber {
		t.Errorf("Heartbeat request handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sentSequenceNumber, heartbeatRequestWithSourceIPReceivedSequenceNumber)
	}
	heartbeatRequestWithSourceIPMu.Unlock()
}

func HeartbeatResponse(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.HeartbeatResponse(HandleHeartbeatResponse)
	sentSequenceNumber := uint32(971)
	recoveryTimeStamp, err := ie.NewRecoveryTimeStamp(time.Now())

	if err != nil {
		t.Fatalf("Error creating Recovery Time Stamp IE: %v", err)
	}

	heartbeatResponseMsg := messages.HeartbeatResponse{
		RecoveryTimeStamp: recoveryTimeStamp,
	}

	go pfcpServer.Run()

	defer pfcpServer.Close()

	time.Sleep(time.Second)

	pfcpClient := client.New("127.0.0.1:8805")
	err = pfcpClient.SendHeartbeatResponse(heartbeatResponseMsg, sentSequenceNumber)
	if err != nil {
		t.Fatalf("Failed to send Heartbeat response: %v", err)
	}

	time.Sleep(time.Second)

	heartbeatResponseMu.Lock()
	if !heartbeatResponsehandlerCalled {
		t.Fatalf("Heartbeat response handler was not called")
	}
	if heartbeatResponsereceivedRecoveryTimestamp != recoveryTimeStamp {
		t.Errorf("Heartbeat response handler was called with wrong timestamp.\n- Sent timestamp: %v\n- Received timestamp %v\n", recoveryTimeStamp, heartbeatResponsereceivedRecoveryTimestamp)
	}
	if heartbeatResponseReceivedSequenceNumber != sentSequenceNumber {
		t.Errorf("Heartbeat response handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sentSequenceNumber, heartbeatResponseReceivedSequenceNumber)
	}

	heartbeatResponseMu.Unlock()
}
