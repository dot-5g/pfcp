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
	Reports []Report
}

func NewReportType(reports []Report) (ReportType, error) {
	return ReportType{
		Reports: reports,
	}, nil
}

func (reportType ReportType) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octet 5: Reports
	// Bit 1: DLDR, Bit 2: USAR, Bit 3: ERIR, Bit 4: UPIR, Bit 5: TMIR, Bit 6: SESR, Bit 7: UISR, Bit 8: Spare
	var reportsByte byte = 0
	for _, report := range reportType.Reports {
		reportsByte |= 1 << report
	}
	buf.WriteByte(reportsByte)

	return buf.Bytes()
}

func (reportType ReportType) GetType() IEType {
	return ReportTypeIEType
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
