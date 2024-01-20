package ie

import (
	"bytes"
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
	Header  Header
	Reports []Report
}

func NewReportType(reports []Report) (ReportType, error) {
	ieHeader := Header{
		Type:   IEType(ReportTypeIEType),
		Length: 1,
	}

	return ReportType{
		Header:  ieHeader,
		Reports: reports,
	}, nil
}

func (reportType ReportType) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 4: Header
	buf.Write(reportType.Header.Serialize())

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
	return reportType.Header.Length == 0
}

func (reportType ReportType) SetHeader(header Header) InformationElement {
	reportType.Header = header
	return reportType
}

func DeserializeReportType(ieValue []byte) (ReportType, error) {

	var reports []Report
	reportsByte := ieValue[0]

	for i := 0; i < 8; i++ {
		if reportsByte&(1<<i) != 0 {
			reports = append(reports, Report(i))
		}
	}

	return ReportType{
		Reports: reports,
	}, nil
}
