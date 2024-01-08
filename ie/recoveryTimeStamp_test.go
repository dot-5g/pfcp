package ie_test

import (
	"testing"
	"time"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectTimeWhenNewRecoveryTimeStampThenFieldsSetCorrectly(t *testing.T) {
	time := time.Now()

	recoveryTimeStamp, err := ie.NewRecoveryTimeStamp(time)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if recoveryTimeStamp.Header.Type != 96 {
		t.Errorf("Expected IEType %d, got %d", 96, recoveryTimeStamp.Header.Type)
	}

	if recoveryTimeStamp.Header.Length != 4 {
		t.Errorf("Expected Length %d, got %d", 4, recoveryTimeStamp.Header.Length)
	}

	// Validate that secodns match num of seconds since 1900
	if recoveryTimeStamp.Value != time.Unix()+2208988800 {
		t.Errorf("Expected Value %d, got %d", time.Unix()+2208988800, recoveryTimeStamp.Value)
	}

}

func TestGivenRecoveryTimeStampSerializedWhenDeserializeThenFieldsSetCorrectly(t *testing.T) {
	time := time.Now()
	recoveryTimeStamp, err := ie.NewRecoveryTimeStamp(time)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	recoveryTimeStampSerialized := recoveryTimeStamp.Serialize()
	ieHeader := ie.IEHeader{
		Type:   96,
		Length: 4,
	}

	deserializedRecoveryTimeStamp, err := ie.DeserializeRecoveryTimeStamp(ieHeader, recoveryTimeStampSerialized[4:])

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deserializedRecoveryTimeStamp.Header.Type != 96 {
		t.Errorf("Expected IEType %d, got %d", 96, deserializedRecoveryTimeStamp.Header.Type)
	}

	if deserializedRecoveryTimeStamp.Header.Length != 4 {
		t.Errorf("Expected Length %d, got %d", 4, deserializedRecoveryTimeStamp.Header.Length)
	}

	// Validate that secodns match num of seconds since 1900
	if deserializedRecoveryTimeStamp.Value != time.Unix()+2208988800 {
		t.Errorf("Expected Value %d, got %d", time.Unix()+2208988800, deserializedRecoveryTimeStamp.Value)
	}
}
