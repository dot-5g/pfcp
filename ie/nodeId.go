package ie

import (
	"bytes"
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
	Header Header
	Type   NodeIDType
	Value  []byte
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
	header := Header{
		Type:   NodeIDIEType,
		Length: length,
	}
	return NodeID{
		Header: header,
		Type:   nodeIDType,
		Value:  nodeIDValueBytes,
	}, nil
}

func (n NodeID) String() string {
	switch n.Type {
	case IPv4:
		return net.IP(n.Value).To4().String()
	case IPv6:
		return net.IP(n.Value).To16().String()
	case FQDN:
		return string(n.Value)
	default:
		return ""
	}
}

func (n NodeID) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 4: Header
	buf.Write(n.Header.Serialize())

	// Octet 5: Spare (4 bits) + Node ID Type (4 bits)
	spareAndType := byte(n.Type & 0x0F) // Ensure NodeIDType is only 4 bits
	buf.WriteByte(spareAndType)

	// Octets 6 to n+5: Node ID Value
	buf.Write(n.Value)

	return buf.Bytes()
}

func (n NodeID) IsZeroValue() bool {
	return n.Header.Length == 0
}

func (n NodeID) SetHeader(ieHeader Header) InformationElement {
	n.Header = ieHeader
	return n
}

func DeserializeNodeID(ieValue []byte) (NodeID, error) {

	if len(ieValue) < 1 {
		return NodeID{}, fmt.Errorf("invalid length for NodeID: got %d bytes, expected at least 1", len(ieValue))
	}

	nodeIDType := NodeIDType(ieValue[0] & 0x0F)
	if nodeIDType != IPv4 && nodeIDType != IPv6 && nodeIDType != FQDN {
		return NodeID{}, fmt.Errorf("invalid NodeIDType: %d", nodeIDType)
	}
	switch nodeIDType {
	case IPv4:
		if len(ieValue[1:]) != net.IPv4len {
			return NodeID{}, fmt.Errorf("invalid length for IPv4 NodeID: expected %d, got %d", net.IPv4len, len(ieValue[1:]))
		}
	case IPv6:
		if len(ieValue[1:]) != net.IPv6len {
			return NodeID{}, fmt.Errorf("invalid length for IPv6 NodeID: expected %d, got %d", net.IPv6len, len(ieValue[1:]))
		}
	}

	nodeID := NodeID{
		Type:  nodeIDType,
		Value: ieValue[1:],
	}

	return nodeID, nil
}

func (n NodeID) GetHeader() Header {
	return n.Header
}
