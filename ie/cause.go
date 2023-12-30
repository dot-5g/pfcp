package ie

import (
	"bytes"
	"encoding/binary"
)

type Cause struct {
	IEtype uint16
	Length uint16
	Value  uint8
}

func NewCause(value int) Cause {
	return Cause{
		IEtype: 19,
		Length: 1,
		Value:  uint8(value),
	}
}

func (cause Cause) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(cause.IEtype))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(cause.Length))

	// Octet 5: Value (1 byte)
	buf.WriteByte(cause.Value)

	return buf.Bytes()
}

func (cause Cause) IsZeroValue() bool {
	return cause.Length == 0
}

func DeserializeCause(ieType uint16, ieLength uint16, ieValue []byte) Cause {
	return Cause{
		IEtype: ieType,
		Length: ieLength,
		Value:  ieValue[0],
	}
}
