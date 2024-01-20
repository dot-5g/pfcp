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

	if sourceInterface.Header.Type != 20 {
		t.Errorf("Expected IEType %d, got %d", 20, sourceInterface.Header.Type)
	}

	if sourceInterface.Header.Length != 1 {
		t.Errorf("Expected Length %d, got %d", 1, sourceInterface.Header.Length)
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

	deserializedSourceInterface, err := ie.DeserializeSourceInterface(sourceInterfaceSerialized[4:])

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deserializedSourceInterface.Value != value {
		t.Errorf("Expected Value %d, got %d", value, deserializedSourceInterface.Value)
	}
}
