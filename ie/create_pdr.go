package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type CreatePDR struct {
	Header     IEHeader
	PDRID      PDRID
	Precedence Precedence
	PDI        PDI
}

func NewCreatePDR(pdrID PDRID, precedence Precedence, pdi PDI) (CreatePDR, error) {
	ieHeader := IEHeader{
		Type:   IEType(CreatePDRIEType),
		Length: pdrID.Header.Length + precedence.Header.Length + pdi.Header.Length + 12,
	}

	return CreatePDR{
		Header:     ieHeader,
		PDRID:      pdrID,
		Precedence: precedence,
		PDI:        pdi,
	}, nil
}

func (createPDR CreatePDR) IsZeroValue() bool {
	return createPDR.Header.Length == 0
}

func (createPDR CreatePDR) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 4: Header
	buf.Write(createPDR.Header.Serialize())

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

func DeserializeCreatePDR(ieHeader IEHeader, value []byte) (CreatePDR, error) {
	createPDR := CreatePDR{
		Header:     ieHeader,
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
			pdrIDIEHeader := IEHeader{
				Type:   IEType(currentIEType),
				Length: currentIELength,
			}
			pdrID, err := DeserializePDRID(pdrIDIEHeader, currentIEValue)
			if err != nil {
				return CreatePDR{}, fmt.Errorf("failed to deserialize PDR ID: %v", err)
			}
			createPDR.PDRID = pdrID
		case PrecedenceIEType:
			precedenceIEHeader := IEHeader{
				Type:   IEType(currentIEType),
				Length: currentIELength,
			}
			precedence, err := DeserializePrecedence(precedenceIEHeader, currentIEValue)
			if err != nil {
				return CreatePDR{}, fmt.Errorf("failed to deserialize Precedence: %v", err)
			}
			createPDR.Precedence = precedence
		case PDIIEType:
			pdiIEHEader := IEHeader{
				Type:   IEType(currentIEType),
				Length: currentIELength,
			}
			pdi, err := DeserializePDI(pdiIEHEader, currentIEValue)
			if err != nil {
				return CreatePDR{}, fmt.Errorf("failed to deserialize PDI: %v", err)
			}
			createPDR.PDI = pdi
		}

		index += 4 + int(currentIELength)
	}

	return createPDR, nil
}
