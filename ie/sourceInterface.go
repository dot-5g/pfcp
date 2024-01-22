package ie

import (
	"bytes"
	"fmt"
)

type SourceInterface struct {
	Value int
}

func NewSourceInterface(value int) (SourceInterface, error) {
	if value < 0 || value > 15 {
		return SourceInterface{}, fmt.Errorf("invalid value for SourceInterface: got %d, want 0-15", value)
	}

	return SourceInterface{
		Value: value,
	}, nil
}

func (sourceInterface SourceInterface) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octet 5: Spare (4 bits), Interface Value (4 bits)
	spareAndValue := (0x00 << 4) | (sourceInterface.Value & 0x0F)
	buf.WriteByte(byte(spareAndValue))

	return buf.Bytes()
}

func (sourceInterface SourceInterface) GetType() IEType {
	return SourceInterfaceIEType
}

func DeserializeSourceInterface(ieValue []byte) (SourceInterface, error) {
	if len(ieValue) != 1 {
		return SourceInterface{}, fmt.Errorf("invalid length for PDRID: got %d bytes, want 2", len(ieValue))
	}
	value := int(ieValue[0] & 0x0F)

	return SourceInterface{
		Value: value,
	}, nil
}
