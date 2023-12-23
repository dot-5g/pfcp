package ie

import (
	"bytes"
	"encoding/binary"
)

type NodeID struct {
	Type        uint16
	Length      uint16
	NodeIDType  int
	NodeIDValue []byte
}

func NewNodeID(nodeIDType int, nodeIDValue []byte) NodeID {
	return NodeID{
		Type:        60,
		Length:      uint16(len(nodeIDValue) + 1),
		NodeIDType:  nodeIDType,
		NodeIDValue: nodeIDValue,
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
