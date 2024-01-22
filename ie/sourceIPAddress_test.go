package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectIPv4AddressWhenSourceIPAddressThenFieldsSetCorrectly(t *testing.T) {
	sourceIPAddress, err := ie.NewSourceIPAddress("1.2.3.4/24", "")

	if err != nil {
		t.Fatalf("Error creating SourceIPAddress: %v", err)
	}

	if sourceIPAddress.MPL != true {
		t.Errorf("Expected NodeID MPL true, got %v", sourceIPAddress.MPL)
	}

	if sourceIPAddress.V4 != true {
		t.Errorf("Expected NodeID V4 true, got %v", sourceIPAddress.V4)
	}

	if sourceIPAddress.V6 != false {
		t.Errorf("Expected NodeID V6 false, got %v", sourceIPAddress.V6)
	}

	if sourceIPAddress.MaskPrefixLength != 24 {
		t.Errorf("Expected NodeID MaskPrefixLength 24, got %d", sourceIPAddress.MaskPrefixLength)
	}
}

func TestGivenCorrectIPv6AddressWhenSourceIPAddressThenFieldsSetCorrectly(t *testing.T) {
	sourceIPAddress, err := ie.NewSourceIPAddress("", "2001:db8::/32")

	if err != nil {
		t.Fatalf("Error creating SourceIPAddress: %v", err)
	}

	if sourceIPAddress.MPL != true {
		t.Errorf("Expected NodeID MPL true, got %v", sourceIPAddress.MPL)
	}

	if sourceIPAddress.V4 != false {
		t.Errorf("Expected NodeID V4 false, got %v", sourceIPAddress.V4)
	}

	if sourceIPAddress.V6 != true {
		t.Errorf("Expected NodeID V6 true, got %v", sourceIPAddress.V6)
	}

	if sourceIPAddress.MaskPrefixLength != 32 {
		t.Errorf("Expected NodeID MaskPrefixLength 32, got %d", sourceIPAddress.MaskPrefixLength)
	}
}

func TestGivenSerializedAddressWhenDeserializeThenFieldsSetCorrectly(t *testing.T) {
	sourceIPAddress, err := ie.NewSourceIPAddress("2.2.3.1/24", "")

	if err != nil {
		t.Fatalf("Error creating SourceIPAddress: %v", err)
	}

	serializedSourceIPAddress := sourceIPAddress.Serialize()

	deserializedSourceIPAddress, err := ie.DeserializeSourceIPAddress(serializedSourceIPAddress)

	if err != nil {
		t.Fatalf("Error deserializing SourceIPAddress: %v", err)
	}

	if deserializedSourceIPAddress.MPL != true {
		t.Errorf("Expected NodeID MPL true, got %v", deserializedSourceIPAddress.MPL)
	}

	if deserializedSourceIPAddress.V4 != true {
		t.Errorf("Expected NodeID V4 true, got %v", deserializedSourceIPAddress.V4)
	}

	if deserializedSourceIPAddress.V6 != false {
		t.Errorf("Expected NodeID V6 false, got %v", deserializedSourceIPAddress.V6)
	}

	if deserializedSourceIPAddress.MaskPrefixLength != 24 {
		t.Errorf("Expected NodeID MaskPrefixLength 24, got %d", deserializedSourceIPAddress.MaskPrefixLength)
	}

	deserializedIPv4Address := deserializedSourceIPAddress.IPv4Address
	if len(deserializedIPv4Address) != 4 {
		t.Errorf("Expected IPv4 address length 4, got %d", len(deserializedIPv4Address))
	}

	expectedIPv4Address := []byte{2, 2, 3, 1}
	for i := range deserializedIPv4Address {
		if deserializedIPv4Address[i] != expectedIPv4Address[i] {
			t.Errorf("Expected IPv4 address %v, got %v", expectedIPv4Address, deserializedIPv4Address)
		}
	}
}
