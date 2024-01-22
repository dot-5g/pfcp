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

	createFar, err := ie.NewCreateFAR(farId, applyAction)

	if err != nil {
		t.Fatalf("Error creating CreateFAR: %v", err)
	}

	if createFar.FARID.Value != 1 {
		t.Errorf("Expected FARID 1, got %d", createFar.FARID.Value)
	}

	if createFar.ApplyAction.FORW != true {
		t.Errorf("Expected FORW true, got %v", createFar.ApplyAction.FORW)
	}

	if createFar.ApplyAction.DFRT != false {
		t.Errorf("Expected DFRT false, got %v", createFar.ApplyAction.DFRT)
	}

	if createFar.ApplyAction.EDRT != false {
		t.Errorf("Expected EDRT false, got %v", createFar.ApplyAction.EDRT)
	}

	if createFar.ApplyAction.DROP != false {
		t.Errorf("Expected DROP false, got %v", createFar.ApplyAction.DROP)
	}

	if createFar.ApplyAction.BUFF != false {
		t.Errorf("Expected BUFF false, got %v", createFar.ApplyAction.BUFF)
	}

	if createFar.ApplyAction.IPMA != false {
		t.Errorf("Expected IPMA false, got %v", createFar.ApplyAction.IPMA)
	}

	if createFar.ApplyAction.IPMD != false {
		t.Errorf("Expected IPMD false, got %v", createFar.ApplyAction.IPMD)
	}

	if createFar.ApplyAction.DUPL != false {
		t.Errorf("Expected DUPL false, got %v", createFar.ApplyAction.DUPL)
	}

	if createFar.ApplyAction.NOCP != false {
		t.Errorf("Expected NOCP false, got %v", createFar.ApplyAction.NOCP)
	}

	if createFar.ApplyAction.DDPN != false {
		t.Errorf("Expected DDPN false, got %v", createFar.ApplyAction.DDPN)
	}

	if createFar.ApplyAction.BDPN != false {
		t.Errorf("Expected BDPN false, got %v", createFar.ApplyAction.BDPN)
	}

}

func TestGivenSerializedWhenDeserializeCreateFarThenFieldsSetCorrectly(t *testing.T) {
	farId, err := ie.NewFarID(1)
	if err != nil {
		t.Fatalf("Error creating FARID: %v", err)
	}

	flag := ie.FORW
	applyAction, err := ie.NewApplyAction(flag, []ie.ApplyActionExtraFlag{})

	if err != nil {
		t.Fatalf("Error creating ApplyAction: %v", err)
	}
	createFar, err := ie.NewCreateFAR(farId, applyAction)

	if err != nil {
		t.Fatalf("Error creating CreateFAR: %v", err)
	}

	serialized := createFar.Serialize()

	deserialized, err := ie.DeserializeCreateFAR(serialized)

	if err != nil {
		t.Fatalf("Error deserializing CreateFAR: %v", err)
	}

	if deserialized.FARID != farId {
		t.Errorf("Expected FARID %v, got %v", farId, deserialized.FARID)
	}

	if deserialized.ApplyAction != applyAction {
		t.Errorf("Expected ApplyAction %v, got %v", applyAction, deserialized.ApplyAction)
	}
}
