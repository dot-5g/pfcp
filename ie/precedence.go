package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Precedence struct {
	Header Header
	Value  uint32
}

func NewPrecedence(value uint32) (Precedence, error) {
	ieHeader := Header{
		Type:   PrecedenceIEType,
		Length: 4,
	}

	return Precedence{
		Header: ieHeader,
		Value:  value,
	}, nil
}

func (precedence Precedence) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 4: Header
	buf.Write(precedence.Header.Serialize())

	// Octets 5 to 8: Value
	binary.Write(buf, binary.BigEndian, precedence.Value)

	return buf.Bytes()
}

func (precedence Precedence) IsZeroValue() bool {
	return precedence.Header.Length == 0
}

func DeserializePrecedence(ieHeader Header, ieValue []byte) (Precedence, error) {
	if len(ieValue) != 4 {
		return Precedence{}, fmt.Errorf("invalid length for Precedence: got %d bytes, want 4", len(ieValue))
	}
	return Precedence{
		Header: ieHeader,
		Value:  binary.BigEndian.Uint32(ieValue),
	}, nil
}
