package client_test

import (
	"testing"
	"time"

	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/ie"
)

type MockUdpSender struct {
	SendFunc func(msg []byte) error
}

func (m *MockUdpSender) Send(msg []byte) error {
	return m.SendFunc(msg)
}

func TestGivenPfcpWhenSendHeartbeatRequestThenNoError(t *testing.T) {
	mockSender := &MockUdpSender{
		SendFunc: func(msg []byte) error {
			return nil
		},
	}
	pfcpClient := client.New("127.0.0.1:8805")
	pfcpClient.Udp = mockSender
	recoveryTimeStamp := ie.NewRecoveryTimeStamp(time.Now())
	sequenceNumber := uint32(21)

	_, err := pfcpClient.SendHeartbeatRequest(recoveryTimeStamp, sequenceNumber)

	if err != nil {
		t.Errorf("SendHeartbeatRequest failed: %v", err)
	}

}
