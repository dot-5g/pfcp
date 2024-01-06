package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

const (
	IPv4 NodeIDType = iota
	IPv6
	FQDN
)

type NodeIDType int

type NodeID struct {
	IEtype      uint16
	Length      uint16
	NodeIDType  NodeIDType
	NodeIDValue []byte
}

func NewNodeID(nodeID string) (NodeID, error) {
	var nodeIDValueBytes []byte
	var length uint16
	var nodeIDType NodeIDType

	ip := net.ParseIP(nodeID)

	if ip.To4() != nil {
		nodeIDValueBytes = ip.To4()
		length = uint16(len(nodeIDValueBytes)) + 1
		nodeIDType = IPv4
	} else if ip.To16() != nil {
		nodeIDValueBytes = ip.To16()
		length = uint16(len(nodeIDValueBytes)) + 1
		nodeIDType = IPv6
	} else {
		fqdn := []byte(nodeID)
		if len(fqdn) > 255 {
			return NodeID{}, fmt.Errorf("invalid length for FQDN NodeID: got %d bytes, want <= 255", len(fqdn))
		}
		nodeIDValueBytes = fqdn
		length = uint16(len(nodeIDValueBytes)) + 1
		nodeIDType = FQDN
	}

	return NodeID{
		IEtype:      uint16(NodeIDIEType),
		Length:      length,
		NodeIDType:  nodeIDType,
		NodeIDValue: nodeIDValueBytes,
	}, nil
}

func (n NodeID) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(n.IEtype))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(n.Length))

	// Octet 5: Spare (4 bits) + Node ID Type (4 bits)
	spareAndType := byte(n.NodeIDType & 0x0F) // Ensure NodeIDType is only 4 bits
	buf.WriteByte(spareAndType)

	// Octets 6 to n+5: Node ID Value
	buf.Write(n.NodeIDValue)

	return buf.Bytes()
}

func (n NodeID) IsZeroValue() bool {
	return n.Length == 0
}

func DeserializeNodeID(ieType uint16, ieLength uint16, ieValue []byte) (NodeID, error) {
	var nodeID NodeID

	if len(ieValue) < 1 {
		return nodeID, fmt.Errorf("invalid length for NodeID: got %d bytes, expected at least 1", len(ieValue))
	}

	if ieType != uint16(NodeIDIEType) {
		return nodeID, fmt.Errorf("invalid IE type: expected %d, got %d", NodeIDIEType, ieType)
	}

	nodeIDType := NodeIDType(ieValue[0] & 0x0F)
	if nodeIDType != IPv4 && nodeIDType != IPv6 && nodeIDType != FQDN {
		return nodeID, fmt.Errorf("invalid NodeIDType: %d", nodeIDType)
	}

	switch nodeIDType {
	case IPv4:
		if len(ieValue[1:]) != net.IPv4len {
			return nodeID, fmt.Errorf("invalid length for IPv4 NodeID: expected %d, got %d", net.IPv4len, len(ieValue[1:]))
		}
	case IPv6:
		if len(ieValue[1:]) != net.IPv6len {
			return nodeID, fmt.Errorf("invalid length for IPv6 NodeID: expected %d, got %d", net.IPv6len, len(ieValue[1:]))
		}
	}

	nodeID = NodeID{
		IEtype:      ieType,
		Length:      ieLength,
		NodeIDType:  nodeIDType,
		NodeIDValue: ieValue[1:],
	}

	return nodeID, nil
}
