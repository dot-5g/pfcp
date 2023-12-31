package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenValidIPAddressWhenNewFSEIDThenFieldsAreSetCorrectly(t *testing.T) {
	seid := uint64(0x1234567890ABCDEF)

	fseid, err := ie.NewFSEID(seid, "1.2.3.4", "")

	if err != nil {
		t.Fatalf("Error creating FSEID: %v", err)
	}

	if fseid.IEType != 57 {
		t.Errorf("Expected FSEID IEType 97, got %d", fseid.IEType)
	}

	if fseid.Length != 13 {
		t.Errorf("Expected FSEID length 12, got %d", fseid.Length)
	}

	if fseid.V4 != true {
		t.Errorf("Expected FSEID V4 true, got %v", fseid.V4)
	}

	if fseid.V6 != false {
		t.Errorf("Expected FSEID V6 false, got %v", fseid.V6)
	}

	if fseid.SEID != seid {
		t.Errorf("Expected FSEID SEID %d, got %d", seid, fseid.SEID)
	}

	expectedIPv4 := []byte{1, 2, 3, 4}
	for i := range fseid.IPv4 {
		if fseid.IPv4[i] != expectedIPv4[i] {
			t.Errorf("Expected FSEID IPv4 %v, got %v", expectedIPv4, fseid.IPv4)
		}
	}

}

func TestGivenSerializedWhenDeserializeThenFieldsSetCorrectly(t *testing.T) {
	seid := uint64(0x1234567890ABCDEF)
	ipv4 := "2.3.4.5"
	ipv6 := "2001:db8::68"

	fseid, err := ie.NewFSEID(seid, ipv4, ipv6)

	if err != nil {
		t.Fatalf("Error creating FSEID: %v", err)
	}

	serialized := fseid.Serialize()

	deserialized, err := ie.DeserializeFSEID(57, 12, serialized[4:])

	if err != nil {
		t.Fatalf("Error deserializing FSEID: %v", err)
	}

	if deserialized.IEType != 57 {
		t.Errorf("Expected FSEID IEType 57, got %d", deserialized.IEType)
	}

	if deserialized.Length != 12 {
		t.Errorf("Expected FSEID length 12, got %d", deserialized.Length)
	}

	if deserialized.V4 != true {
		t.Errorf("Expected FSEID V4 true, got %v", deserialized.V4)
	}

	if deserialized.V6 != true {
		t.Errorf("Expected FSEID V6 true, got %v", deserialized.V6)
	}

	if deserialized.SEID != seid {
		t.Errorf("Expected FSEID SEID %d, got %d", seid, deserialized.SEID)
	}

	expectedIPv4 := []byte{2, 3, 4, 5}
	for i := range deserialized.IPv4 {
		if deserialized.IPv4[i] != expectedIPv4[i] {
			t.Errorf("Expected FSEID IPv4 %v, got %v", expectedIPv4, deserialized.IPv4)
		}
	}

	expectedIPv6 := []byte{32, 1, 13, 184, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x00, 0x00, 0x68}
	for i := range deserialized.IPv6 {
		if deserialized.IPv6[i] != expectedIPv6[i] {
			t.Errorf("Expected FSEID IPv6 %v, got %v", expectedIPv6, deserialized.IPv6)
		}
	}

}
