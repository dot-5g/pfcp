package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type NodeReportType struct {
	IEtype uint16
	Length uint16
	GPQR   bool
	CKDR   bool
	UPRR   bool
	UPFR   bool
}

func NewNodeReportType(gpqr bool, ckdr bool, uprr bool, upfr bool) (NodeReportType, error) {
	return NodeReportType{
		IEtype: uint16(NodeReportTypeIEType),
		Length: 1,
		GPQR:   gpqr,
		CKDR:   ckdr,
		UPRR:   uprr,
		UPFR:   upfr,
	}, nil
}

func (nrt NodeReportType) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(nrt.IEtype))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(nrt.Length))

	// Octet 5: Spare, Spare, Spare, Spare, GPQR, CKDR, UPRR, UPFR
	var octet5 byte
	if nrt.GPQR {
		octet5 |= 1 << 3
	}
	if nrt.CKDR {
		octet5 |= 1 << 2
	}
	if nrt.UPRR {
		octet5 |= 1 << 1
	}
	if nrt.UPFR {
		octet5 |= 1
	}
	buf.WriteByte(octet5)

	return buf.Bytes()
}

func (nrt NodeReportType) IsZeroValue() bool {
	return nrt.Length == 0
}

func DeserializeNodeReportType(ieType uint16, ieLength uint16, ieValue []byte) (NodeReportType, error) {
	var nrt NodeReportType

	if len(ieValue) < 1 {
		return nrt, fmt.Errorf("invalid length for NodeReportType: got %d bytes, expected at least 1", len(ieValue))
	}

	if ieType != uint16(NodeReportTypeIEType) {
		return nrt, fmt.Errorf("invalid IE type: expected %d, got %d", NodeReportTypeIEType, ieType)
	}

	buf := bytes.NewBuffer(ieValue)

	nrt.IEtype = ieType
	nrt.Length = ieLength

	var octet5 byte
	var err error
	if octet5, err = buf.ReadByte(); err != nil {
		return nrt, fmt.Errorf("error reading NodeReportType flags: %v", err)
	}

	nrt.GPQR = octet5&0x08 != 0
	nrt.CKDR = octet5&0x04 != 0
	nrt.UPRR = octet5&0x02 != 0
	nrt.UPFR = octet5&0x01 != 0

	return nrt, nil
}
