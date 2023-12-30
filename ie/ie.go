package ie

import (
	"encoding/binary"
	"fmt"
)

const IEHeaderLength = 4

type IEType uint16

const (
	CauseIEType              IEType = 19
	NodeIDIEType             IEType = 60
	RecoveryTimeStampIEType  IEType = 96
	NodeReportTypeIEType     IEType = 101
	SourceIPAddressIEType    IEType = 192
	UPFunctionFeaturesIEType IEType = 43
)

type InformationElement interface {
	Serialize() []byte
	IsZeroValue() bool
}

func ParseInformationElements(b []byte) ([]InformationElement, error) {
	var ies []InformationElement
	var err error

	index := 0

	for index < len(b) {
		if len(b[index:]) < IEHeaderLength {
			return nil, fmt.Errorf("not enough bytes for IE header")
		}

		ieType := IEType(binary.BigEndian.Uint16(b[index : index+2]))
		ieLength := binary.BigEndian.Uint16(b[index+2 : index+4])
		index += IEHeaderLength
		if len(b[index:]) < int(ieLength) {
			return nil, fmt.Errorf("not enough bytes for IE data, expected %d, got %d", ieLength, len(b[index:]))
		}

		ieValue := b[index : index+int(ieLength)]
		var ie InformationElement
		switch ieType {
		case CauseIEType:
			ie = DeserializeCause(uint16(ieType), ieLength, ieValue)
		case NodeIDIEType:
			ie = DeserializeNodeID(uint16(ieType), ieLength, ieValue)
		case RecoveryTimeStampIEType:
			ie = DeserializeRecoveryTimeStamp(uint16(ieType), ieLength, ieValue)
		case NodeReportTypeIEType:
			ie = DeserializeNodeReportType(uint16(ieType), ieLength, ieValue)
		case SourceIPAddressIEType:
			ie, err = DeserializeSourceIPAddress(uint16(ieType), ieLength, ieValue)
		case UPFunctionFeaturesIEType:
			ie, err = DeserializeUPFunctionFeatures(uint16(ieType), ieLength, ieValue)
		default:
			err = fmt.Errorf("unknown IE type %d", ieType)
		}

		if ie != nil {
			ies = append(ies, ie)
		}

		index += int(ieLength)
	}

	return ies, err
}
