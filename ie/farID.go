package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type FARID struct {
	Header Header
	Value  uint32
}

func NewFarID(value uint32) (FARID, error) {
	ieHeader := Header{
		Type:   IEType(FARIDIEType),
		Length: 4,
	}

	return FARID{
		Header: ieHeader,
		Value:  value,
	}, nil
}

func (farID FARID) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 4: Header
	buf.Write(farID.Header.Serialize())

	// Octets 5 to 8: Value
	binary.Write(buf, binary.BigEndian, farID.Value)

	return buf.Bytes()
}

func (farID FARID) IsZeroValue() bool {
	return farID.Header.Length == 0
}

func DeserializeFARID(ieHeader Header, ieValue []byte) (FARID, error) {
	if len(ieValue) != 4 {
		return FARID{}, fmt.Errorf("invalid length for FARID: got %d bytes, want 4", len(ieValue))
	}

	if ieHeader.Type != FARIDIEType {
		return FARID{}, fmt.Errorf("invalid IE type for FARID: got %d, want %d", ieHeader.Type, FARIDIEType)
	}

	return FARID{
		Header: ieHeader,
		Value:  binary.BigEndian.Uint32(ieValue),
	}, nil
}
