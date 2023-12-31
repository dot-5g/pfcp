package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectRuleIDWhenNewPdrIDThenFieldsSetCorrectly(t *testing.T) {
	ruleID := uint16(1234)

	pdrID := ie.NewPDRID(ruleID)

	if pdrID.IEType != 56 {
		t.Errorf("Expected IEType %d, got %d", 56, pdrID.IEType)
	}

	if pdrID.Length != 2 {
		t.Errorf("Expected Length %d, got %d", 2, pdrID.Length)
	}

	if pdrID.RuleID != ruleID {
		t.Errorf("Expected RuleID %d, got %d", ruleID, pdrID.RuleID)
	}

}

func TestGivenPDRIDSerializedWhenDeserializeThenFieldsSetCorrectly(t *testing.T) {
	ruleID := uint16(1234)
	pdrID := ie.NewPDRID(ruleID)

	pdrIDSerialized := pdrID.Serialize()

	deserializedPDRID, err := ie.DeserializePDRID(56, 2, pdrIDSerialized[4:])

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deserializedPDRID.IEType != 56 {
		t.Errorf("Expected IEType %d, got %d", 56, deserializedPDRID.IEType)
	}

	if deserializedPDRID.Length != 2 {
		t.Errorf("Expected Length %d, got %d", 2, deserializedPDRID.Length)
	}

	if deserializedPDRID.RuleID != ruleID {
		t.Errorf("Expected RuleID %d, got %d", ruleID, deserializedPDRID.RuleID)
	}
}
