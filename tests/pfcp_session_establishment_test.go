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
	pfcpSessionEstablishmentRequestMu                     sync.Mutex
	pfcpSessionEstablishmentRequesthandlerCalled          bool
	pfcpSessionEstablishmentRequestReceivedSequenceNumber uint32
	pfcpSessionEstablishmentRequestReceivedSEID           uint64
	pfcpSessionEstablishmentRequestReceivedNodeID         ie.NodeID
	pfcpSessionEstablishmentRequestReceivedCPFSEID        ie.FSEID
	pfcpSessionEstablishmentRequestReceivedCreatePDR      ie.CreatePDR
	pfcpSessionEstablishmentRequestReceivedCreateFAR      ie.CreateFAR
)

var (
	pfcpSessionEstablishmentResponseMu                     sync.Mutex
	pfcpSessionEstablishmentResponsehandlerCalled          bool
	pfcpSessionEstablishmentResponseReceivedSequenceNumber uint32
	pfcpSessionEstablishmentResponseReceivedNodeID         ie.NodeID
	pfcpSessionEstablishmentResponseReceivedCause          ie.Cause
)

func HandlePFCPSessionEstablishmentRequest(sequenceNumber uint32, seid uint64, msg messages.PFCPSessionEstablishmentRequest) {
	pfcpSessionEstablishmentRequestMu.Lock()
	defer pfcpSessionEstablishmentRequestMu.Unlock()
	pfcpSessionEstablishmentRequesthandlerCalled = true
	pfcpSessionEstablishmentRequestReceivedSequenceNumber = sequenceNumber
	pfcpSessionEstablishmentRequestReceivedSEID = seid
	pfcpSessionEstablishmentRequestReceivedNodeID = msg.NodeID
	pfcpSessionEstablishmentRequestReceivedCPFSEID = msg.CPFSEID
	pfcpSessionEstablishmentRequestReceivedCreatePDR = msg.CreatePDR
	pfcpSessionEstablishmentRequestReceivedCreateFAR = msg.CreateFAR
}

func HandlePFCPSessionEstablishmentResponse(sequenceNumber uint32, seid uint64, msg messages.PFCPSessionEstablishmentResponse) {
	pfcpSessionEstablishmentResponseMu.Lock()
	defer pfcpSessionEstablishmentResponseMu.Unlock()
	pfcpSessionEstablishmentResponsehandlerCalled = true
	pfcpSessionEstablishmentResponseReceivedSequenceNumber = sequenceNumber
	pfcpSessionEstablishmentResponseReceivedNodeID = msg.NodeID
	pfcpSessionEstablishmentResponseReceivedCause = msg.Cause
}

func TestPFCPSessionEstablishment(t *testing.T) {
	t.Run("TestPFCPSessionEstablishmentRequest", PFCPSessionEstablishmentRequest)
	t.Run("TestPFCPSessionEstablishmentResponse", PFCPSessionEstablishmentResponse)
}

