package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type PDRID struct {
	IEType uint16
	Length uint16
	RuleID uint16
}

func NewPDRID(ruleID uint16) PDRID {
	return PDRID{
		IEType: uint16(PDRIDIEType),
		Length: 2,
		RuleID: ruleID,
	}
}

func (pdrID PDRID) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(pdrID.IEType))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(pdrID.Length))

	// Octets 5 to 6: RuleID
	binary.Write(buf, binary.BigEndian, pdrID.RuleID)

	return buf.Bytes()
}

func (pdrID PDRID) IsZeroValue() bool {
	return pdrID.Length == 0
}

func DeserializePDRID(ieType uint16, ieLength uint16, ieValue []byte) (PDRID, error) {
	if len(ieValue) != 2 {
		return PDRID{}, fmt.Errorf("invalid length for PDRID: got %d bytes, want 2", len(ieValue))
	}
	return PDRID{
		IEType: ieType,
		Length: ieLength,
		RuleID: binary.BigEndian.Uint16(ieValue),
	}, nil
}
