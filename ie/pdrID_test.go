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

	deserializedPDRID, err := ie.DeserializePDRID(pdrIDSerialized)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deserializedPDRID.RuleID != ruleID {
		t.Errorf("Expected RuleID %d, got %d", ruleID, deserializedPDRID.RuleID)
	}
}
