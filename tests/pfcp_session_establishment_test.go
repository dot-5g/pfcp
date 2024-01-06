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

func TestPFCPSessionEstablishment(t *testing.T) {
	t.Run("TestPFCPSessionEstablishmentRequest", PFCPSessionEstablishmentRequest)
}

func PFCPSessionEstablishmentRequest(t *testing.T) {
	pfcpServer := server.New("127.0.0.1:8805")
	pfcpServer.PFCPSessionEstablishmentRequest(HandlePFCPSessionEstablishmentRequest)

	pfcpServer.Run()

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

	pfcpClient.SendPFCPSessionEstablishmentRequest(PFCPSessionEstablishmentRequestMsg, seid, sequenceNumber)

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

	if pfcpSessionEstablishmentRequestReceivedNodeID.Length != nodeID.Length {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong node ID length.\n- Sent node ID length: %v\n- Received node ID length %v\n", nodeID.Length, pfcpSessionEstablishmentRequestReceivedNodeID.Length)
	}

	if pfcpSessionEstablishmentRequestReceivedNodeID.NodeIDType != nodeID.NodeIDType {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong node ID type.\n- Sent node ID type: %v\n- Received node ID type %v\n", nodeID.NodeIDType, pfcpSessionEstablishmentRequestReceivedNodeID.NodeIDType)
	}

	if len(pfcpSessionEstablishmentRequestReceivedNodeID.NodeIDValue) != len(nodeID.NodeIDValue) {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong node ID value length.\n- Sent node ID value length: %v\n- Received node ID value length %v\n", len(nodeID.NodeIDValue), len(pfcpSessionEstablishmentRequestReceivedNodeID.NodeIDValue))
	}

	for i := range nodeID.NodeIDValue {
		if pfcpSessionEstablishmentRequestReceivedNodeID.NodeIDValue[i] != nodeID.NodeIDValue[i] {
			t.Errorf("PFCP Session Establishment Request handler was called with wrong node ID value.\n- Sent node ID value: %v\n- Received node ID value %v\n", nodeID.NodeIDValue, pfcpSessionEstablishmentRequestReceivedNodeID.NodeIDValue)
		}
	}

	if pfcpSessionEstablishmentRequestReceivedCPFSEID.Length != fseid.Length {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong FSEID length.\n- Sent FSEID length: %v\n- Received FSEID length %v\n", fseid.Length, pfcpSessionEstablishmentRequestReceivedCPFSEID.Length)
	}

	if pfcpSessionEstablishmentRequestReceivedCPFSEID.IEType != fseid.IEType {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong FSEID type.\n- Sent FSEID type: %v\n- Received FSEID type %v\n", fseid.IEType, pfcpSessionEstablishmentRequestReceivedCPFSEID.IEType)
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

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.Length != createPDR.Length {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR length.\n- Sent CreatePDR length: %v\n- Received CreatePDR length %v\n", createPDR.Length, pfcpSessionEstablishmentRequestReceivedCreatePDR.Length)
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.IEType != createPDR.IEType {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR type.\n- Sent CreatePDR type: %v\n- Received CreatePDR type %v\n", createPDR.IEType, pfcpSessionEstablishmentRequestReceivedCreatePDR.IEType)
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.PDRID != createPDR.PDRID {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR PDRID.\n- Sent CreatePDR PDRID: %v\n- Received CreatePDR PDRID %v\n", createPDR.PDRID, pfcpSessionEstablishmentRequestReceivedCreatePDR.PDRID)
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.Precedence != createPDR.Precedence {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR Precedence.\n- Sent CreatePDR Precedence: %v\n- Received CreatePDR Precedence %v\n", createPDR.Precedence, pfcpSessionEstablishmentRequestReceivedCreatePDR.Precedence)
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.Length != createPDR.PDI.Length {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR PDI length.\n- Sent CreatePDR PDI length: %v\n- Received CreatePDR PDI length %v\n", createPDR.PDI.Length, pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.Length)
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.IEType != createPDR.PDI.IEType {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR PDI type.\n- Sent CreatePDR PDI type: %v\n- Received CreatePDR PDI type %v\n", createPDR.PDI.IEType, pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.IEType)
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.SourceInterface.Length != createPDR.PDI.SourceInterface.Length {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR PDI SourceInterface length.\n- Sent CreatePDR PDI SourceInterface length: %v\n- Received CreatePDR PDI SourceInterface length %v\n", createPDR.PDI.SourceInterface.Length, pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.SourceInterface.Length)
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.SourceInterface.IEType != createPDR.PDI.SourceInterface.IEType {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR PDI SourceInterface type.\n- Sent CreatePDR PDI SourceInterface type: %v\n- Received CreatePDR PDI SourceInterface type %v\n", createPDR.PDI.SourceInterface.IEType, pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.SourceInterface.IEType)
	}

	if pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.SourceInterface.Value != createPDR.PDI.SourceInterface.Value {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreatePDR PDI SourceInterface Value.\n- Sent CreatePDR PDI SourceInterface Value: %v\n- Received CreatePDR PDI SourceInterface Value %v\n", createPDR.PDI.SourceInterface.Value, pfcpSessionEstablishmentRequestReceivedCreatePDR.PDI.SourceInterface.Value)
	}

	if pfcpSessionEstablishmentRequestReceivedCreateFAR.IEType != createFAR.IEType {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreateFAR IEType.\n- Sent CreateFAR IEType: %v\n- Received CreateFAR IEType %v\n", createFAR.IEType, pfcpSessionEstablishmentRequestReceivedCreateFAR.IEType)
	}

	if pfcpSessionEstablishmentRequestReceivedCreateFAR.Length != createFAR.Length {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreateFAR Length.\n- Sent CreateFAR Length: %v\n- Received CreateFAR Length %v\n", createFAR.Length, pfcpSessionEstablishmentRequestReceivedCreateFAR.Length)
	}

	if pfcpSessionEstablishmentRequestReceivedCreateFAR.FARID != createFAR.FARID {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreateFAR FARID.\n- Sent CreateFAR FARID: %v\n- Received CreateFAR FARID %v\n", createFAR.FARID, pfcpSessionEstablishmentRequestReceivedCreateFAR.FARID)
	}

	if pfcpSessionEstablishmentRequestReceivedCreateFAR.ApplyAction.Length != createFAR.ApplyAction.Length {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreateFAR ApplyAction Length.\n- Sent CreateFAR ApplyAction Length: %v\n- Received CreateFAR ApplyAction Length %v\n", createFAR.ApplyAction.Length, pfcpSessionEstablishmentRequestReceivedCreateFAR.ApplyAction.Length)
	}

	if pfcpSessionEstablishmentRequestReceivedCreateFAR.ApplyAction.IEType != createFAR.ApplyAction.IEType {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreateFAR ApplyAction IEType.\n- Sent CreateFAR ApplyAction IEType: %v\n- Received CreateFAR ApplyAction IEType %v\n", createFAR.ApplyAction.IEType, pfcpSessionEstablishmentRequestReceivedCreateFAR.ApplyAction.IEType)
	}

	if pfcpSessionEstablishmentRequestReceivedCreateFAR.ApplyAction.FORW != createFAR.ApplyAction.FORW {
		t.Errorf("PFCP Session Establishment Request handler was called with wrong CreateFAR ApplyAction Forw.\n- Sent CreateFAR ApplyAction Forw: %v\n- Received CreateFAR ApplyAction Forw %v\n", createFAR.ApplyAction.FORW, pfcpSessionEstablishmentRequestReceivedCreateFAR.ApplyAction.FORW)
	}

	pfcpSessionEstablishmentRequestMu.Unlock()
}
