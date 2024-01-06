package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Cause struct {
	IEType uint16
	Length uint16
	Value  uint8
}

func NewCause(value int) (Cause, error) {
	return Cause{
		IEType: uint16(CauseIEType),
		Length: 1,
		Value:  uint8(value),
	}, nil
}

func (cause Cause) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(cause.IEType))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(cause.Length))

	// Octet 5: Value (1 byte)
	buf.WriteByte(cause.Value)

	return buf.Bytes()
}

func (cause Cause) IsZeroValue() bool {
	return cause.Length == 0
}

func DeserializeCause(ieType uint16, ieLength uint16, ieValue []byte) (Cause, error) {
	var cause Cause

	if len(ieValue) != 1 {
		return cause, fmt.Errorf("invalid length for Cause: got %d bytes, want 1", len(ieValue))
	}

	if ieType != uint16(CauseIEType) {
		return cause, fmt.Errorf("invalid IE type: expected %d, got %d", CauseIEType, ieType)
	}

	if ieLength != 1 {
		return cause, fmt.Errorf("invalid length field for Cause: expected 1, got %d", ieLength)
	}

	return Cause{
		IEType: ieType,
		Length: ieLength,
		Value:  ieValue[0],
	}, nil
}
