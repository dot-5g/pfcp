package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type FARID struct {
	Value uint32
}

func NewFarID(value uint32) (FARID, error) {
	return FARID{
		Value: value,
	}, nil
}

func (farID FARID) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 5 to 8: Value
	binary.Write(buf, binary.BigEndian, farID.Value)

	return buf.Bytes()
}

func (farID FARID) GetType() IEType {
	return FARIDIEType
}

func DeserializeFARID(ieValue []byte) (FARID, error) {
	if len(ieValue) != 4 {
		return FARID{}, fmt.Errorf("invalid length for FARID: got %d bytes, want 4", len(ieValue))
	}

	return FARID{
		Value: binary.BigEndian.Uint32(ieValue),
	}, nil
}
