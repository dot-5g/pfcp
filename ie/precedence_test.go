package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectPrecedenceValueWhenNewPrecedenceThenFieldsSetCorrectly(t *testing.T) {
	precedenceValue := uint32(123)

	precedence := ie.NewPrecedence(precedenceValue)

	if precedence.IEType != 29 {
		t.Errorf("Expected IEType %d, got %d", 29, precedence.IEType)
	}

	if precedence.Length != 4 {
		t.Errorf("Expected Length %d, got %d", 4, precedence.Length)
	}

	if precedence.Value != precedenceValue {
		t.Errorf("Expected PrecedenceValue %d, got %d", precedenceValue, precedence.Value)
	}
}

func TestGivenPrecedenceSerializedWhenDeserializeThenFieldsSetCorrectly(t *testing.T) {
	precedenceValue := uint32(123)
	precedence := ie.NewPrecedence(precedenceValue)

	precedenceSerialized := precedence.Serialize()

	deserializedPrecedence, err := ie.DeserializePrecedence(29, 4, precedenceSerialized[4:])

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deserializedPrecedence.IEType != 29 {
		t.Errorf("Expected IEType %d, got %d", 29, deserializedPrecedence.IEType)
	}

	if deserializedPrecedence.Length != 4 {
		t.Errorf("Expected Length %d, got %d", 4, deserializedPrecedence.Length)
	}

	if deserializedPrecedence.Value != precedenceValue {
		t.Errorf("Expected PrecedenceValue %d, got %d", precedenceValue, deserializedPrecedence.Value)
	}
}
