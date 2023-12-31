package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type SourceInterface struct {
	IEType uint16
	Length uint16
	Value  int
}

func NewSourceInterface(value int) (SourceInterface, error) {
	if value < 0 || value > 15 {
		return SourceInterface{}, fmt.Errorf("invalid value for SourceInterface: got %d, want 0-15", value)
	}

	return SourceInterface{
		IEType: uint16(SourceInterfaceIEType),
		Length: 1,
		Value:  value,
	}, nil
}

func (sourceInterface SourceInterface) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(sourceInterface.IEType))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(sourceInterface.Length))

	// Octet 5: Spare (4 bits), Interface Value (4 bits)
	spareAndValue := (0x00 << 4) | (sourceInterface.Value & 0x0F)
	buf.WriteByte(byte(spareAndValue))

	return buf.Bytes()
}

func (sourceInterface SourceInterface) IsZeroValue() bool {
	return sourceInterface.Length == 0
}

func DeserializeSourceInterface(ieType uint16, ieLength uint16, ieValue []byte) (SourceInterface, error) {
	if len(ieValue) != 1 {
		return SourceInterface{}, fmt.Errorf("invalid length for PDRID: got %d bytes, want 2", len(ieValue))
	}
	value := int(ieValue[0] & 0x0F)

	return SourceInterface{
		IEType: ieType,
		Length: ieLength,
		Value:  value,
	}, nil
}
