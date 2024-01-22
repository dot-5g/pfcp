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

	deserializedRecoveryTimeStamp, err := ie.DeserializeRecoveryTimeStamp(recoveryTimeStampSerialized)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Validate that secodns match num of seconds since 1900
	if deserializedRecoveryTimeStamp.Value != time.Unix()+2208988800 {
		t.Errorf("Expected Value %d, got %d", time.Unix()+2208988800, deserializedRecoveryTimeStamp.Value)
	}
}
