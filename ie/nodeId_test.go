package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestNewNodeIDIPv4(t *testing.T) {
	nodeID := ie.NewNodeID(ie.IPv4, "1.2.3.4")

	if nodeID.Type != 60 {
		t.Errorf("Expected NodeID, got %d", nodeID.Type)
	}

	if nodeID.Length != 4 {
		t.Errorf("Expected NodeID length 4, got %d", nodeID.Length)
	}

	if nodeID.NodeIDType != 0 {
		t.Errorf("Expected NodeID type IPv4, got %d", nodeID.NodeIDType)
	}

	if len(nodeID.NodeIDValue) != 4 {
		t.Errorf("Expected NodeID value length 4, got %d", len(nodeID.NodeIDValue))
	}

	expectedNodeIDValue := []byte{1, 2, 3, 4}
	for i := range nodeID.NodeIDValue {
		if nodeID.NodeIDValue[i] != expectedNodeIDValue[i] {
			t.Errorf("Expected NodeID value %v, got %v", expectedNodeIDValue, nodeID.NodeIDValue)
		}
	}

}

func TestNewNodeIDIPv6(t *testing.T) {
	nodeID := ie.NewNodeID(ie.IPv6, "2001:db8::68")

	if nodeID.Type != 60 {
		t.Errorf("Expected NodeID, got %d", nodeID.Type)
	}

	if nodeID.Length != 16 {
		t.Errorf("Expected NodeID length 16, got %d", nodeID.Length)
	}

	if nodeID.NodeIDType != 1 {
		t.Errorf("Expected NodeID type IPv6, got %d", nodeID.NodeIDType)
	}

	if len(nodeID.NodeIDValue) != 16 {
		t.Errorf("Expected NodeID value length 16, got %d", len(nodeID.NodeIDValue))
	}

	expectedNodeIDValue := []byte{32, 1, 13, 184, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 104}
	for i := range nodeID.NodeIDValue {
		if nodeID.NodeIDValue[i] != expectedNodeIDValue[i] {
			t.Errorf("Expected NodeID value %v, got %v", expectedNodeIDValue, nodeID.NodeIDValue)
		}
	}

}

func TestNewNodeIDFQDN(t *testing.T) {
	nodeID := ie.NewNodeID(ie.FQDN, "www.example.com")

	if nodeID.Type != 60 {
		t.Errorf("Expected NodeID, got %d", nodeID.Type)
	}

	if nodeID.Length != 15 {
		t.Errorf("Expected NodeID length 15, got %d", nodeID.Length)
	}

	if nodeID.NodeIDType != 2 {
		t.Errorf("Expected NodeID type FQDN, got %d", nodeID.NodeIDType)
	}

	if len(nodeID.NodeIDValue) != 15 {
		t.Errorf("Expected NodeID value length 15, got %d", len(nodeID.NodeIDValue))
	}

	expectedNodeIDValue := []byte{119, 119, 119, 46, 101, 120, 97, 109, 112, 108, 101, 46, 99, 111, 109}
	for i := range nodeID.NodeIDValue {
		if nodeID.NodeIDValue[i] != expectedNodeIDValue[i] {
			t.Errorf("Expected NodeID value %v, got %v", expectedNodeIDValue, nodeID.NodeIDValue)
		}
	}

}
