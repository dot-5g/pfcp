package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type CreatePDR struct {
	PDRID      PDRID
	Precedence Precedence
	PDI        PDI
}

func NewCreatePDR(pdrID PDRID, precedence Precedence, pdi PDI) (CreatePDR, error) {
	return CreatePDR{
		PDRID:      pdrID,
		Precedence: precedence,
		PDI:        pdi,
	}, nil
}

func (createPDR CreatePDR) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 5 to n: PDR ID
	serializedPDRID := createPDR.PDRID.Serialize()
	pdrIDLength := uint16(len(serializedPDRID))
	pdrIDHeader := Header{
		Type:   createPDR.PDRID.GetType(),
		Length: pdrIDLength,
	}
	buf.Write(pdrIDHeader.Serialize())
	buf.Write(serializedPDRID)

	// Octets n+1 to m: Precedence
	serializedPrecedence := createPDR.Precedence.Serialize()
	precedenceLength := uint16(len(serializedPrecedence))
	precedenceHeader := Header{
		Type:   createPDR.Precedence.GetType(),
		Length: precedenceLength,
	}
	buf.Write(precedenceHeader.Serialize())
	buf.Write(serializedPrecedence)

	// Octets m+1 to o: PDI
	serializedPDI := createPDR.PDI.Serialize()
	pdiLength := uint16(len(serializedPDI))
	pdiHeader := Header{
		Type:   createPDR.PDI.GetType(),
		Length: pdiLength,
	}
	buf.Write(pdiHeader.Serialize())
	buf.Write(serializedPDI)

	return buf.Bytes()

}

func (createPDR CreatePDR) GetType() IEType {
	return CreatePDRIEType
}

func DeserializeCreatePDR(value []byte) (CreatePDR, error) {
	createPDR := CreatePDR{
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
			pdrID, err := DeserializePDRID(currentIEValue)
			if err != nil {
				return CreatePDR{}, fmt.Errorf("failed to deserialize PDR ID: %v", err)
			}

			createPDR.PDRID = pdrID
		case PrecedenceIEType:
			precedence, err := DeserializePrecedence(currentIEValue)
			if err != nil {
				return CreatePDR{}, fmt.Errorf("failed to deserialize Precedence: %v", err)
			}
			createPDR.Precedence = precedence
		case PDIIEType:
			pdi, err := DeserializePDI(currentIEValue)
			if err != nil {
				return CreatePDR{}, fmt.Errorf("failed to deserialize PDI: %v", err)
			}
			createPDR.PDI = pdi
		}

		index += 4 + int(currentIELength)
	}

	return createPDR, nil
}
