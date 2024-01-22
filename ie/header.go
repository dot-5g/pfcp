package ie

import (
	"bytes"
	"encoding/binary"
)

const HeaderLength = 4

type Header struct {
	Type   IEType
	Length uint16
}

func (ieHeader *Header) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(ieHeader.Type))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, ieHeader.Length)

	return buf.Bytes()
}
