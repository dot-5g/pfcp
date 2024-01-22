package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectValueWhenNewReportTypeThenFieldsSetCorrectly(t *testing.T) {
	reports := []ie.Report{ie.UISR, ie.SESR}

	reportType, err := ie.NewReportType(reports)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(reportType.Reports) != 2 {
		t.Errorf("Expected 2 reports, got %d", len(reportType.Reports))
	}

	if reportType.Reports[0] != ie.UISR {
		t.Errorf("Expected report %d, got %d", ie.UISR, reportType.Reports[0])
	}

	if reportType.Reports[1] != ie.SESR {
		t.Errorf("Expected report %d, got %d", ie.SESR, reportType.Reports[1])
	}
}

func TestGivenSerializedWhenDeserializeReportTypeThenFieldsSetCorrectly(t *testing.T) {
	reports := []ie.Report{ie.UISR, ie.SESR}

	reportType, err := ie.NewReportType(reports)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	serializedReportType := reportType.Serialize()

	deserializedReportType, err := ie.DeserializeReportType(serializedReportType)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(deserializedReportType.Reports) != 2 {
		t.Errorf("Expected 2 reports, got %d", len(deserializedReportType.Reports))
	}

	if deserializedReportType.Reports[0] != ie.UISR {
		t.Errorf("Expected report %d, got %d", ie.UISR, deserializedReportType.Reports[0])
	}

	if deserializedReportType.Reports[1] != ie.SESR {
		t.Errorf("Expected report %d, got %d", ie.SESR, deserializedReportType.Reports[1])
	}

}
