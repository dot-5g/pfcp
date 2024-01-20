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

	deserializedFarID, err := ie.DeserializeFARID(farIDSerialized[4:])

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deserializedFarID.Value != farIDValue {
		t.Errorf("Expected FarIDValue %d, got %d", farIDValue, deserializedFarID.Value)
	}
}
