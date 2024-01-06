package ie_test

import (
	"fmt"
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectValuesWhenNewFarThenFieldsSetCorrectly(t *testing.T) {
	farId := ie.NewFarID(1)
	dfrt := false
	ipmd := false
	ipma := true
	dupl := true
	nocp := false
	buff := true
	forw := false
	drop := true
	ddpn := true
	bdpn := true
	edrt := false
	applyAction := ie.NewApplyAction(dfrt, ipmd, ipma, dupl, nocp, buff, forw, drop, ddpn, bdpn, edrt)

	createFar := ie.NewCreateFAR(farId, applyAction)

	if createFar.IEType != 3 {
		t.Errorf("Expected IEType 3, got %d", createFar.IEType)
	}

	if createFar.Length != 14 {
		t.Errorf("Expected Length 14, got %d", createFar.Length)
	}

	if createFar.FARID.Value != 1 {
		t.Errorf("Expected FARID 1, got %d", createFar.FARID.Value)
	}

	if createFar.ApplyAction.DFRT != dfrt {
		t.Errorf("Expected DFRT %v, got %v", dfrt, createFar.ApplyAction.DFRT)
	}

	if createFar.ApplyAction.IPMD != ipmd {
		t.Errorf("Expected IPMD %v, got %v", ipmd, createFar.ApplyAction.IPMD)
	}

	if createFar.ApplyAction.IPMA != ipma {
		t.Errorf("Expected IPMA %v, got %v", ipma, createFar.ApplyAction.IPMA)
	}

	if createFar.ApplyAction.DUPL != dupl {
		t.Errorf("Expected DUPL %v, got %v", dupl, createFar.ApplyAction.DUPL)
	}

	if createFar.ApplyAction.NOCP != nocp {
		t.Errorf("Expected NOCP %v, got %v", nocp, createFar.ApplyAction.NOCP)
	}

	if createFar.ApplyAction.BUFF != buff {
		t.Errorf("Expected BUFF %v, got %v", buff, createFar.ApplyAction.BUFF)
	}

	if createFar.ApplyAction.FORW != forw {
		t.Errorf("Expected FORW %v, got %v", forw, createFar.ApplyAction.FORW)
	}

	if createFar.ApplyAction.DROP != drop {
		t.Errorf("Expected DROP %v, got %v", drop, createFar.ApplyAction.DROP)
	}

	if createFar.ApplyAction.DDPN != ddpn {
		t.Errorf("Expected DDPN %v, got %v", ddpn, createFar.ApplyAction.DDPN)
	}

	if createFar.ApplyAction.BDPN != bdpn {
		t.Errorf("Expected BDPN %v, got %v", bdpn, createFar.ApplyAction.BDPN)
	}

	if createFar.ApplyAction.EDRT != edrt {
		t.Errorf("Expected EDRT %v, got %v", edrt, createFar.ApplyAction.EDRT)
	}
}

func TestGivenSerializedWhenDeserializeCreateFarThenFieldsSetCorrectly(t *testing.T) {
	farId := ie.NewFarID(1)
	dfrt := false
	ipmd := false
	ipma := true
	dupl := true
	nocp := false
	buff := true
	forw := false
	drop := true
	ddpn := true
	bdpn := true
	edrt := false
	applyAction := ie.NewApplyAction(dfrt, ipmd, ipma, dupl, nocp, buff, forw, drop, ddpn, bdpn, edrt)
	createFar := ie.NewCreateFAR(farId, applyAction)

	serialized := createFar.Serialize()

	fmt.Printf("Serialized: %v\n", serialized)
	fmt.Printf("Length of serialized %d\n", len(serialized))

	deserialized, err := ie.DeserializeCreateFAR(3, 14, serialized[4:])

	if err != nil {
		t.Fatalf("Error deserializing CreateFAR: %v", err)
	}

	if deserialized.IEType != 3 {
		t.Errorf("Expected IEType 3, got %d", deserialized.IEType)
	}

	if deserialized.Length != 14 {
		t.Errorf("Expected Length 14, got %d", deserialized.Length)
	}

	if deserialized.FARID != farId {
		t.Errorf("Expected FARID %v, got %v", farId, deserialized.FARID)
	}

	if deserialized.ApplyAction != applyAction {
		t.Errorf("Expected ApplyAction %v, got %v", applyAction, deserialized.ApplyAction)
	}
}
