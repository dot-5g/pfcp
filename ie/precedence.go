package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Precedence struct {
	IEType uint16
	Length uint16
	Value  uint32
}

func NewPrecedence(value uint32) Precedence {
	return Precedence{
		IEType: uint16(PrecedenceIEType),
		Length: 4,
		Value:  value,
	}
}

func (precedence Precedence) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(precedence.IEType))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(precedence.Length))

	// Octets 5 to 8: Value
	binary.Write(buf, binary.BigEndian, precedence.Value)

	return buf.Bytes()
}

func (precedence Precedence) IsZeroValue() bool {
	return precedence.Length == 0
}

func DeserializePrecedence(ieType uint16, ieLength uint16, ieValue []byte) (Precedence, error) {
	if len(ieValue) != 4 {
		return Precedence{}, fmt.Errorf("invalid length for Precedence: got %d bytes, want 4", len(ieValue))
	}
	return Precedence{
		IEType: ieType,
		Length: ieLength,
		Value:  binary.BigEndian.Uint32(ieValue),
	}, nil
}
