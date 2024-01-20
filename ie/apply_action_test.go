package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectValuesWhenNewApplyActionThenFieldsSetCorrectly(t *testing.T) {
	flag := ie.FORW
	applyAction, err := ie.NewApplyAction(flag, []ie.ApplyActionExtraFlag{ie.DFRT, ie.EDRT})

	if err != nil {
		t.Fatalf("Error creating ApplyAction: %v", err)
	}

	if applyAction.Header.Type != 44 {
		t.Errorf("Expected IEType 44, got %d", applyAction.Header.Type)
	}

	if applyAction.Header.Length != 2 {
		t.Errorf("Expected Length 2, got %d", applyAction.Header.Length)
	}

	if applyAction.FORW != true {
		t.Errorf("Expected FORW %v, got %v", flag, applyAction.FORW)
	}

	if applyAction.DFRT != true {
		t.Errorf("Expected DFRT %v, got %v", flag, applyAction.DFRT)
	}

	if applyAction.EDRT != true {
		t.Errorf("Expected EDRT %v, got %v", flag, applyAction.EDRT)
	}

	if applyAction.DROP != false {
		t.Errorf("Expected DROP %v, got %v", flag, applyAction.DROP)
	}

	if applyAction.BUFF != false {
		t.Errorf("Expected BUFF %v, got %v", flag, applyAction.BUFF)
	}

	if applyAction.IPMA != false {
		t.Errorf("Expected IPMA %v, got %v", flag, applyAction.IPMA)
	}

	if applyAction.IPMD != false {
		t.Errorf("Expected IPMD %v, got %v", flag, applyAction.IPMD)
	}

	if applyAction.DUPL != false {
		t.Errorf("Expected DUPL %v, got %v", flag, applyAction.DUPL)
	}

	if applyAction.NOCP != false {
		t.Errorf("Expected NOCP %v, got %v", flag, applyAction.NOCP)
	}

	if applyAction.DDPN != false {
		t.Errorf("Expected DDPN %v, got %v", flag, applyAction.DDPN)
	}

	if applyAction.BDPN != false {
		t.Errorf("Expected BDPN %v, got %v", flag, applyAction.BDPN)
	}
}

func TestGivenNOCPFlagUsedWithFORWWhenNewApplyActionThenErrorReturned(t *testing.T) {
	flag := ie.FORW
	_, err := ie.NewApplyAction(flag, []ie.ApplyActionExtraFlag{ie.NOCP})

	if err == nil {
		t.Errorf("Expected error creating ApplyAction, got nil")
	}
}

func TestGivenDUPLUsedWithIPMAWhenNewApplyActionThenErrorReturned(t *testing.T) {
	flag := ie.IPMA
	_, err := ie.NewApplyAction(flag, []ie.ApplyActionExtraFlag{ie.DUPL})

	if err == nil {
		t.Errorf("Expected error creating ApplyAction, got nil")
	}
}

func TestGivenDFRTUsedWithNonFORWWhenNewApplyActionThenErrorReturned(t *testing.T) {
	flag := ie.BUFF
	_, err := ie.NewApplyAction(flag, []ie.ApplyActionExtraFlag{ie.DFRT})

	if err == nil {
		t.Errorf("Expected error creating ApplyAction, got nil")
	}
}

func TestGivenEDRTUsedWithNonFORWWhenNewApplyActionThenErrorReturned(t *testing.T) {
	flag := ie.BUFF
	_, err := ie.NewApplyAction(flag, []ie.ApplyActionExtraFlag{ie.EDRT})

	if err == nil {
		t.Errorf("Expected error creating ApplyAction, got nil")
	}
}

func TestGivenApplyActionSerializedWhenDeserializeThenFieldsSetCorrectly(t *testing.T) {
	flag := ie.FORW

	applyAction, err := ie.NewApplyAction(flag, []ie.ApplyActionExtraFlag{})

	if err != nil {
		t.Fatalf("Error creating ApplyAction: %v", err)
	}

	serialized := applyAction.Serialize()

	deserialized, err := ie.DeserializeApplyAction(serialized[4:])

	if err != nil {
		t.Fatalf("Error deserializing ApplyAction: %v", err)
	}

	if deserialized.DFRT != false {
		t.Errorf("Expected no DFRT, got %v", deserialized.DFRT)
	}

	if deserialized.EDRT != false {
		t.Errorf("Expected no EDRT, got %v", deserialized.EDRT)
	}

	if deserialized.IPMD != false {
		t.Errorf("Expected no IPMD, got %v", deserialized.IPMD)
	}

	if deserialized.IPMA != false {
		t.Errorf("Expected no IPMA, got %v", deserialized.IPMA)
	}

	if deserialized.DUPL != false {
		t.Errorf("Expected no DUPL, got %v", deserialized.DUPL)
	}

	if deserialized.NOCP != false {
		t.Errorf("Expected no NOCP, got %v", deserialized.NOCP)
	}

	if deserialized.BUFF != false {
		t.Errorf("Expected no BUFF, got %v", deserialized.BUFF)
	}

	if deserialized.FORW != true {
		t.Errorf("Expected FORW %v, got %v", flag, deserialized.FORW)
	}

	if deserialized.DROP != false {
		t.Errorf("Expected no DROP, got %v", deserialized.DROP)
	}

	if deserialized.DDPN != false {
		t.Errorf("Expected no DDPN, got %v", deserialized.DDPN)
	}

	if deserialized.BDPN != false {
		t.Errorf("Expected no BDPN, got %v", deserialized.BDPN)
	}
}
