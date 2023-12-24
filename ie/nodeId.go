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
	Type        uint16
	Length      uint16
	NodeIDType  NodeIDType
	NodeIDValue []byte
}

func NewNodeID(nodeIDType NodeIDType, nodeIDValue string) NodeID {
	var nodeIDValueBytes []byte
	var length uint16

	switch nodeIDType {
	case IPv4:
		ip := net.ParseIP(nodeIDValue)
		if ip == nil {
			panic("invalid IPv4 address")
		}
		ipv4 := ip.To4()
		if ipv4 == nil {
			panic("invalid IPv4 address")
		}
		nodeIDValueBytes = ipv4
		length = uint16(len(nodeIDValueBytes))
	case IPv6:
		ip := net.ParseIP(nodeIDValue)
		if ip == nil {
			panic("invalid IPv6 address")
		}
		ipv6 := ip.To16()
		if ipv6 == nil {
			panic("invalid IPv6 address")
		}
		nodeIDValueBytes = ipv6
		length = uint16(len(nodeIDValueBytes))
	case FQDN:
		fqdn := []byte(nodeIDValue)
		if len(fqdn) > 255 {
			panic("FQDN too long")
		}
		nodeIDValueBytes = fqdn
		length = uint16(len(nodeIDValueBytes))

	default:
		panic(fmt.Sprintf("invalid NodeIDType %d", nodeIDType))
	}
	return NodeID{
		Type:        60,
		Length:      length,
		NodeIDType:  nodeIDType,
		NodeIDValue: nodeIDValueBytes,
	}
}

func (n NodeID) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type (60)
	binary.Write(buf, binary.BigEndian, uint16(n.Type))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(n.Length))

	// Octet 5: Spare (4 bits) + Node ID Type (4 bits)
	spareAndType := byte(n.NodeIDType & 0x0F) // Ensure NodeIDType is only 4 bits
	buf.WriteByte(spareAndType)

	// Octets 6 to n+5: Node ID Value
	buf.Write(n.NodeIDValue)

	return buf.Bytes()
}
