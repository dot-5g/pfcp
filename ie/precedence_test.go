package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectPrecedenceValueWhenNewPrecedenceThenFieldsSetCorrectly(t *testing.T) {
	precedenceValue := uint32(123)

	precedence, err := ie.NewPrecedence(precedenceValue)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if precedence.Value != precedenceValue {
		t.Errorf("Expected PrecedenceValue %d, got %d", precedenceValue, precedence.Value)
	}
}

func TestGivenPrecedenceSerializedWhenDeserializeThenFieldsSetCorrectly(t *testing.T) {
	precedenceValue := uint32(123)
	precedence, err := ie.NewPrecedence(precedenceValue)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	precedenceSerialized := precedence.Serialize()

	deserializedPrecedence, err := ie.DeserializePrecedence(precedenceSerialized)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deserializedPrecedence.Value != precedenceValue {
		t.Errorf("Expected PrecedenceValue %d, got %d", precedenceValue, deserializedPrecedence.Value)
	}
}
