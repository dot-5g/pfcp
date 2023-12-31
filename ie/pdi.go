package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type PDI struct {
	IEType          uint16
	Length          uint16
	SourceInterface SourceInterface
}

func NewPDI(sourceInterface SourceInterface) PDI {
	return PDI{
		IEType:          uint16(PDIIEType),
		Length:          sourceInterface.Length + 4,
		SourceInterface: sourceInterface,
	}
}

func (pdi PDI) Serialize() []byte {

	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(pdi.IEType))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(pdi.Length))

	// Octets 5 to n: Source Interface
	serializedSourceInterface := pdi.SourceInterface.Serialize()
	buf.Write(serializedSourceInterface)

	return buf.Bytes()

}

func (pdi PDI) IsZeroValue() bool {
	return pdi.Length == 0
}

func DeserializePDI(ieType uint16, ieLength uint16, ieValue []byte) (PDI, error) {
	if len(ieValue) < 1 {
		return PDI{}, fmt.Errorf("invalid length for PDI: got %d bytes, want at least 1", len(ieValue))
	}

	sourceInterfaceIELength := ieLength - IEHeaderLength
	sourceInterfaceIEValue := ieValue[4 : 4+sourceInterfaceIELength]
	sourceInterfaceIEType := binary.BigEndian.Uint16(ieValue[:2])

	sourceInterface, err := DeserializeSourceInterface(sourceInterfaceIEType, sourceInterfaceIELength, sourceInterfaceIEValue)
	if err != nil {
		return PDI{}, err
	}

	return PDI{
		IEType:          ieType,
		Length:          ieLength,
		SourceInterface: sourceInterface,
	}, nil
}
