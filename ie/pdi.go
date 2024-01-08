package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type PDI struct {
	Header          IEHeader
	SourceInterface SourceInterface
}

func NewPDI(sourceInterface SourceInterface) (PDI, error) {
	ieHeader := IEHeader{
		Type:   PDIIEType,
		Length: sourceInterface.Header.Length + 4,
	}

	return PDI{
		Header:          ieHeader,
		SourceInterface: sourceInterface,
	}, nil
}

func (pdi PDI) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 4: Header
	buf.Write(pdi.Header.Serialize())

	// Octets 5 to n: Source Interface
	serializedSourceInterface := pdi.SourceInterface.Serialize()
	buf.Write(serializedSourceInterface)

	return buf.Bytes()

}

func (pdi PDI) IsZeroValue() bool {
	return pdi.Header.Length == 0
}

func DeserializePDI(ieHeader IEHeader, ieValue []byte) (PDI, error) {
	if len(ieValue) < 1 {
		return PDI{}, fmt.Errorf("invalid length for PDI: got %d bytes, want at least 1", len(ieValue))
	}

	sourceInterfaceIELength := ieHeader.Length - HeaderLength
	sourceInterfaceIEValue := ieValue[4 : 4+sourceInterfaceIELength]
	sourceInterfaceIEType := binary.BigEndian.Uint16(ieValue[:2])

	sourceInterfaceIEHeader := IEHeader{
		Type:   IEType(sourceInterfaceIEType),
		Length: sourceInterfaceIELength,
	}
	sourceInterface, err := DeserializeSourceInterface(sourceInterfaceIEHeader, sourceInterfaceIEValue)
	if err != nil {
		return PDI{}, err
	}

	return PDI{
		Header:          ieHeader,
		SourceInterface: sourceInterface,
	}, nil
}
