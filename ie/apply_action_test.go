package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectValuesWhenNewApplyActionThenFieldsSetCorrectly(t *testing.T) {
	dfrt := true
	ipmd := false
	ipma := true
	dupl := false
	nocp := false
	buff := true
	forw := false
	drop := true
	ddpn := true
	bdpn := true
	edrt := false

	applyAction := ie.NewApplyAction(dfrt, ipmd, ipma, dupl, nocp, buff, forw, drop, ddpn, bdpn, edrt)

	if applyAction.IEType != 44 {
		t.Errorf("Expected IEType 44, got %d", applyAction.IEType)
	}

	if applyAction.Length != 2 {
		t.Errorf("Expected Length 2, got %d", applyAction.Length)
	}

	if applyAction.DFRT != dfrt {
		t.Errorf("Expected DFRT %v, got %v", dfrt, applyAction.DFRT)
	}

	if applyAction.IPMD != ipmd {
		t.Errorf("Expected IPMD %v, got %v", ipmd, applyAction.IPMD)
	}

	if applyAction.IPMA != ipma {
		t.Errorf("Expected IPMA %v, got %v", ipma, applyAction.IPMA)
	}

	if applyAction.DUPL != dupl {
		t.Errorf("Expected DUPL %v, got %v", dupl, applyAction.DUPL)
	}

	if applyAction.NOCP != nocp {
		t.Errorf("Expected NOCP %v, got %v", nocp, applyAction.NOCP)
	}

	if applyAction.BUFF != buff {
		t.Errorf("Expected BUFF %v, got %v", buff, applyAction.BUFF)
	}

	if applyAction.FORW != forw {
		t.Errorf("Expected FORW %v, got %v", forw, applyAction.FORW)
	}

	if applyAction.DROP != drop {
		t.Errorf("Expected DROP %v, got %v", drop, applyAction.DROP)
	}

	if applyAction.DDPN != ddpn {
		t.Errorf("Expected DDPN %v, got %v", ddpn, applyAction.DDPN)
	}

	if applyAction.BDPN != bdpn {
		t.Errorf("Expected BDPN %v, got %v", bdpn, applyAction.BDPN)
	}

	if applyAction.EDRT != edrt {
		t.Errorf("Expected EDRT %v, got %v", edrt, applyAction.EDRT)
	}

}

func TestGivenApplyActionSerializedWhenDeserializeThenFieldsSetCorrectly(t *testing.T) {
	dfrt := true
	ipmd := false
	ipma := true
	dupl := false
	nocp := false
	buff := true
	forw := false
	drop := true
	ddpn := true
	bdpn := true
	edrt := false

	applyAction := ie.NewApplyAction(dfrt, ipmd, ipma, dupl, nocp, buff, forw, drop, ddpn, bdpn, edrt)

	serialized := applyAction.Serialize()

	deserialized, err := ie.DeserializeApplyAction(44, 2, serialized[4:])

	if err != nil {
		t.Fatalf("Error deserializing ApplyAction: %v", err)
	}

	if deserialized.IEType != 44 {
		t.Errorf("Expected IEType 44, got %d", deserialized.IEType)
	}

	if deserialized.Length != 2 {
		t.Errorf("Expected Length 2, got %d", deserialized.Length)
	}

	if deserialized.DFRT != dfrt {
		t.Errorf("Expected DFRT %v, got %v", dfrt, deserialized.DFRT)
	}

	if deserialized.IPMD != ipmd {
		t.Errorf("Expected IPMD %v, got %v", ipmd, deserialized.IPMD)
	}

	if deserialized.IPMA != ipma {
		t.Errorf("Expected IPMA %v, got %v", ipma, deserialized.IPMA)
	}

	if deserialized.DUPL != dupl {
		t.Errorf("Expected DUPL %v, got %v", dupl, deserialized.DUPL)
	}

	if deserialized.NOCP != nocp {
		t.Errorf("Expected NOCP %v, got %v", nocp, deserialized.NOCP)
	}

	if deserialized.BUFF != buff {
		t.Errorf("Expected BUFF %v, got %v", buff, deserialized.BUFF)
	}

	if deserialized.FORW != forw {
		t.Errorf("Expected FORW %v, got %v", forw, deserialized.FORW)
	}

	if deserialized.DROP != drop {
		t.Errorf("Expected DROP %v, got %v", drop, deserialized.DROP)
	}

	if deserialized.DDPN != ddpn {
		t.Errorf("Expected DDPN %v, got %v", ddpn, deserialized.DDPN)
	}

	if deserialized.BDPN != bdpn {
		t.Errorf("Expected BDPN %v, got %v", bdpn, deserialized.BDPN)
	}

	if deserialized.EDRT != edrt {
		t.Errorf("Expected EDRT %v, got %v", edrt, deserialized.EDRT)
	}

}
