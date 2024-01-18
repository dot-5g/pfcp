package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectParametersWhenNewPDRThenFieldsSetCorrectly(t *testing.T) {
	pdrID, err := ie.NewPDRID(1)

	if err != nil {
		t.Fatalf("Error creating PDRID: %v", err)
	}

	precedence, err := ie.NewPrecedence(1)

	if err != nil {
		t.Fatalf("Error creating Precedence: %v", err)
	}

	sourceInterface, err := ie.NewSourceInterface(1)

	if err != nil {
		t.Fatalf("Error creating SourceInterface: %v", err)
	}

	pdi, err := ie.NewPDI(sourceInterface)

	if err != nil {
		t.Fatalf("Error creating PDI: %v", err)
	}

	pdr, err := ie.NewPDR(pdrID, precedence, pdi)

	if err != nil {
		t.Fatalf("Error creating pdr: %v", err)
	}

	if pdr.Header.Type != 1 {
		t.Errorf("Expected pdr IEType 1, got %d", pdr.Header.Type)
	}

	if pdr.Header.Length != 23 {
		t.Errorf("Expected pdr length 23, got %d", pdr.Header.Length)
	}

	if pdr.PDRID != pdrID {
		t.Errorf("Expected pdr PDRID %v, got %v", pdrID, pdr.PDRID)
	}

	if pdr.Precedence != precedence {
		t.Errorf("Expected pdr Precedence %v, got %v", precedence, pdr.Precedence)
	}

	if pdr.PDI != pdi {
		t.Errorf("Expected pdr PDI %v, got %v", pdi, pdr.PDI)
	}
}

func TestGivenSerializedWhenDeserializePDRThenFieldsSetCorrectly(t *testing.T) {
	pdrID, err := ie.NewPDRID(1)

	if err != nil {
		t.Fatalf("Error creating PDRID: %v", err)
	}

	precedence, err := ie.NewPrecedence(1)

	if err != nil {
		t.Fatalf("Error creating Precedence: %v", err)
	}

	sourceInterface, _ := ie.NewSourceInterface(1)
	pdi, err := ie.NewPDI(sourceInterface)

	if err != nil {
		t.Fatalf("Error creating PDI: %v", err)
	}

	pdr, err := ie.NewPDR(pdrID, precedence, pdi)

	if err != nil {
		t.Fatalf("Error creating PDR: %v", err)
	}

	serialized := pdr.Serialize()

	ieHeader := ie.Header{
		Type:   1,
		Length: 17,
	}

	deserialized, err := ie.DeserializePDR(ieHeader, serialized[4:])

	if err != nil {
		t.Fatalf("Error deserializing PDR: %v", err)
	}

	if deserialized.Header.Type != 1 {
		t.Errorf("Expected PDR IEType 1, got %d", deserialized.Header.Type)
	}

	if deserialized.Header.Length != 17 {
		t.Errorf("Expected PDR length 17, got %d", deserialized.Header.Length)
	}

	if deserialized.PDRID != pdrID {
		t.Errorf("Expected PDR PDRID %v, got %v", pdrID, deserialized.PDRID)
	}

	if deserialized.Precedence != precedence {
		t.Errorf("Expected PDR Precedence %v, got %v", precedence, deserialized.Precedence)
	}

	if deserialized.PDI != pdi {
		t.Errorf("Expected PDR PDI %v, got %v", pdi, deserialized.PDI)
	}
}
