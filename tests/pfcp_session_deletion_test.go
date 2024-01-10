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
	pfcpSessionDeletionRequestMu                     sync.Mutex
	pfcpSessionDeletionRequesthandlerCalled          bool
	pfcpSessionDeletionRequestReceivedSequenceNumber uint32
	pfcpSessionDeletionRequestReceivedSEID           uint64
)

var (
	pfcpSessionDeletionResponseMu                     sync.Mutex
	pfcpSessionDeletionResponsehandlerCalled          bool
	pfcpSessionDeletionResponseReceivedSequenceNumber uint32
	pfcpSessionDeletionResponseReceivedSEID           uint64
	pfcpSessionDeletionResponseReceivedCause          ie.Cause
)

func HandlePFCPSessionDeletionRequest(client *client.Pfcp, sequenceNumber uint32, seid uint64, msg messages.PFCPSessionDeletionRequest) {
	pfcpSessionDeletionRequestMu.Lock()
	defer pfcpSessionDeletionRequestMu.Unlock()
	pfcpSessionDeletionRequesthandlerCalled = true
	pfcpSessionDeletionRequestReceivedSequenceNumber = sequenceNumber
	pfcpSessionDeletionRequestReceivedSEID = seid
}

func HandlePFCPSessionDeletionResponse(client *client.Pfcp, sequenceNumber uint32, seid uint64, msg messages.PFCPSessionDeletionResponse) {
	pfcpSessionDeletionResponseMu.Lock()
	defer pfcpSessionDeletionResponseMu.Unlock()
	pfcpSessionDeletionResponsehandlerCalled = true
	pfcpSessionDeletionResponseReceivedSequenceNumber = sequenceNumber
	pfcpSessionDeletionResponseReceivedSEID = seid
	pfcpSessionDeletionResponseReceivedCause = msg.Cause
}

func TestPFCPSessionDeletion(t *testing.T) {
	t.Run("TestPFCPSessionDeletionRequest", PFCPSessionDeletionRequest)
	// t.Run("TestPFCPSessionDeletionResponse", PFCPSessionDeletionResponse)
}

func PFCPSessionDeletionRequest(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPSessionDeletionRequest(HandlePFCPSessionDeletionRequest)
	go pfcpServer.Run()
	defer pfcpServer.Close()

	time.Sleep(time.Second)
	pfcpClient := client.New("127.0.0.1:8805")

	PFCPSessionDeletionRequestMsg := messages.PFCPSessionDeletionRequest{}
	seid := uint64(1234567890)
	sequenceNumber := uint32(32)

	err := pfcpClient.SendPFCPSessionDeletionRequest(PFCPSessionDeletionRequestMsg, seid, sequenceNumber)
	if err != nil {
		t.Fatalf("Error sending PFCP Session Deletion Request: %v", err)
	}

	time.Sleep(time.Second)

	pfcpSessionDeletionRequestMu.Lock()
	if !pfcpSessionDeletionRequesthandlerCalled {
		t.Fatalf("PFCP Session Deletion Request handler was not called")
	}

	if pfcpSessionDeletionRequestReceivedSequenceNumber != sequenceNumber {
		t.Errorf("PFCP Session Deletion Request handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sequenceNumber, pfcpSessionDeletionRequestReceivedSequenceNumber)
	}

	if pfcpSessionDeletionRequestReceivedSEID != seid {
		t.Errorf("PFCP Session Deletion Request handler was called with wrong SEID.\n- Sent SEID: %v\n- Received SEID %v\n", seid, pfcpSessionDeletionRequestReceivedSEID)
	}

	pfcpSessionDeletionRequestMu.Unlock()
}

func PFCPSessionDeletionResponse(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPSessionDeletionResponse(HandlePFCPSessionDeletionResponse)
	go pfcpServer.Run()
	defer pfcpServer.Close()

	time.Sleep(time.Second)
	pfcpClient := client.New("127.0.0.1:8805")

	PFCPSessionDeletionResponseMsg := messages.PFCPSessionDeletionResponse{}
	seid := uint64(1234567890)
	sequenceNumber := uint32(31232)

	err := pfcpClient.SendPFCPSessionDeletionResponse(PFCPSessionDeletionResponseMsg, seid, sequenceNumber)
	if err != nil {
		t.Fatalf("Error sending PFCP Session Deletion Response: %v", err)
	}

	time.Sleep(time.Second)

	pfcpSessionDeletionResponseMu.Lock()

	if !pfcpSessionDeletionResponsehandlerCalled {
		t.Fatalf("PFCP Session Deletion Response handler was not called")
	}

	if pfcpSessionDeletionResponseReceivedSequenceNumber != sequenceNumber {
		t.Errorf("PFCP Session Deletion Response handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sequenceNumber, pfcpSessionDeletionResponseReceivedSequenceNumber)
	}

	if pfcpSessionDeletionResponseReceivedSEID != seid {
		t.Errorf("PFCP Session Deletion Response handler was called with wrong SEID.\n- Sent SEID: %v\n- Received SEID %v\n", seid, pfcpSessionDeletionResponseReceivedSEID)
	}

	if pfcpSessionDeletionResponseReceivedCause != PFCPSessionDeletionResponseMsg.Cause {
		t.Errorf("PFCP Session Deletion Response handler was called with wrong cause.\n- Sent cause: %v\n- Received cause %v\n", PFCPSessionDeletionResponseMsg.Cause, pfcpSessionDeletionResponseReceivedCause)
	}

}
