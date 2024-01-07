package ie

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type Report int

const (
	UISR Report = iota
	SESR
	TMIR
	UPIR
	ERIR
	USAR
	DLDR
)

type ReportType struct {
	IEType  uint16
	Length  uint16
	Reports []Report
}

func NewReportType(reports []Report) (ReportType, error) {
	return ReportType{
		IEType:  uint16(ReportTypeIEType),
		Length:  1,
		Reports: reports,
	}, nil
}

func (reportType ReportType) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(ReportTypeIEType))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(reportType.Length))

	// Octet 5: Reports
	// Bit 1: DLDR, Bit 2: USAR, Bit 3: ERIR, Bit 4: UPIR, Bit 5: TMIR, Bit 6: SESR, Bit 7: UISR, Bit 8: Spare
	var reportsByte byte = 0
	for _, report := range reportType.Reports {
		reportsByte |= 1 << report
	}
	buf.WriteByte(reportsByte)

	return buf.Bytes()
}

func (reportType ReportType) IsZeroValue() bool {
	return reportType.Length == 0
}

func DeserializeReportType(ieType uint16, ieLength uint16, ieValue []byte) (ReportType, error) {
	if len(ieValue) != int(ieLength) {
		return ReportType{}, errors.New("invalid length for ReportType")
	}

	var reports []Report
	reportsByte := ieValue[0]

	for i := 0; i < 8; i++ {
		if reportsByte&(1<<i) != 0 {
			reports = append(reports, Report(i))
		}
	}

	return ReportType{
		IEType:  ieType,
		Length:  ieLength,
		Reports: reports,
	}, nil
}
