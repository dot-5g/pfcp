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

	if precedence.Header.Type != 29 {
		t.Errorf("Expected IEType %d, got %d", 29, precedence.Header.Type)
	}

	if precedence.Header.Length != 4 {
		t.Errorf("Expected Length %d, got %d", 4, precedence.Header.Length)
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
	ieHeader := ie.IEHeader{
		Type:   29,
		Length: 4,
	}

	deserializedPrecedence, err := ie.DeserializePrecedence(ieHeader, precedenceSerialized[4:])

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deserializedPrecedence.Header.Type != 29 {
		t.Errorf("Expected IEType %d, got %d", 29, deserializedPrecedence.Header.Type)
	}

	if deserializedPrecedence.Header.Length != 4 {
		t.Errorf("Expected Length %d, got %d", 4, deserializedPrecedence.Header.Length)
	}

	if deserializedPrecedence.Value != precedenceValue {
		t.Errorf("Expected PrecedenceValue %d, got %d", precedenceValue, deserializedPrecedence.Value)
	}
}
