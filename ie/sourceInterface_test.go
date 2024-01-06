package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectValueWhenNewSourceInterfaceThenFieldsSetCorrectly(t *testing.T) {
	value := 12

	sourceInterface, err := ie.NewSourceInterface(value)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if sourceInterface.IEType != 20 {
		t.Errorf("Expected IEType %d, got %d", 20, sourceInterface.IEType)
	}

	if sourceInterface.Length != 1 {
		t.Errorf("Expected Length %d, got %d", 1, sourceInterface.Length)
	}

	if sourceInterface.Value != value {
		t.Errorf("Expected Value %d, got %d", value, sourceInterface.Value)
	}
}

func TestGivenSourceInterfaceSerializedWhenDeserializeThenFieldsSetCorrectly(t *testing.T) {
	value := 9
	sourceInterface, err := ie.NewSourceInterface(value)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	sourceInterfaceSerialized := sourceInterface.Serialize()

	deserializedSourceInterface, err := ie.DeserializeSourceInterface(20, 1, sourceInterfaceSerialized[4:])

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deserializedSourceInterface.IEType != 20 {
		t.Errorf("Expected IEType %d, got %d", 20, deserializedSourceInterface.IEType)
	}

	if deserializedSourceInterface.Length != 1 {
		t.Errorf("Expected Length %d, got %d", 1, deserializedSourceInterface.Length)
	}

	if deserializedSourceInterface.Value != value {
		t.Errorf("Expected Value %d, got %d", value, deserializedSourceInterface.Value)
	}
}
