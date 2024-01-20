package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type CreatePDR struct {
	Header     Header
	PDRID      PDRID
	Precedence Precedence
	PDI        PDI
}

func NewCreatePDR(pdrID PDRID, precedence Precedence, pdi PDI) (CreatePDR, error) {
	ieHeader := Header{
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

func (createPDR CreatePDR) SetHeader(header Header) InformationElement {
	createPDR.Header = header
	return createPDR
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
			pdrIDHeader := Header{
				Type:   IEType(currentIEType),
				Length: currentIELength,
			}

			tempPdrID, err := DeserializePDRID(currentIEValue)
			if err != nil {
				return CreatePDR{}, fmt.Errorf("failed to deserialize PDR ID: %v", err)
			}
			pdrID, ok := tempPdrID.SetHeader(pdrIDHeader).(PDRID)
			if !ok {
				return CreatePDR{}, fmt.Errorf("type assertion to PDRID failed")
			}
			createPDR.PDRID = pdrID
		case PrecedenceIEType:
			precedenceHeader := Header{
				Type:   IEType(currentIEType),
				Length: currentIELength,
			}
			tempPrecedence, err := DeserializePrecedence(currentIEValue)
			if err != nil {
				return CreatePDR{}, fmt.Errorf("failed to deserialize Precedence: %v", err)
			}
			precedence, ok := tempPrecedence.SetHeader(precedenceHeader).(Precedence)
			if !ok {
				return CreatePDR{}, fmt.Errorf("type assertion to Precedence failed")
			}
			createPDR.Precedence = precedence
		case PDIIEType:
			pdiIEHeader := Header{
				Type:   IEType(currentIEType),
				Length: currentIELength,
			}
			tempPDI, err := DeserializePDI(currentIEValue)
			if err != nil {
				return CreatePDR{}, fmt.Errorf("failed to deserialize PDI: %v", err)
			}
			pdi, ok := tempPDI.SetHeader(pdiIEHeader).(PDI)
			if !ok {
				return CreatePDR{}, fmt.Errorf("type assertion to PDI failed")
			}
			createPDR.PDI = pdi
		}

		index += 4 + int(currentIELength)
	}

	return createPDR, nil
}
