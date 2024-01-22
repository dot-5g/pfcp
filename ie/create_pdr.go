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

	for _, ie := range createPDR.GetIEs() {
		serializedIE := ie.Serialize()
		ieLength := uint16(len(serializedIE))
		ieHeader := Header{
			Type:   ie.GetType(),
			Length: ieLength,
		}
		buf.Write(ieHeader.Serialize())
		buf.Write(ie.Serialize())
	}

	return buf.Bytes()

}

func (createPDR CreatePDR) GetIEs() []InformationElement {
	return []InformationElement{createPDR.PDRID, createPDR.Precedence, createPDR.PDI}
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
