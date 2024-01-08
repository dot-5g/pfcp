package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const HeaderLength = 4

type IEHeader struct {
	Type   IEType
	Length uint16
}

func (ieHeader *IEHeader) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(ieHeader.Type))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(ieHeader.Length))

	return buf.Bytes()
}

func DeserializeIEHeader(payload []byte) (IEHeader, error) {
	var ieHeader IEHeader

	if len(payload) < HeaderLength {
		return ieHeader, fmt.Errorf("not enough bytes for IE header")
	}

	ieHeader.Type = IEType(binary.BigEndian.Uint16(payload[:2]))
	ieHeader.Length = binary.BigEndian.Uint16(payload[2:4])

	return ieHeader, nil
}
