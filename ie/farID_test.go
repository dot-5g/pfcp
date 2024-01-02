package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectFarIDValueWhenNewFarIDThenFieldsSetCorrectly(t *testing.T) {
	farIDValue := uint32(123)

	farID := ie.NewFarID(farIDValue)

	if farID.IEType != 108 {
		t.Errorf("Expected IEType %d, got %d", 108, farID.IEType)
	}

	if farID.Length != 4 {
		t.Errorf("Expected Length %d, got %d", 4, farID.Length)
	}

	if farID.Value != farIDValue {
		t.Errorf("Expected FarIDValue %d, got %d", farIDValue, farID.Value)
	}
}

func TestGivenFarIDSerializedWhenDeserializeThenFieldsSetCorrectly(t *testing.T) {
	farIDValue := uint32(123)
	farID := ie.NewFarID(farIDValue)

	farIDSerialized := farID.Serialize()

	deserializedFarID, err := ie.DeserializeFARID(108, 4, farIDSerialized[4:])

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deserializedFarID.IEType != 108 {
		t.Errorf("Expected IEType %d, got %d", 108, deserializedFarID.IEType)
	}

	if deserializedFarID.Length != 4 {
		t.Errorf("Expected Length %d, got %d", 4, deserializedFarID.Length)
	}

	if deserializedFarID.Value != farIDValue {
		t.Errorf("Expected FarIDValue %d, got %d", farIDValue, deserializedFarID.Value)
	}
}
