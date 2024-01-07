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
	pfcpAssociationSetupRequestMu                         sync.Mutex
	pfcpAssociationSetupRequesthandlerCalled              bool
	pfcpAssociationSetupRequestReceivedSequenceNumber     uint32
	pfcpAssociationSetupRequestReceivedRecoveryTimeStamp  ie.RecoveryTimeStamp
	pfcpAssociationSetupRequestReceivedNodeID             ie.NodeID
	pfcpAssociationSetupRequestReceivedUPFunctionFeatures ie.UPFunctionFeatures
)

var (
	pfcpAssociationSetupResponseMu                        sync.Mutex
	pfcpAssociationSetupResponsehandlerCalled             bool
	pfcpAssociationSetupResponseReceivedSequenceNumber    uint32
	pfcpAssociationSetupResponseReceivedRecoveryTimeStamp ie.RecoveryTimeStamp
	pfcpAssociationSetupResponseReceivedNodeID            ie.NodeID
	pfcpAssociationSetupResponseReceivedCause             ie.Cause
)

func HandlePFCPAssociationSetupRequest(sequenceNumber uint32, msg messages.PFCPAssociationSetupRequest) {
	pfcpAssociationSetupRequestMu.Lock()
	defer pfcpAssociationSetupRequestMu.Unlock()
	pfcpAssociationSetupRequesthandlerCalled = true
	pfcpAssociationSetupRequestReceivedSequenceNumber = sequenceNumber
	pfcpAssociationSetupRequestReceivedRecoveryTimeStamp = msg.RecoveryTimeStamp
	pfcpAssociationSetupRequestReceivedNodeID = msg.NodeID
	pfcpAssociationSetupRequestReceivedUPFunctionFeatures = msg.UPFunctionFeatures
}

func HandlePFCPAssociationSetupResponse(sequenceNumber uint32, msg messages.PFCPAssociationSetupResponse) {
	pfcpAssociationSetupResponseMu.Lock()
	defer pfcpAssociationSetupResponseMu.Unlock()
	pfcpAssociationSetupResponsehandlerCalled = true
	pfcpAssociationSetupResponseReceivedSequenceNumber = sequenceNumber
	pfcpAssociationSetupResponseReceivedRecoveryTimeStamp = msg.RecoveryTimeStamp
	pfcpAssociationSetupResponseReceivedNodeID = msg.NodeID
	pfcpAssociationSetupResponseReceivedCause = msg.Cause
}

func TestPFCPAssociationSetup(t *testing.T) {
	t.Run("TestPFCPAssociationSetupRequest", PFCPAssociationSetupRequest)
	t.Run("TestPFCPAssociationSetupResponse", PFCPAssociationSetupResponse)
}

