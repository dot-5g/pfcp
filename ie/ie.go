package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const IEHeaderLength = 4

type IEType uint16

const (
	CreatePDRIEType          IEType = 1
	CreateFARIEType          IEType = 3
	PDIIEType                IEType = 17
	CauseIEType              IEType = 19
	SourceInterfaceIEType    IEType = 20
	PrecedenceIEType         IEType = 29
	ReportTypeIEType         IEType = 39
	UPFunctionFeaturesIEType IEType = 43
	ApplyActionIEType        IEType = 44
	PDRIDIEType              IEType = 56
	FSEIDIEType              IEType = 57
	NodeIDIEType             IEType = 60
	RecoveryTimeStampIEType  IEType = 96
	NodeReportTypeIEType     IEType = 101
	FARIDIEType              IEType = 108
	SourceIPAddressIEType    IEType = 192
)

type IEHeader struct {
	Type   IEType
	Length uint16
}

func (ieHeader *IEHeader) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(ieHeader.Type))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(ieHeader.Length))

	return buf.Bytes()
}

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
			ie, err = DeserializeCause(uint16(ieType), ieLength, ieValue)
		case NodeIDIEType:
			ie, err = DeserializeNodeID(uint16(ieType), ieLength, ieValue)
		case RecoveryTimeStampIEType:
			ie, err = DeserializeRecoveryTimeStamp(uint16(ieType), ieLength, ieValue)
		case NodeReportTypeIEType:
			ie, err = DeserializeNodeReportType(uint16(ieType), ieLength, ieValue)
		case SourceIPAddressIEType:
			ie, err = DeserializeSourceIPAddress(uint16(ieType), ieLength, ieValue)
		case UPFunctionFeaturesIEType:
			ie, err = DeserializeUPFunctionFeatures(uint16(ieType), ieLength, ieValue)
		case FSEIDIEType:
			ie, err = DeserializeFSEID(uint16(ieType), ieLength, ieValue)
		case PDRIDIEType:
			ie, err = DeserializePDRID(uint16(ieType), ieLength, ieValue)
		case PrecedenceIEType:
			ie, err = DeserializePrecedence(uint16(ieType), ieLength, ieValue)
		case SourceInterfaceIEType:
			ie, err = DeserializeSourceInterface(uint16(ieType), ieLength, ieValue)
		case PDIIEType:
			ie, err = DeserializePDI(uint16(ieType), ieLength, ieValue)
		case CreatePDRIEType:
			ie, err = DeserializeCreatePDR(uint16(ieType), ieLength, ieValue)
		case FARIDIEType:
			ie, err = DeserializeFARID(uint16(ieType), ieLength, ieValue)
		case ApplyActionIEType:
			ie, err = DeserializeApplyAction(uint16(ieType), ieLength, ieValue)
		case CreateFARIEType:
			ie, err = DeserializeCreateFAR(uint16(ieType), ieLength, ieValue)
		case ReportTypeIEType:
			ie, err = DeserializeReportType(uint16(ieType), ieLength, ieValue)
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
