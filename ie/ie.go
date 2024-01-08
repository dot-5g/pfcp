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

func DeserializeIEHeader(payload []byte) (IEHeader, error) {
	var ieHeader IEHeader

	if len(payload) < IEHeaderLength {
		return ieHeader, fmt.Errorf("not enough bytes for IE header")
	}

	ieHeader.Type = IEType(binary.BigEndian.Uint16(payload[:2]))
	ieHeader.Length = binary.BigEndian.Uint16(payload[2:4])

	return ieHeader, nil
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

		ieHeader := IEHeader{
			Type:   ieType,
			Length: ieLength,
		}

		if len(b[index:]) < int(ieHeader.Length) {
			return nil, fmt.Errorf("not enough bytes for IE data, expected %d, got %d", ieHeader.Length, len(b[index:]))
		}

		ieValue := b[index : index+int(ieHeader.Length)]
		var ie InformationElement
		switch ieHeader.Type {
		case CauseIEType:
			ie, err = DeserializeCause(ieHeader, ieValue)
		case NodeIDIEType:
			ie, err = DeserializeNodeID(ieHeader, ieValue)
		case RecoveryTimeStampIEType:
			ie, err = DeserializeRecoveryTimeStamp(ieHeader, ieValue)
		case NodeReportTypeIEType:
			ie, err = DeserializeNodeReportType(ieHeader, ieValue)
		case SourceIPAddressIEType:
			ie, err = DeserializeSourceIPAddress(ieHeader, ieValue)
		case UPFunctionFeaturesIEType:
			ie, err = DeserializeUPFunctionFeatures(ieHeader, ieValue)
		case FSEIDIEType:
			ie, err = DeserializeFSEID(ieHeader, ieValue)
		case PDRIDIEType:
			ie, err = DeserializePDRID(ieHeader, ieValue)
		case PrecedenceIEType:
			ie, err = DeserializePrecedence(uint16(ieHeader.Type), ieHeader.Length, ieValue)
		case SourceInterfaceIEType:
			ie, err = DeserializeSourceInterface(uint16(ieHeader.Type), ieHeader.Length, ieValue)
		case PDIIEType:
			ie, err = DeserializePDI(uint16(ieHeader.Type), ieHeader.Length, ieValue)
		case CreatePDRIEType:
			ie, err = DeserializeCreatePDR(uint16(ieHeader.Type), ieHeader.Length, ieValue)
		case FARIDIEType:
			ie, err = DeserializeFARID(uint16(ieHeader.Type), ieHeader.Length, ieValue)
		case ApplyActionIEType:
			ie, err = DeserializeApplyAction(uint16(ieHeader.Type), ieHeader.Length, ieValue)
		case CreateFARIEType:
			ie, err = DeserializeCreateFAR(uint16(ieHeader.Type), ieHeader.Length, ieValue)
		case ReportTypeIEType:
			ie, err = DeserializeReportType(uint16(ieHeader.Type), ieHeader.Length, ieValue)
		default:
			err = fmt.Errorf("unknown IE type %d", ieHeader.Type)
		}

		if ie != nil {
			ies = append(ies, ie)
		}

		index += int(ieHeader.Length)
	}

	return ies, err
}
