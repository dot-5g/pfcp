package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectFarIDValueWhenNewFarIDThenFieldsSetCorrectly(t *testing.T) {
	farIDValue := uint32(123)

	farID, err := ie.NewFarID(farIDValue)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if farID.Header.Type != 108 {
		t.Errorf("Expected IEType %d, got %d", 108, farID.Header.Type)
	}

	if farID.Header.Length != 4 {
		t.Errorf("Expected Length %d, got %d", 4, farID.Header.Length)
	}

	if farID.Value != farIDValue {
		t.Errorf("Expected FarIDValue %d, got %d", farIDValue, farID.Value)
	}
}

func TestGivenFarIDSerializedWhenDeserializeThenFieldsSetCorrectly(t *testing.T) {
	farIDValue := uint32(123)
	farID, err := ie.NewFarID(farIDValue)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	farIDSerialized := farID.Serialize()

	ieHeader := ie.Header{
		Type:   108,
		Length: 4,
	}
	deserializedFarID, err := ie.DeserializeFARID(ieHeader, farIDSerialized[4:])

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deserializedFarID.Header.Type != 108 {
		t.Errorf("Expected IEType %d, got %d", 108, deserializedFarID.Header.Type)
	}

	if deserializedFarID.Header.Length != 4 {
		t.Errorf("Expected Length %d, got %d", 4, deserializedFarID.Header.Length)
	}

	if deserializedFarID.Value != farIDValue {
		t.Errorf("Expected FarIDValue %d, got %d", farIDValue, deserializedFarID.Value)
	}
}
