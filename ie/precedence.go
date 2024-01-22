package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Precedence struct {
	Value uint32
}

func NewPrecedence(value uint32) (Precedence, error) {
	return Precedence{
		Value: value,
	}, nil
}

func (precedence Precedence) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 5 to 8: Value
	binary.Write(buf, binary.BigEndian, precedence.Value)

	return buf.Bytes()
}

func (precedence Precedence) GetType() IEType {
	return PrecedenceIEType
}

func DeserializePrecedence(ieValue []byte) (Precedence, error) {
	if len(ieValue) != 4 {
		return Precedence{}, fmt.Errorf("invalid length for Precedence: got %d bytes, want 4", len(ieValue))
	}
	return Precedence{
		Value: binary.BigEndian.Uint32(ieValue),
	}, nil
}