func PFCPSessionEstablishmentRequest(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPSessionEstablishmentRequest(HandlePFCPSessionEstablishmentRequest)
	go pfcpServer.Run()
	defer pfcpServer.Close()

	time.Sleep(time.Second)
	pfcpClient := client.New("127.0.0.1:8805")

	nodeID, err := ie.NewNodeID("12.23.34.45")

	if err != nil {
		t.Fatalf("Error creating Node ID: %v", err)
	}

	seid := uint64(1234567890)

	fseid, err := ie.NewFSEID(seid, "1.2.3.4", "")
	if err != nil {
		t.Fatalf("Error creating FSEID: %v", err)
	}

	pdrID, err := ie.NewPDRID(1)

	if err != nil {
		t.Fatalf("Error creating PDR ID: %v", err)
	}

	precedence, err := ie.NewPrecedence(uint32(1))

	if err != nil {
		t.Fatalf("Error creating Precedence: %v", err)
	}

	sourceInterface, err := ie.NewSourceInterface(3)
	if err != nil {
		t.Fatalf("Error creating SourceInterface: %v", err)
	}

	pdi, err := ie.NewPDI(sourceInterface)
	if err != nil {
		t.Fatalf("Error creating PDI: %v", err)
	}

	createPDR, err := ie.NewCreatePDR(pdrID, precedence, pdi)
	if err != nil {
		t.Fatalf("Error creating CreatePDR: %v", err)
	}

	farID, err := ie.NewFarID(uint32(1))

	if err != nil {
		t.Fatalf("Error creating FarID: %v", err)
	}

	applyAction, err := ie.NewApplyAction(ie.FORW, []ie.ApplyActionExtraFlag{})

	if err != nil {
		t.Fatalf("Error creating ApplyAction: %v", err)
	}

	createFAR, err := ie.NewCreateFAR(farID, applyAction)
	if err != nil {
		t.Fatalf("Error creating CreateFAR: %v", err)
	}

	PFCPSessionEstablishmentRequestMsg := messages.PFCPSessionEstablishmentRequest{
		NodeID:    nodeID,
		CPFSEID:   fseid,
		CreatePDR: createPDR,
		CreateFAR: createFAR,
	}
	sequenceNumber := uint32(32)

	err = pfcpClient.SendPFCPSessionEstablishmentRequest(&PFCPSessionEstablishmentRequestMsg, seid, sequenceNumber)
	if err != nil {
		t.Fatalf("Error sending PFCP Session Establishment Request: %v", err)
	}

	time.Sleep(time.Second)

	pfcpSessionEstablishmentRequestMu.Lock()
	if !pfcpSessionEstablishmentRequesthandlerCalled {
		t.Fatalf("PFCP Session Establishment Request handler was not called")
	}

	if pfcpSessionEstablishmentRequestReceivedSequenceNumber != sequenceNumber {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sequenceNumber, pfcpSessionEstablishmentRequestReceivedSequenceNumber)
	}

	if pfcpSessionEstablishmentRequestReceivedSEID != seid {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong SEID.\n- Sent SEID: %v\n- Received SEID %v\n", seid, pfcpSessionEstablishmentRequestReceivedSEID)
	}

	if pfcpSessionEstablishmentRequestReceivedNodeID.Header.Length != nodeID.Header.Length {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong node ID length.\n- Sent node ID length: %v\n- Received node ID length %v\n", nodeID.Header.Length, pfcpSessionEstablishmentRequestReceivedNodeID.Header.Length)
	}

	if pfcpSessionEstablishmentRequestReceivedNodeID.Type != nodeID.Type {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong node ID type.\n- Sent node ID type: %v\n- Received node ID type %v\n", nodeID.Type, pfcpSessionEstablishmentRequestReceivedNodeID.Type)
	}

	if len(pfcpSessionEstablishmentRequestReceivedNodeID.Value) != len(nodeID.Value) {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong node ID value length.\n- Sent node ID value length: %v\n- Received node ID value length %v\n", len(nodeID.Value), len(pfcpSessionEstablishmentRequestReceivedNodeID.Value))
	}

	for i := range nodeID.Value {
		if pfcpSessionEstablishmentRequestReceivedNodeID.Value[i] != nodeID.Value[i] {
			t.Errorf("PFCP Session Establishment Request handler was called with wrong node ID value.\n- Sent node ID value: %v\n- Received node ID value %v\n", nodeID.Value, pfcpSessionEstablishmentRequestReceivedNodeID.Value)
		}
	}

	if pfcpSessionEstablishmentRequestReceivedCPFSEID.Header.Length != fseid.Header.Length {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong FSEID length.\n- Sent FSEID length: %v\n- Received FSEID length %v\n", fseid.Header.Length, pfcpSessionEstablishmentRequestReceivedCPFSEID.Header.Length)
	}

	if pfcpSessionEstablishmentRequestReceivedCPFSEID.Header.Type != fseid.Header.Type {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong FSEID type.\n- Sent FSEID type: %v\n- Received FSEID type %v\n", fseid.Header.Type, pfcpSessionEstablishmentRequestReceivedCPFSEID.Header.Type)
	}

	if pfcpSessionEstablishmentRequestReceivedCPFSEID.V4 != fseid.V4 {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong FSEID V4.\n- Sent FSEID V4: %v\n- Received FSEID V4 %v\n", fseid.V4, pfcpSessionEstablishmentRequestReceivedCPFSEID.V4)
	}

	if pfcpSessionEstablishmentRequestReceivedCPFSEID.V6 != fseid.V6 {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong FSEID V6.\n- Sent FSEID V6: %v\n- Received FSEID V6 %v\n", fseid.V6, pfcpSessionEstablishmentRequestReceivedCPFSEID.V6)
	}

	if pfcpSessionEstablishmentRequestReceivedCPFSEID.SEID != fseid.SEID {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong FSEID SEID.\n- Sent FSEID SEID: %v\n- Received FSEID SEID %v\n", fseid.SEID, pfcpSessionEstablishmentRequestReceivedCPFSEID.SEID)
	}

	if len(pfcpSessionEstablishmentRequestReceivedCPFSEID.IPv4) != len(fseid.IPv4) {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong FSEID IPv4 length.\n- Sent FSEID IPv4 length: %v\n- Received FSEID IPv4 length %v\n", len(fseid.IPv4), len(pfcpSessionEstablishmentRequestReceivedCPFSEID.IPv4))
	}

	for i := range fseid.IPv4 {
		if pfcpSessionEstablishmentRequestReceivedCPFSEID.IPv4[i] != fseid.IPv4[i] {
			t.Errorf("PFCP Session Establishment Request handler was called with wrong FSEID IPv4.\n- Sent FSEID IPv4: %v\n- Received FSEID IPv4 %v\n", fseid.IPv4, pfcpSessionEstablishmentRequestReceivedCPFSEID.IPv4)
		}
	}

	if len(pfcpSessionEstablishmentRequestReceivedCPFSEID.IPv6) != len(fseid.IPv6) {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong FSEID IPv6 length.\n- Sent FSEID IPv6 length: %v\n- Received FSEID IPv6 length %v\n", len(fseid.IPv6), len(pfcpSessionEstablishmentRequestReceivedCPFSEID.IPv6))
	}

	for i := range fseid.IPv6 {
		if pfcpSessionEstablishmentRequestReceivedCPFSEID.IPv6[i] != fseid.IPv6[i] {
			t.Errorf("PFCP Session Establishment Request handler was called with wrong FSEID IPv6.\n- Sent FSEID IPv6: %v\n- Received FSEID IPv6 %v\n", fseid.IPv6, pfcpSessionEstablishmentRequestReceivedCPFSEID.IPv6)
		}
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.Header.Length != createPDR.Header.Length {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR length.\n- Sent CreatePDR length: %v\n- Received CreatePDR length %v\n", createPDR.Header.Length, pfcpSessionEstablishmentRequestReceivedCreatePDR.Header.Length)
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.Header.Type != createPDR.Header.Type {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR type.\n- Sent CreatePDR type: %v\n- Received CreatePDR type %v\n", createPDR.Header.Type, pfcpSessionEstablishmentRequestReceivedCreatePDR.Header.Type)
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.PDRID != createPDR.PDRID {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR PDRID.\n- Sent CreatePDR PDRID: %v\n- Received CreatePDR PDRID %v\n", createPDR.PDRID, pfcpSessionEstablishmentRequestReceivedCreatePDR.PDRID)
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.Precedence != createPDR.Precedence {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR Precedence.\n- Sent CreatePDR Precedence: %v\n- Received CreatePDR Precedence %v\n", createPDR.Precedence, pfcpSessionEstablishmentRequestReceivedCreatePDR.Precedence)
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.Header.Length != createPDR.PDI.Header.Length {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR PDI length.\n- Sent CreatePDR PDI length: %v\n- Received CreatePDR PDI length %v\n", createPDR.PDI.Header.Length, pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.Header.Length)
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.Header.Type != createPDR.PDI.Header.Type {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR PDI type.\n- Sent CreatePDR PDI type: %v\n- Received CreatePDR PDI type %v\n", createPDR.PDI.Header.Type, pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.Header.Type)
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.SourceInterface.Header.Length != createPDR.PDI.SourceInterface.Header.Length {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR PDI SourceInterface length.\n- Sent CreatePDR PDI SourceInterface length: %v\n- Received CreatePDR PDI SourceInterface length %v\n", createPDR.PDI.SourceInterface.Header.Length, pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.SourceInterface.Header.Length)
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.SourceInterface.Header.Type != createPDR.PDI.SourceInterface.Header.Type {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR PDI SourceInterface type.\n- Sent CreatePDR PDI SourceInterface type: %v\n- Received CreatePDR PDI SourceInterface type %v\n", createPDR.PDI.SourceInterface.Header.Type, pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.SourceInterface.Header.Type)
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.SourceInterface.Value != createPDR.PDI.SourceInterface.Value {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR PDI SourceInterface Value.\n- Sent CreatePDR PDI SourceInterface Value: %v\n- Received CreatePDR PDI SourceInterface Value %v\n", createPDR.PDI.SourceInterface.Value, pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.SourceInterface.Value)
	}

	if pfcpSessionEstablishmentRequestReceivedCreateFAR.Header.Type != createFAR.Header.Type {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreateFAR IEType.\n- Sent CreateFAR IEType: %v\n- Received CreateFAR IEType %v\n", createFAR.Header.Type, pfcpSessionEstablishmentRequestReceivedCreateFAR.Header.Type)
	}

	if pfcpSessionEstablishmentRequestReceivedCreateFAR.Header.Length != createFAR.Header.Length {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreateFAR Length.\n- Sent CreateFAR Length: %v\n- Received CreateFAR Length %v\n", createFAR.Header.Length, pfcpSessionEstablishmentRequestReceivedCreateFAR.Header.Length)
	}

	if pfcpSessionEstablishmentRequestReceivedCreateFAR.FARID != createFAR.FARID {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreateFAR FARID.\n- Sent CreateFAR FARID: %v\n- Received CreateFAR FARID %v\n", createFAR.FARID, pfcpSessionEstablishmentRequestReceivedCreateFAR.FARID)
	}

	if pfcpSessionEstablishmentRequestReceivedCreateFAR.ApplyAction.Header.Length != createFAR.ApplyAction.Header.Length {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreateFAR ApplyAction Length.\n- Sent CreateFAR ApplyAction Length: %v\n- Received CreateFAR ApplyAction Length %v\n", createFAR.ApplyAction.Header.Length, pfcpSessionEstablishmentRequestReceivedCreateFAR.ApplyAction.Header.Length)
	}

	if pfcpSessionEstablishmentRequestReceivedCreateFAR.ApplyAction.Header.Type != createFAR.ApplyAction.Header.Type {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreateFAR ApplyAction IEType.\n- Sent CreateFAR ApplyAction IEType: %v\n- Received CreateFAR ApplyAction IEType %v\n", createFAR.ApplyAction.Header.Type, pfcpSessionEstablishmentRequestReceivedCreateFAR.ApplyAction.Header.Type)
	}

	if pfcpSessionEstablishmentRequestReceivedCreateFAR.ApplyAction.FORW != createFAR.ApplyAction.FORW {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreateFAR ApplyAction Forw.\n- Sent CreateFAR ApplyAction Forw: %v\n- Received CreateFAR ApplyAction Forw %v\n", createFAR.ApplyAction.FORW, pfcpSessionEstablishmentRequestReceivedCreateFAR.ApplyAction.FORW)
	}

	pfcpSessionEstablishmentRequestMu.Unlock()
}

func PFCPSessionEstablishmentResponse(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPSessionEstablishmentResponse(HandlePFCPSessionEstablishmentResponse)
	go pfcpServer.Run()
	defer pfcpServer.Close()

	time.Sleep(time.Second)
	pfcpClient := client.New("127.0.0.1:8805")

	nodeID, err := ie.NewNodeID("")
	if err != nil {
		t.Fatalf("Error creating Node ID: %v", err)
	}

	cause, err := ie.NewCause(ie.RequestAccepted)
	if err != nil {
		t.Fatalf("Error creating Cause: %v", err)
	}

	PFCPSessionEstablishmentResponseMsg := messages.PFCPSessionEstablishmentResponse{
		NodeID: nodeID,
		Cause:  cause,
	}

	sequenceNumber := uint32(32)
	seid := uint64(1234567890)

	err = pfcpClient.SendPFCPSessionEstablishmentResponse(&PFCPSessionEstablishmentResponseMsg, seid, sequenceNumber)
	if err != nil {
		t.Fatalf("Error sending PFCP Session Establishment Response: %v", err)
	}

	time.Sleep(time.Second)

	pfcpSessionEstablishmentResponseMu.Lock()

	if !pfcpSessionEstablishmentResponsehandlerCalled {
		t.Fatalf("PFCP Session Establishment Response handler was not called")
	}

	if pfcpSessionEstablishmentResponseReceivedSequenceNumber != sequenceNumber {
		t.Errorf("PFCP Session Establishment Response handler was called with wrong sequence number.\n- Sent sequence number: %v\n- Received sequence number %v\n", sequenceNumber, pfcpSessionEstablishmentResponseReceivedSequenceNumber)
	}

	if pfcpSessionEstablishmentResponseReceivedNodeID.Header.Length != nodeID.Header.Length {
		t.Errorf("PFCP Session Establishment Response handler was called with wrong node ID length.\n- Sent node ID length: %v\n- Received node ID length %v\n", nodeID.Header.Length, pfcpSessionEstablishmentResponseReceivedNodeID.Header.Length)
	}

	if pfcpSessionEstablishmentResponseReceivedNodeID.Type != nodeID.Type {
		t.Errorf("PFCP Session Establishment Response handler was called with wrong node ID type.\n- Sent node ID type: %v\n- Received node ID type %v\n", nodeID.Type, pfcpSessionEstablishmentResponseReceivedNodeID.Type)
	}

	if len(pfcpSessionEstablishmentResponseReceivedNodeID.Value) != len(nodeID.Value) {
		t.Errorf("PFCP Session Establishment Response handler was called with wrong node ID value length.\n- Sent node ID value length: %v\n- Received node ID value length %v\n", len(nodeID.Value), len(pfcpSessionEstablishmentResponseReceivedNodeID.Value))
	}

	for i := range nodeID.Value {
		if pfcpSessionEstablishmentResponseReceivedNodeID.Value[i] != nodeID.Value[i] {
			t.Errorf("PFCP Session Establishment Response handler was called with wrong node ID value.\n- Sent node ID value: %v\n- Received node ID value %v\n", nodeID.Value, pfcpSessionEstablishmentResponseReceivedNodeID.Value)
		}
	}

	if pfcpSessionEstablishmentResponseReceivedCause.Header.Length != cause.Header.Length {
		t.Errorf("PFCP Session Establishment Response handler was called with wrong cause length.\n- Sent cause length: %v\n- Received cause length %v\n", cause.Header.Length, pfcpSessionEstablishmentResponseReceivedCause.Header.Length)
	}

	if pfcpSessionEstablishmentResponseReceivedCause.Header.Type != cause.Header.Type {
		t.Errorf("PFCP Session Establishment Response handler was called with wrong cause type.\n- Sent cause type: %v\n- Received cause type %v\n", cause.Header.Type, pfcpSessionEstablishmentResponseReceivedCause.Header.Type)
	}

	if pfcpSessionEstablishmentResponseReceivedCause.Value != cause.Value {
		t.Errorf("PFCP Session Establishment Response handler was called with wrong cause value.\n- Sent cause value: %v\n- Received cause value %v\n", cause.Value, pfcpSessionEstablishmentResponseReceivedCause.Value)
	}

	pfcpSessionEstablishmentResponseMu.Unlock()

}
