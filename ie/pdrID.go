package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type PDRID struct {
	RuleID uint16
}

func NewPDRID(ruleID uint16) (PDRID, error) {
	return PDRID{
		RuleID: ruleID,
	}, nil
}

func (pdrID PDRID) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 5 to 6: RuleID
	binary.Write(buf, binary.BigEndian, pdrID.RuleID)

	return buf.Bytes()
}

func (pdrID PDRID) GetType() IEType {
	return PDRIDIEType
}

func DeserializePDRID(ieValue []byte) (PDRID, error) {
	if len(ieValue) != 2 {
		return PDRID{}, fmt.Errorf("invalid length for PDRID: got %d bytes, want 2", len(ieValue))
	}
	return PDRID{
		RuleID: binary.BigEndian.Uint16(ieValue),
	}, nil
}
