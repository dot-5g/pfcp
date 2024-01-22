package client_test

import (
	"testing"
	"time"

	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/messages"
)

type MockUDPSender struct {
	SendFunc func(msg []byte) error
}

func (m *MockUDPSender) Send(msg []byte) error {
	return m.SendFunc(msg)
}

func TestGivenPfcpWhenSendHeartbeatRequestThenNoError(t *testing.T) {
	mockSender := &MockUDPSender{
		SendFunc: func(msg []byte) error {
			return nil
		},
	}
	pfcpClient := client.New("127.0.0.1:8805")
	pfcpClient.Udp = mockSender
	recoveryTimeStamp, err := ie.NewRecoveryTimeStamp(time.Now())

	if err != nil {
		t.Fatalf("Error creating Recovery TimeStamp: %v", err)
	}

	sequenceNumber := uint32(21)
	heartbeatRequestMsg := messages.HeartbeatRequest{
		RecoveryTimeStamp: recoveryTimeStamp,
	}

	err = pfcpClient.SendHeartbeatRequest(heartbeatRequestMsg, sequenceNumber)

	if err != nil {
		t.Errorf("SendHeartbeatRequest failed: %v", err)
	}
}
