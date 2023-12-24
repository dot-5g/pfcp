package ie

import (
	"encoding/binary"
	"fmt"
)

const IEHeaderLength = 4

type InformationElement interface {
	Serialize() []byte
}

func ParseInformationElements(b []byte) ([]InformationElement, error) {
	var ies []InformationElement
	var err error

	index := 0

	for index < len(b) {
		if len(b[index:]) < IEHeaderLength {
			return nil, fmt.Errorf("not enough bytes for IE header")
		}

		ieType := binary.BigEndian.Uint16(b[index : index+2])
		ieLength := binary.BigEndian.Uint16(b[index+2 : index+4])
		index += IEHeaderLength
		if len(b[index:]) < int(ieLength) {
			return nil, fmt.Errorf("not enough bytes for IE data, expected %d, got %d", ieLength, len(b[index:]))
		}

		ieValue := b[index : index+int(ieLength)]
		var ie InformationElement
		switch ieType {
		case 60:
			ie = DeserializeNodeID(ieType, ieLength, ieValue)
		case 96:
			ie = DeserializeRecoveryTimeStamp(ieType, ieLength, ieValue)
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
