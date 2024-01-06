package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type CreatePDR struct {
	IEType     uint16
	Length     uint16
	PDRID      PDRID
	Precedence Precedence
	PDI        PDI
}

func NewCreatePDR(pdrID PDRID, precedence Precedence, pdi PDI) (CreatePDR, error) {
	return CreatePDR{
		IEType:     uint16(CreatePDRIEType),
		Length:     pdrID.Length + precedence.Length + pdi.Length + 12,
		PDRID:      pdrID,
		Precedence: precedence,
		PDI:        pdi,
	}, nil
}

func (createPDR CreatePDR) IsZeroValue() bool {
	return createPDR.Length == 0
}

func (createPDR CreatePDR) Serialize() []byte {

	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(createPDR.IEType))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(createPDR.Length))

	// Octets 5 to n: PDR ID
	serializedPDRID := createPDR.PDRID.Serialize()
	buf.Write(serializedPDRID)

	// Octets n+1 to m: Precedence
	serializedPrecedence := createPDR.Precedence.Serialize()
	buf.Write(serializedPrecedence)

	// Octets m+1 to o: PDI
	serializedPDI := createPDR.PDI.Serialize()
	buf.Write(serializedPDI)

	return buf.Bytes()

}

func DeserializeCreatePDR(ieType uint16, length uint16, value []byte) (CreatePDR, error) {
	createPDR := CreatePDR{
		IEType:     ieType,
		Length:     length,
		PDRID:      PDRID{},
		Precedence: Precedence{},
		PDI:        PDI{},
	}

	index := 0
	for index < len(value) {
		if index+4 > len(value) {
			return CreatePDR{}, fmt.Errorf("slice bounds out of range")
		}

		currentIEType := binary.BigEndian.Uint16(value[index : index+2])
		currentIELength := binary.BigEndian.Uint16(value[index+2 : index+4])

		if index+4+int(currentIELength) > len(value) {
			return CreatePDR{}, fmt.Errorf("slice bounds out of range")
		}

		currentIEValue := value[index+4 : index+4+int(currentIELength)]

		switch IEType(currentIEType) {
		case PDRIDIEType:
			pdrID, err := DeserializePDRID(currentIEType, currentIELength, currentIEValue)
			if err != nil {
				return CreatePDR{}, fmt.Errorf("failed to deserialize PDR ID: %v", err)
			}
			createPDR.PDRID = pdrID
		case PrecedenceIEType:
			precedence, err := DeserializePrecedence(currentIEType, currentIELength, currentIEValue)
			if err != nil {
				return CreatePDR{}, fmt.Errorf("failed to deserialize Precedence: %v", err)
			}
			createPDR.Precedence = precedence
		case PDIIEType:
			pdi, err := DeserializePDI(currentIEType, currentIELength, currentIEValue)
			if err != nil {
				return CreatePDR{}, fmt.Errorf("failed to deserialize PDI: %v", err)
			}
			createPDR.PDI = pdi
		}

		index += 4 + int(currentIELength)
	}

	return createPDR, nil
}
