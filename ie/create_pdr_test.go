package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectParametersWhenNewCreatePDRThenFieldsSetCorrectly(t *testing.T) {
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

	createPDR, err := ie.NewCreatePDR(pdrID, precedence, pdi)

	if err != nil {
		t.Fatalf("Error creating CreatePDR: %v", err)
	}

	if createPDR.IEType != 1 {
		t.Errorf("Expected CreatePDR IEType 1, got %d", createPDR.IEType)
	}

	if createPDR.Length != 23 {
		t.Errorf("Expected CreatePDR length 23, got %d", createPDR.Length)
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

	createPDR, err := ie.NewCreatePDR(pdrID, precedence, pdi)

	if err != nil {
		t.Fatalf("Error creating CreatePDR: %v", err)
	}

	serialized := createPDR.Serialize()

	deserialized, err := ie.DeserializeCreatePDR(1, 17, serialized[4:])

	if err != nil {
		t.Fatalf("Error deserializing CreatePDR: %v", err)
	}

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
