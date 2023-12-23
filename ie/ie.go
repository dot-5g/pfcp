package ie

import (
	"encoding/binary"
	"fmt"
)

type InformationElement interface {
	Serialize() []byte
}

func ParseInformationElements(b []byte) ([]InformationElement, error) {
	var ies []InformationElement

	index := 0

	for index < len(b) {
		if len(b[index:]) < 4 {
			return nil, fmt.Errorf("not enough bytes for IE header")
		}

		ieType := int(binary.BigEndian.Uint16(b[index : index+2]))
		ieLength := int(binary.BigEndian.Uint16(b[index+2 : index+4]))
		index += 4 // Move past the header
		fmt.Printf("IE type: %d, length: %d\n", ieType, ieLength)

		if len(b[index:]) < ieLength {
			return nil, fmt.Errorf("not enough bytes for IE data, expected %d, got %d", ieLength, len(b[index:]))
		}

		ieValue := b[index : index+ieLength]
		var ie InformationElement
		switch ieType {
		case 96:
			ie = DeserializeRecoveryTimeStamp(ieType, ieLength, ieValue)
		default:
			return nil, fmt.Errorf("unknown IE type %d", ieType)
		}

		if ie != nil {
			ies = append(ies, ie)
		}

		index += ieLength
	}

	return ies, nil
}
