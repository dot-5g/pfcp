package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestNewNodeIDIPv4(t *testing.T) {
	nodeID, err := ie.NewNodeID("1.2.3.4")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if nodeID.Header.Type != 60 {
		t.Errorf("Expected NodeID, got %d", nodeID.Header.Type)
	}

	if nodeID.Header.Length != 4+1 {
		t.Errorf("Expected NodeID length 4, got %d", nodeID.Header.Length)
	}

	if nodeID.Type != 0 {
		t.Errorf("Expected NodeID type IPv4, got %d", nodeID.Type)
	}

	if len(nodeID.Value) != 4 {
		t.Errorf("Expected NodeID value length 4, got %d", len(nodeID.Value))
	}

	expectedNodeIDValue := []byte{1, 2, 3, 4}
	for i := range nodeID.Value {
		if nodeID.Value[i] != expectedNodeIDValue[i] {
			t.Errorf("Expected NodeID value %v, got %v", expectedNodeIDValue, nodeID.Value)
		}
	}

}

func TestString(t *testing.T) {
	nodeIDstring := "1.2.3.4"
	nodeID, err := ie.NewNodeID(nodeIDstring)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if nodeID.String() != nodeIDstring {
		t.Errorf("Expected NodeID string, got %s", nodeID.String())
	}

}

func TestNewNodeIDIPv6(t *testing.T) {
	nodeID, err := ie.NewNodeID("2001:db8::68")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if nodeID.Header.Type != 60 {
		t.Errorf("Expected NodeID, got %d", nodeID.Header.Type)
	}

	if nodeID.Header.Length != 16+1 {
		t.Errorf("Expected NodeID length 16, got %d", nodeID.Header.Length)
	}

	if nodeID.Type != 1 {
		t.Errorf("Expected NodeID type IPv6, got %d", nodeID.Type)
	}

	if len(nodeID.Value) != 16 {
		t.Errorf("Expected NodeID value length 16, got %d", len(nodeID.Value))
	}

	expectedNodeIDValue := []byte{32, 1, 13, 184, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 104}
	for i := range nodeID.Value {
		if nodeID.Value[i] != expectedNodeIDValue[i] {
			t.Errorf("Expected NodeID value %v, got %v", expectedNodeIDValue, nodeID.Value)
		}
	}

}

func TestNewNodeIDFQDN(t *testing.T) {
	nodeID, err := ie.NewNodeID("www.example.com")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if nodeID.Header.Type != 60 {
		t.Errorf("Expected NodeID, got %d", nodeID.Header.Type)
	}

	if nodeID.Header.Length != 15+1 {
		t.Errorf("Expected NodeID length 15, got %d", nodeID.Header.Length)
	}

	if nodeID.Type != 2 {
		t.Errorf("Expected NodeID type FQDN, got %d", nodeID.Type)
	}

	if len(nodeID.Value) != 15 {
		t.Errorf("Expected NodeID value length 15, got %d", len(nodeID.Value))
	}

	expectedNodeIDValue := []byte{119, 119, 119, 46, 101, 120, 97, 109, 112, 108, 101, 46, 99, 111, 109}
	for i := range nodeID.Value {
		if nodeID.Value[i] != expectedNodeIDValue[i] {
			t.Errorf("Expected NodeID value %v, got %v", expectedNodeIDValue, nodeID.Value)
		}
	}

}

func TestGivenSerializedWhenDeserializNodeIDThenFieldsSetCorrectly(t *testing.T) {
	nodeID, err := ie.NewNodeID("1.2.3.4")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	serializedNodeID := nodeID.Serialize()

	deserializedNodeID, err := ie.DeserializeNodeID(serializedNodeID[4:])

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deserializedNodeID.Type != 0 {
		t.Errorf("Expected NodeID type FQDN, got %d", deserializedNodeID.Type)
	}

	if len(deserializedNodeID.Value) != 4 {
		t.Errorf("Expected NodeID value length 4, got %d", len(deserializedNodeID.Value))
	}

}
