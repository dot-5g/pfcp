package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type PDRID struct {
	Header IEHeader
	RuleID uint16
}

func NewPDRID(ruleID uint16) (PDRID, error) {
	ieHeader := IEHeader{
		Type:   PDRIDIEType,
		Length: 2,
	}

	return PDRID{
		Header: ieHeader,
		RuleID: ruleID,
	}, nil
}

func (pdrID PDRID) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 4: Header
	buf.Write(pdrID.Header.Serialize())

	// Octets 5 to 6: RuleID
	binary.Write(buf, binary.BigEndian, pdrID.RuleID)

	return buf.Bytes()
}

func (pdrID PDRID) IsZeroValue() bool {
	return pdrID.Header.Length == 0
}

func DeserializePDRID(ieHeader IEHeader, ieValue []byte) (PDRID, error) {
	if len(ieValue) != 2 {
		return PDRID{}, fmt.Errorf("invalid length for PDRID: got %d bytes, want 2", len(ieValue))
	}
	return PDRID{
		Header: ieHeader,
		RuleID: binary.BigEndian.Uint16(ieValue),
	}, nil
}
