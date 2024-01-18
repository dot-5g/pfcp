package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectValuesWhenNewFarThenFieldsSetCorrectly(t *testing.T) {
	farId, err := ie.NewFarID(1)

	if err != nil {
		t.Fatalf("Error creating FARID: %v", err)
	}

	flag := ie.FORW
	applyAction, err := ie.NewApplyAction(flag, []ie.ApplyActionExtraFlag{})

	if err != nil {
		t.Fatalf("Error creating ApplyAction: %v", err)
	}

	far, err := ie.NewFAR(farId, applyAction)

	if err != nil {
		t.Fatalf("Error creating far: %v", err)
	}

	if far.Header.Type != 3 {
		t.Errorf("Expected IEType 3, got %d", far.Header.Type)
	}

	if far.Header.Length != 14 {
		t.Errorf("Expected Length 14, got %d", far.Header.Length)
	}

	if far.FARID.Value != 1 {
		t.Errorf("Expected FARID 1, got %d", far.FARID.Value)
	}

	if far.ApplyAction.FORW != true {
		t.Errorf("Expected FORW true, got %v", far.ApplyAction.FORW)
	}

	if far.ApplyAction.DFRT != false {
		t.Errorf("Expected DFRT false, got %v", far.ApplyAction.DFRT)
	}

	if far.ApplyAction.EDRT != false {
		t.Errorf("Expected EDRT false, got %v", far.ApplyAction.EDRT)
	}

	if far.ApplyAction.DROP != false {
		t.Errorf("Expected DROP false, got %v", far.ApplyAction.DROP)
	}

	if far.ApplyAction.BUFF != false {
		t.Errorf("Expected BUFF false, got %v", far.ApplyAction.BUFF)
	}

	if far.ApplyAction.IPMA != false {
		t.Errorf("Expected IPMA false, got %v", far.ApplyAction.IPMA)
	}

	if far.ApplyAction.IPMD != false {
		t.Errorf("Expected IPMD false, got %v", far.ApplyAction.IPMD)
	}

	if far.ApplyAction.DUPL != false {
		t.Errorf("Expected DUPL false, got %v", far.ApplyAction.DUPL)
	}

	if far.ApplyAction.NOCP != false {
		t.Errorf("Expected NOCP false, got %v", far.ApplyAction.NOCP)
	}

	if far.ApplyAction.DDPN != false {
		t.Errorf("Expected DDPN false, got %v", far.ApplyAction.DDPN)
	}

	if far.ApplyAction.BDPN != false {
		t.Errorf("Expected BDPN false, got %v", far.ApplyAction.BDPN)
	}

	if far.ApplyAction.Header.Length != 2 {
		t.Errorf("Expected Length 2, got %d", far.ApplyAction.Header.Length)
	}

	if far.ApplyAction.Header.Type != 44 {
		t.Errorf("Expected IEType 44, got %d", far.ApplyAction.Header.Type)
	}

}

func TestGivenSerializedWhenDeserializeFarThenFieldsSetCorrectly(t *testing.T) {
	farId, err := ie.NewFarID(1)

	if err != nil {
		t.Fatalf("Error creating FARID: %v", err)
	}

	flag := ie.FORW
	applyAction, err := ie.NewApplyAction(flag, []ie.ApplyActionExtraFlag{})

	if err != nil {
		t.Fatalf("Error creating ApplyAction: %v", err)
	}
	far, err := ie.NewFAR(farId, applyAction)

	if err != nil {
		t.Fatalf("Error creating FAR: %v", err)
	}

	serialized := far.Serialize()

	ieHeader := ie.Header{
		Type:   3,
		Length: 14,
	}

	deserialized, err := ie.DeserializeFAR(ieHeader, serialized[4:])

	if err != nil {
		t.Fatalf("Error deserializing FAR: %v", err)
	}

	if deserialized.Header.Type != 3 {
		t.Errorf("Expected IEType 3, got %d", deserialized.Header.Type)
	}

	if deserialized.Header.Length != 14 {
		t.Errorf("Expected Length 14, got %d", deserialized.Header.Length)
	}

	if deserialized.FARID != farId {
		t.Errorf("Expected FARID %v, got %v", farId, deserialized.FARID)
	}

	if deserialized.ApplyAction != applyAction {
		t.Errorf("Expected ApplyAction %v, got %v", applyAction, deserialized.ApplyAction)
	}
}
