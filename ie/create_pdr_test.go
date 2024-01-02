package ie

import (
	"fmt"
	"testing"
)

func TestGivenCorrectParametersWhenNewCreatePDRThenFieldsSetCorrectly(t *testing.T) {
	pdrID := NewPDRID(1)
	precedence := NewPrecedence(1)
	sourceInterface, _ := NewSourceInterface(1)
	pdi := NewPDI(sourceInterface)
	createPDR := NewCreatePDR(pdrID, precedence, pdi)

	if createPDR.IEType != 1 {
		t.Errorf("Expected CreatePDR IEType 1, got %d", createPDR.IEType)
	}

	if createPDR.Length != 17 {
		t.Errorf("Expected CreatePDR length 17, got %d", createPDR.Length)
	}

	if createPDR.PDRID != pdrID {
		t.Errorf("Expected CreatePDR PDRID %v, got %v", pdrID, createPDR.PDRID)
	}

	if createPDR.Precedence != precedence {
		t.Errorf("Expected CreatePDR Precedence %v, got %v", precedence, createPDR.Precedence)
	}

	if createPDR.PDI != pdi {
		t.Errorf("Expected CreatePDR PDI %v, got %v", pdi, createPDR.PDI)
	}
}

func TestGivenSerializedWhenDeserializeCreatePDRThenFieldsSetCorrectly(t *testing.T) {
	pdrID := NewPDRID(1)
	precedence := NewPrecedence(1)
	sourceInterface, _ := NewSourceInterface(1)
	pdi := NewPDI(sourceInterface)
	createPDR := NewCreatePDR(pdrID, precedence, pdi)

	serialized := createPDR.Serialize()
	fmt.Printf("Serialized: %v\n", serialized)

	deserialized, err := DeserializeCreatePDR(1, 17, serialized[4:])

	if err != nil {
		t.Fatalf("Error deserializing CreatePDR: %v", err)
	}

	fmt.Printf("Deserialized: %v\n", deserialized)

	if deserialized.IEType != 1 {
		t.Errorf("Expected CreatePDR IEType 1, got %d", deserialized.IEType)
	}

	if deserialized.Length != 17 {
		t.Errorf("Expected CreatePDR length 17, got %d", deserialized.Length)
	}

	if deserialized.PDRID != pdrID {
		t.Errorf("Expected CreatePDR PDRID %v, got %v", pdrID, deserialized.PDRID)
	}

	if deserialized.Precedence != precedence {
		t.Errorf("Expected CreatePDR Precedence %v, got %v", precedence, deserialized.Precedence)
	}

	if deserialized.PDI != pdi {
		t.Errorf("Expected CreatePDR PDI %v, got %v", pdi, deserialized.PDI)
	}
}
