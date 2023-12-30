package ie

import (
	"bytes"
	"encoding/binary"
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

func NewNodeID(nodeID string) NodeID {
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
			panic("FQDN too long")
		}
		nodeIDValueBytes = fqdn
		length = uint16(len(nodeIDValueBytes)) + 1
		nodeIDType = FQDN
	}

	return NodeID{
		IEtype:      60,
		Length:      length,
		NodeIDType:  nodeIDType,
		NodeIDValue: nodeIDValueBytes,
	}
}

func (n NodeID) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type (60)
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

func DeserializeNodeID(ieType uint16, ieLength uint16, ieValue []byte) NodeID {
	nodeIDType := NodeIDType(ieValue[0] & 0x0F) // Ensure NodeIDType is only 4 bits
	nodeIDValue := ieValue[1:]

	return NodeID{
		IEtype:      ieType,
		Length:      ieLength,
		NodeIDType:  nodeIDType,
		NodeIDValue: nodeIDValue,
	}
}
