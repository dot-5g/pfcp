package ie

import (
	"bytes"
	"fmt"
)

type SourceInterface struct {
	Header Header
	Value  int
}

func NewSourceInterface(value int) (SourceInterface, error) {
	if value < 0 || value > 15 {
		return SourceInterface{}, fmt.Errorf("invalid value for SourceInterface: got %d, want 0-15", value)
	}

	ieHeader := Header{
		Type:   SourceInterfaceIEType,
		Length: 1,
	}

	return SourceInterface{
		Header: ieHeader,
		Value:  value,
	}, nil
}

func (sourceInterface SourceInterface) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 4: Header
	buf.Write(sourceInterface.Header.Serialize())

	// Octet 5: Spare (4 bits), Interface Value (4 bits)
	spareAndValue := (0x00 << 4) | (sourceInterface.Value & 0x0F)
	buf.WriteByte(byte(spareAndValue))

	return buf.Bytes()
}

func (sourceInterface SourceInterface) IsZeroValue() bool {
	return sourceInterface.Header.Length == 0
}

func (sourceInterface SourceInterface) SetHeader(header Header) InformationElement {
	sourceInterface.Header = header
	return sourceInterface
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
