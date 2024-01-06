package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type FARID struct {
	IEType uint16
	Length uint16
	Value  uint32
}

func NewFarID(value uint32) FARID {
	return FARID{
		IEType: uint16(FARIDIEType),
		Length: 4,
		Value:  value,
	}
}

func (farID FARID) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(farID.IEType))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(farID.Length))

	// Octets 5 to 8: Value
	binary.Write(buf, binary.BigEndian, farID.Value)

	return buf.Bytes()
}

func (farID FARID) IsZeroValue() bool {
	return farID.Length == 0
}

func DeserializeFARID(ieType uint16, ieLength uint16, ieValue []byte) (FARID, error) {
	if len(ieValue) != 4 {
		return FARID{}, fmt.Errorf("invalid length for FARID: got %d bytes, want 4", len(ieValue))
	}

	if ieType != uint16(FARIDIEType) {
		return FARID{}, fmt.Errorf("invalid IE type for FARID: got %d, want %d", ieType, FARIDIEType)
	}

	return FARID{
		IEType: ieType,
		Length: ieLength,
		Value:  binary.BigEndian.Uint32(ieValue),
	}, nil
}
