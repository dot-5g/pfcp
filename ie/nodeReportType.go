package ie

import (
	"bytes"
	"fmt"
)

type NodeReportType struct {
	GPQR bool
	CKDR bool
	UPRR bool
	UPFR bool
}

func NewNodeReportType(gpqr bool, ckdr bool, uprr bool, upfr bool) (NodeReportType, error) {
	return NodeReportType{
		GPQR: gpqr,
		CKDR: ckdr,
		UPRR: uprr,
		UPFR: upfr,
	}, nil
}

func (nrt NodeReportType) Serialize() []byte {
	buf := new(bytes.Buffer)

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

func (nrt NodeReportType) GetType() IEType {
	return NodeReportTypeIEType
}

func DeserializeNodeReportType(ieValue []byte) (NodeReportType, error) {

	if len(ieValue) < 1 {
		return NodeReportType{}, fmt.Errorf("invalid length for NodeReportType: got %d bytes, expected at least 1", len(ieValue))
	}

	buf := bytes.NewBuffer(ieValue)

	var octet5 byte
	var err error
	if octet5, err = buf.ReadByte(); err != nil {
		return NodeReportType{}, fmt.Errorf("error reading NodeReportType flags: %v", err)
	}

	nrt := NodeReportType{
		GPQR: octet5&0x08 != 0,
		CKDR: octet5&0x04 != 0,
		UPRR: octet5&0x02 != 0,
		UPFR: octet5&0x01 != 0,
	}

	return nrt, nil
}