func PFCPAssociationSetupRequest(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPAssociationSetupRequest(HandlePFCPAssociationSetupRequest)

	go pfcpServer.Run()

	defer pfcpServer.Close()

	time.Sleep(time.Second)
	pfcpClient := client.New("127.0.0.1:8805")
	nodeID, err := ie.NewNodeID("12.23.34.45")

	if err != nil {
		t.Fatalf("Error creating node ID IE: %v", err)
	}

	recoveryTimeStamp, err := ie.NewRecoveryTimeStamp(time.Now())

	if err != nil {
		t.Fatalf("Error creating recovery timestamp IE: %v", err)
	}

	sequenceNumber := uint32(32)
	features := [](ie.UPFeature){
		ie.BUCP,
		ie.TRACE,
	}
	upFeatures, err := ie.NewUPFunctionFeatures(features)

	if err != nil {
		t.Fatalf("Error creating UP function features IE: %v", err)
	}

	PFCPAssociationSetupRequestMsg := messages.PFCPAssociationSetupRequest{
		NodeID:             nodeID,
		RecoveryTimeStamp:  recoveryTimeStamp,
		UPFunctionFeatures: upFeatures,
	}

	pfcpClient.SendPFCPAssociationSetupRequest(PFCPAssociationSetupRequestMsg, sequenceNumber)

	time.Sleep(time.Second)

	pfcpAssociationSetupRequestMu.Lock()
	if !pfcpAssociationSetupRequesthandlerCalled {
		t.Fatalf("PFCP Association Setup Request handler was not called")
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

	if pfcpAssociationSetupRequestReceivedUPFunctionFeatures.Length != upFeatures.Length {
		t.Errorf("PFCP Association Setup Request handler was called with wrong UP function features length.\n- Sent UP function features length: %v\n- Received UP function features length %v\n", upFeatures.Length, pfcpAssociationSetupRequestReceivedUPFunctionFeatures.Length)
	}

	receivedFeatures := pfcpAssociationSetupRequestReceivedUPFunctionFeatures.GetFeatures()

	if len(receivedFeatures) != len(features) {
		t.Errorf("PFCP Association Setup Request handler was called with wrong UP function features.\n- Sent UP function features: %v\n- Received UP function features %v\n", features, receivedFeatures)
	}

	for i := range features {
		if receivedFeatures[i] != features[i] {
			t.Errorf("PFCP Association Setup Request handler was called with wrong UP function features.\n- Sent UP function features: %v\n- Received UP function features %v\n", features, receivedFeatures)
		}
	}

	pfcpAssociationSetupRequestMu.Unlock()
}

func PFCPAssociationSetupResponse(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPAssociationSetupResponse(HandlePFCPAssociationSetupResponse)

	go pfcpServer.Run()

	defer pfcpServer.Close()

	time.Sleep(time.Second)

	pfcpClient := client.New("127.0.0.1:8805")
	nodeID, err := ie.NewNodeID("1.2.3.4")

	if err != nil {
		t.Fatalf("Error creating node ID IE: %v", err)
	}

	cause, err := ie.NewCause(ie.RequestAccepted)

	if err != nil {
		t.Fatalf("Error creating cause IE: %v", err)
	}

	recoveryTimeStamp, err := ie.NewRecoveryTimeStamp(time.Now())

	if err != nil {
		t.Fatalf("Error creating recovery timestamp IE: %v", err)
	}

	sequenceNumber := uint32(32)
	PFCPAssociationSetupResponseMsg := messages.PFCPAssociationSetupResponse{
		NodeID:            nodeID,
		Cause:             cause,
		RecoveryTimeStamp: recoveryTimeStamp,
	}

	pfcpClient.SendPFCPAssociationSetupResponse(PFCPAssociationSetupResponseMsg, sequenceNumber)

	time.Sleep(time.Second)

	pfcpAssociationSetupResponseMu.Lock()

	if !pfcpAssociationSetupResponsehandlerCalled {
		t.Fatalf("PFCP Association Setup Response handler was not called")
	}

	if pfcpAssociationSetupResponseReceivedSequenceNumber != sequenceNumber {
		t.Errorf("PFCP Association Setup Response handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sequenceNumber, pfcpAssociationSetupResponseReceivedSequenceNumber)
	}

	if pfcpAssociationSetupResponseReceivedRecoveryTimeStamp != recoveryTimeStamp {
		t.Errorf("PFCP Association Setup Response handler was called with wrong recovery timestamp.\n- Sent recovery timestamp: %v\n- Received recovery timestamp %v\n", recoveryTimeStamp, pfcpAssociationSetupResponseReceivedRecoveryTimeStamp)
	}

	if pfcpAssociationSetupResponseReceivedNodeID.Length != nodeID.Length {
		t.Errorf("PFCP Association Setup Response handler was called with wrong node ID length.\n- Sent node ID length: %v\n- Received node ID length %v\n", nodeID.Length, pfcpAssociationSetupResponseReceivedNodeID.Length)
	}

	if pfcpAssociationSetupResponseReceivedNodeID.NodeIDType != nodeID.NodeIDType {
		t.Errorf("PFCP Association Setup Response handler was called with wrong node ID type.\n- Sent node ID type: %v\n- Received node ID type %v\n", nodeID.NodeIDType, pfcpAssociationSetupResponseReceivedNodeID.NodeIDType)
	}

	if len(pfcpAssociationSetupResponseReceivedNodeID.NodeIDValue) != len(nodeID.NodeIDValue) {
		t.Errorf("PFCP Association Setup Response handler was called with wrong node ID value length.\n- Sent node ID value length: %v\n- Received node ID value length %v\n", len(nodeID.NodeIDValue), len(pfcpAssociationSetupResponseReceivedNodeID.NodeIDValue))
	}

	for i := range nodeID.NodeIDValue {
		if pfcpAssociationSetupResponseReceivedNodeID.NodeIDValue[i] != nodeID.NodeIDValue[i] {
			t.Errorf("PFCP Association Setup Response handler was called with wrong node ID value.\n- Sent node ID value: %v\n- Received node ID value %v\n", nodeID.NodeIDValue, pfcpAssociationSetupResponseReceivedNodeID.NodeIDValue)
		}
	}

	if pfcpAssociationSetupResponseReceivedCause.IEType != cause.IEType {
		t.Errorf("PFCP Association Setup Response handler was called with wrong cause type.\n- Sent cause type: %v\n- Received cause type %v\n", cause.IEType, pfcpAssociationSetupResponseReceivedCause.IEType)
	}

	if pfcpAssociationSetupResponseReceivedCause.Length != cause.Length {
		t.Errorf("PFCP Association Setup Response handler was called with wrong cause length.\n- Sent cause length: %v\n- Received cause length %v\n", cause.Length, pfcpAssociationSetupResponseReceivedCause.Length)
	}

	if pfcpAssociationSetupResponseReceivedCause.Value != cause.Value {
		t.Errorf("PFCP Association Setup Response handler was called with wrong cause value.\n- Sent cause value: %v\n- Received cause value %v\n", cause.Value, pfcpAssociationSetupResponseReceivedCause.Value)
	}
	pfcpAssociationSetupResponseMu.Unlock()

}
