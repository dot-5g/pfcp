package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectPDIWhenNewPDIThenFieldsSetCorrectly(t *testing.T) {
	sourceInterface, err := ie.NewSourceInterface(4)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	pdi, err := ie.NewPDI(sourceInterface)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if pdi.IEType != 17 {
		t.Errorf("Expected IEType %d, got %d", 17, pdi.IEType)
	}

	if pdi.Length != 5 {
		t.Errorf("Expected Length %d, got %d", 5, pdi.Length)
	}

	if pdi.SourceInterface != sourceInterface {
		t.Errorf("Expected SourceInterface %v, got %v", sourceInterface, pdi.SourceInterface)
	}

}

func TestGivenPDISerializedWhenDeserializeThenFieldsSetCorrectly(t *testing.T) {
	sourceInterface, err := ie.NewSourceInterface(4)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	pdi, err := ie.NewPDI(sourceInterface)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	pdiSerialized := pdi.Serialize()

	deserializedPDI, err := ie.DeserializePDI(17, 5, pdiSerialized[4:])

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deserializedPDI.IEType != 17 {
		t.Errorf("Expected IEType %d, got %d", 17, deserializedPDI.IEType)
	}

	if deserializedPDI.Length != 5 {
		t.Errorf("Expected Length %d, got %d", 5, deserializedPDI.Length)
	}

	if deserializedPDI.SourceInterface != sourceInterface {
		t.Errorf("Expected SourceInterface %v, got %v", sourceInterface, deserializedPDI.SourceInterface)
	}

}
