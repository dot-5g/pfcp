package ie_test

import (
	"net"
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenValidIPv4AddressWhenNewFSEIDThenFieldsAreSetCorrectly(t *testing.T) {
	seid := uint64(0x1234567890ABCDEF)

	fseid, err := ie.NewFSEID(seid, "1.2.3.4", "")

	if err != nil {
		t.Fatalf("Error creating FSEID: %v", err)
	}

	if fseid.Header.Type != 57 {
		t.Errorf("Expected FSEID IEType 97, got %d", fseid.Header.Type)
	}

	if fseid.Header.Length != 13 {
		t.Errorf("Expected FSEID length 13, got %d", fseid.Header.Length)
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

func TestGivenValidIPv6AddressWhenNewFSEIDThenFieldsAreSetCorrectly(t *testing.T) {
	seid := uint64(0x1234567890ABCDEF)

	fseid, err := ie.NewFSEID(seid, "", "2001:db8::68")

	if err != nil {
		t.Fatalf("Error creating FSEID: %v", err)
	}

	if fseid.Header.Type != 57 {
		t.Errorf("Expected FSEID IEType 97, got %d", fseid.Header.Type)
	}

	if fseid.Header.Length != 25 {
		t.Errorf("Expected FSEID length 25, got %d", fseid.Header.Length)
	}

	if fseid.V4 != false {
		t.Errorf("Expected FSEID V4 false, got %v", fseid.V4)
	}

	if fseid.V6 != true {
		t.Errorf("Expected FSEID V6 true, got %v", fseid.V6)
	}

	if fseid.SEID != seid {
		t.Errorf("Expected FSEID SEID %d, got %d", seid, fseid.SEID)
	}

	expectedIPv6 := []byte{32, 1, 13, 184, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 104}
	for i := range fseid.IPv6 {
		if fseid.IPv6[i] != expectedIPv6[i] {
			t.Errorf("Expected FSEID IPv6 %v, got %v", expectedIPv6, fseid.IPv6)
		}
	}

}

func TestGivenIPv4AndIPv6AddressWhenNewFSEIDThenFieldsAreSetCorrectly(t *testing.T) {
	seid := uint64(0x1234567890ABCDEF)

	fseid, err := ie.NewFSEID(seid, "1.2.3.4", "2001:db8::68")

	if err != nil {
		t.Fatalf("Error creating FSEID: %v", err)
	}

	if fseid.Header.Type != 57 {
		t.Errorf("Expected FSEID IEType 57, got %d", fseid.Header.Type)
	}

	if fseid.Header.Length != 29 {
		t.Errorf("Expected FSEID length 29, got %d", fseid.Header.Length)
	}

	if fseid.V4 != true {
		t.Errorf("Expected FSEID V4 true, got %v", fseid.V4)
	}

	if fseid.V6 != true {
		t.Errorf("Expected FSEID V6 true, got %v", fseid.V6)
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

	expectedIPv6 := []byte{32, 1, 13, 184, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 104}
	for i := range fseid.IPv6 {
		if fseid.IPv6[i] != expectedIPv6[i] {
			t.Errorf("Expected FSEID IPv6 %v, got %v", expectedIPv6, fseid.IPv6)
		}
	}

}

func TestGivenIPv4SerializedWhenDeserializeThenFieldsSetCorrectly(t *testing.T) {
	seid := uint64(0x1234567890ABCDEF)
	ipv4 := "2.3.4.5"
	ipv6 := ""

	fseid, err := ie.NewFSEID(seid, ipv4, ipv6)

	if err != nil {
		t.Fatalf("Error creating FSEID: %v", err)
	}

	serialized := fseid.Serialize()

	ieHeader := ie.Header{
		Type:   57,
		Length: 13,
	}

	deserialized, err := ie.DeserializeFSEID(ieHeader, serialized[4:])

	if err != nil {
		t.Fatalf("Error deserializing FSEID: %v", err)
	}

	if deserialized.Header.Type != 57 {
		t.Errorf("Expected FSEID IEType 57, got %d", deserialized.Header.Type)
	}

	if deserialized.Header.Length != 13 {
		t.Errorf("Expected FSEID length 13, got %d", deserialized.Header.Length)
	}

	if deserialized.V4 != true {
		t.Errorf("Expected FSEID V4 true, got %v", deserialized.V4)
	}

	if deserialized.V6 != false {
		t.Errorf("Expected FSEID V6 false, got %v", deserialized.V6)
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

	ipv6Net := net.ParseIP(ipv6)
	expectedIPv6 := ipv6Net.To16()

	for i := range deserialized.IPv6 {
		if deserialized.IPv6[i] != expectedIPv6[i] {
			t.Errorf("Expected FSEID IPv6 %v, got %v", expectedIPv6, deserialized.IPv6)
		}
	}

}
