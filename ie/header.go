package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

func DeserializeHeader(ieValue []byte) (Header, error) {
	if len(ieValue) != HeaderLength {
		return Header{}, fmt.Errorf("invalid length for Header: got %d bytes, want %d", len(ieValue), HeaderLength)
	}
	return Header{
		Type:   IEType(binary.BigEndian.Uint16(ieValue[0:2])),
		Length: binary.BigEndian.Uint16(ieValue[2:4]),
	}, nil
}
