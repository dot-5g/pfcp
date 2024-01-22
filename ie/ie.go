// Package ie contains the Information Elements (IEs) used by the PFCP protocol.
package ie

import (
	"encoding/binary"
	"fmt"
)

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
	UEIPAddressIEType        IEType = 93
	RecoveryTimeStampIEType  IEType = 96
	NodeReportTypeIEType     IEType = 101
	FARIDIEType              IEType = 108
	SourceIPAddressIEType    IEType = 192
)

type InformationElement interface {
	Serialize() []byte
	GetType() IEType
}

func DeserializeInformationElements(b []byte) ([]InformationElement, error) {
	var ies []InformationElement
	var err error

	index := 0

	for index < len(b) {
		if len(b[index:]) < HeaderLength {
			return nil, fmt.Errorf("not enough bytes for IE header")
		}

		ieType := IEType(binary.BigEndian.Uint16(b[index : index+2]))
		ieLength := binary.BigEndian.Uint16(b[index+2 : index+4])
		index += HeaderLength

		if len(b[index:]) < int(ieLength) {
			return nil, fmt.Errorf("not enough bytes for IE data, expected %d, got %d", ieLength, len(b[index:]))
		}

		ieValue := b[index : index+int(ieLength)]
		var ie InformationElement
		switch ieType {
		case CauseIEType:
			ie, err = DeserializeCause(ieValue)
		case NodeIDIEType:
			ie, err = DeserializeNodeID(ieValue)
		case RecoveryTimeStampIEType:
			ie, err = DeserializeRecoveryTimeStamp(ieValue)
		case NodeReportTypeIEType:
			ie, err = DeserializeNodeReportType(ieValue)
		case SourceIPAddressIEType:
			ie, err = DeserializeSourceIPAddress(ieValue)
		case UPFunctionFeaturesIEType:
			ie, err = DeserializeUPFunctionFeatures(ieValue)
		case FSEIDIEType:
			ie, err = DeserializeFSEID(ieValue)
		case PDRIDIEType:
			ie, err = DeserializePDRID(ieValue)
		case PrecedenceIEType:
			ie, err = DeserializePrecedence(ieValue)
		case SourceInterfaceIEType:
			ie, err = DeserializeSourceInterface(ieValue)
		case PDIIEType:
			ie, err = DeserializePDI(ieValue)
		case CreatePDRIEType:
			ie, err = DeserializeCreatePDR(ieValue)
		case FARIDIEType:
			ie, err = DeserializeFARID(ieValue)
		case ApplyActionIEType:
			ie, err = DeserializeApplyAction(ieValue)
		case CreateFARIEType:
			ie, err = DeserializeCreateFAR(ieValue)
		case ReportTypeIEType:
			ie, err = DeserializeReportType(ieValue)
		case UEIPAddressIEType:
			ie, err = DeserializeUEIPAddress(ieValue)
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
