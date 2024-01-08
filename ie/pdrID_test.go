package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectRuleIDWhenNewPdrIDThenFieldsSetCorrectly(t *testing.T) {
	ruleID := uint16(1234)

	pdrID, err := ie.NewPDRID(ruleID)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if pdrID.Header.Type != 56 {
		t.Errorf("Expected IEType %d, got %d", 56, pdrID.Header.Type)
	}

	if pdrID.Header.Length != 2 {
		t.Errorf("Expected Length %d, got %d", 2, pdrID.Header.Length)
	}

	if pdrID.RuleID != ruleID {
		t.Errorf("Expected RuleID %d, got %d", ruleID, pdrID.RuleID)
	}

}

func TestGivenPDRIDSerializedWhenDeserializeThenFieldsSetCorrectly(t *testing.T) {
	ruleID := uint16(1234)
	pdrID, err := ie.NewPDRID(ruleID)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	pdrIDSerialized := pdrID.Serialize()

	ieHeader := ie.IEHeader{
		Type:   56,
		Length: 2,
	}
	deserializedPDRID, err := ie.DeserializePDRID(ieHeader, pdrIDSerialized[4:])

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deserializedPDRID.Header.Type != 56 {
		t.Errorf("Expected IEType %d, got %d", 56, deserializedPDRID.Header.Type)
	}

	if deserializedPDRID.Header.Length != 2 {
		t.Errorf("Expected Length %d, got %d", 2, deserializedPDRID.Header.Length)
	}

	if deserializedPDRID.RuleID != ruleID {
		t.Errorf("Expected RuleID %d, got %d", ruleID, deserializedPDRID.RuleID)
	}
}
