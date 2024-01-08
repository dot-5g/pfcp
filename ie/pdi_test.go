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

	if pdi.Header.Type != 17 {
		t.Errorf("Expected IEType %d, got %d", 17, pdi.Header.Type)
	}

	if pdi.Header.Length != 5 {
		t.Errorf("Expected Length %d, got %d", 5, pdi.Header.Length)
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

	ieHeader := ie.Header{
		Type:   17,
		Length: 5,
	}

	deserializedPDI, err := ie.DeserializePDI(ieHeader, pdiSerialized[4:])

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deserializedPDI.Header.Type != 17 {
		t.Errorf("Expected IEType %d, got %d", 17, deserializedPDI.Header.Type)
	}

	if deserializedPDI.Header.Length != 5 {
		t.Errorf("Expected Length %d, got %d", 5, deserializedPDI.Header.Length)
	}

	if deserializedPDI.SourceInterface != sourceInterface {
		t.Errorf("Expected SourceInterface %v, got %v", sourceInterface, deserializedPDI.SourceInterface)
	}

}
