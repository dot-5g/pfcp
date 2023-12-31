package ie

import (
	"bytes"
	"encoding/binary"
)

type NodeReportType struct {
	IEtype uint16
	Length uint16
	GPQR   bool
	CKDR   bool
	UPRR   bool
	UPFR   bool
}

func NewNodeReportType(gpqr bool, ckdr bool, uprr bool, upfr bool) NodeReportType {
	return NodeReportType{
		IEtype: uint16(NodeReportTypeIEType),
		Length: 1,
		GPQR:   gpqr,
		CKDR:   ckdr,
		UPRR:   uprr,
		UPFR:   upfr,
	}
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

func DeserializeNodeReportType(ieType uint16, ieLength uint16, ieValue []byte) NodeReportType {
	var nrt NodeReportType

	buf := bytes.NewBuffer(ieValue)

	nrt.IEtype = ieType
	nrt.Length = ieLength

	// Read the bit flags from octet 5
	var octet5 byte
	if len(buf.Bytes()) > 0 {
		octet5, _ = buf.ReadByte()
		nrt.GPQR = octet5&0x08 != 0
		nrt.CKDR = octet5&0x04 != 0
		nrt.UPRR = octet5&0x02 != 0
		nrt.UPFR = octet5&0x01 != 0
	}

	return nrt
}
