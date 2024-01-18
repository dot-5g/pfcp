package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type PDR struct {
	Header     Header
	PDRID      PDRID
	Precedence Precedence
	PDI        PDI
}

func NewPDR(pdrID PDRID, precedence Precedence, pdi PDI) (PDR, error) {
	ieHeader := Header{
		Type:   IEType(PDRIEType),
		Length: pdrID.Header.Length + precedence.Header.Length + pdi.Header.Length + 12,
	}

	return PDR{
		Header:     ieHeader,
		PDRID:      pdrID,
		Precedence: precedence,
		PDI:        pdi,
	}, nil
}

func (pdr PDR) IsZeroValue() bool {
	return pdr.Header.Length == 0
}

func (pdr PDR) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 4: Header
	buf.Write(pdr.Header.Serialize())

	// Octets 5 to n: PDR ID
	serializedPDRID := pdr.PDRID.Serialize()
	buf.Write(serializedPDRID)

	// Octets n+1 to m: Precedence
	serializedPrecedence := pdr.Precedence.Serialize()
	buf.Write(serializedPrecedence)

	// Octets m+1 to o: PDI
	serializedPDI := pdr.PDI.Serialize()
	buf.Write(serializedPDI)

	return buf.Bytes()

}

func DeserializePDR(ieHeader Header, value []byte) (PDR, error) {
	pdr := PDR{
		Header:     ieHeader,
		PDRID:      PDRID{},
		Precedence: Precedence{},
		PDI:        PDI{},
	}

	index := 0
	for index < len(value) {
		if index+4 > len(value) {
			return PDR{}, fmt.Errorf("slice bounds out of range")
		}

		currentIEType := binary.BigEndian.Uint16(value[index : index+2])
		currentIELength := binary.BigEndian.Uint16(value[index+2 : index+4])

		if index+4+int(currentIELength) > len(value) {
			return PDR{}, fmt.Errorf("slice bounds out of range")
		}

		currentIEValue := value[index+4 : index+4+int(currentIELength)]

		switch IEType(currentIEType) {
		case PDRIDIEType:
			pdrIDHeader := Header{
				Type:   IEType(currentIEType),
				Length: currentIELength,
			}
			pdrID, err := DeserializePDRID(pdrIDHeader, currentIEValue)
			if err != nil {
				return PDR{}, fmt.Errorf("failed to deserialize PDR ID: %v", err)
			}
			pdr.PDRID = pdrID
		case PrecedenceIEType:
			precedenceHeader := Header{
				Type:   IEType(currentIEType),
				Length: currentIELength,
			}
			precedence, err := DeserializePrecedence(precedenceHeader, currentIEValue)
			if err != nil {
				return PDR{}, fmt.Errorf("failed to deserialize Precedence: %v", err)
			}
			pdr.Precedence = precedence
		case PDIIEType:
			pdiIEHEader := Header{
				Type:   IEType(currentIEType),
				Length: currentIELength,
			}
			pdi, err := DeserializePDI(pdiIEHEader, currentIEValue)
			if err != nil {
				return PDR{}, fmt.Errorf("failed to deserialize PDI: %v", err)
			}
			pdr.PDI = pdi
		}

		index += 4 + int(currentIELength)
	}

	return pdr, nil
}
