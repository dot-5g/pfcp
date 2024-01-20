package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type PDI struct {
	Header          Header
	SourceInterface SourceInterface // Mandatory
	UEIPAddress     UEIPAddress     // Optional
}

func NewPDI(sourceInterface SourceInterface, ueIPAddress UEIPAddress) (PDI, error) {
	ieHeader := Header{
		Type:   PDIIEType,
		Length: sourceInterface.Header.Length + ueIPAddress.Header.Length + 8,
	}

	return PDI{
		Header:          ieHeader,
		SourceInterface: sourceInterface,
		UEIPAddress:     ueIPAddress,
	}, nil
}

func (pdi PDI) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 4: Header
	buf.Write(pdi.Header.Serialize())

	// Octets 5 to n: Source Interface
	serializedSourceInterface := pdi.SourceInterface.Serialize()
	buf.Write(serializedSourceInterface)

	// Octets (n+1) to (n+m): UE IP Address
	if !pdi.UEIPAddress.IsZeroValue() {
		serializedUEIPAddress := pdi.UEIPAddress.Serialize()
		buf.Write(serializedUEIPAddress)
	}

	return buf.Bytes()

}

func (pdi PDI) IsZeroValue() bool {
	return pdi.Header.Length == 0
}

func (pdi PDI) SetHeader(header Header) InformationElement {
	pdi.Header = header
	return pdi
}

func DeserializePDI(ieValue []byte) (PDI, error) {
	if len(ieValue) < 1 {
		return PDI{}, fmt.Errorf("invalid length for PDI: got %d bytes, want at least 1", len(ieValue))
	}

	pdi := PDI{
		SourceInterface: SourceInterface{},
		UEIPAddress:     UEIPAddress{},
	}

	index := 0
	for index < len(ieValue) {
		if index+4 > len(ieValue) {
			return PDI{}, fmt.Errorf("slice bounds out of range")
		}

		currentIEType := binary.BigEndian.Uint16(ieValue[index : index+2])
		currentIELength := binary.BigEndian.Uint16(ieValue[index+2 : index+4])

		if index+4+int(currentIELength) > len(ieValue) {
			return PDI{}, fmt.Errorf("slice bounds out of range")
		}

		currentIEValue := ieValue[index+4 : index+4+int(currentIELength)]

		switch IEType(currentIEType) {
		case SourceInterfaceIEType:
			sourceInterfaceHeader := Header{
				Type:   IEType(currentIEType),
				Length: currentIELength,
			}

			tempSourceInterface, err := DeserializeSourceInterface(currentIEValue)
			if err != nil {
				return PDI{}, fmt.Errorf("failed to deserialize Source Interface: %v", err)
			}

			sourceInterface, ok := tempSourceInterface.SetHeader(sourceInterfaceHeader).(SourceInterface)
			if !ok {
				return PDI{}, fmt.Errorf("type assertion to FarID failed")
			}

			pdi.SourceInterface = sourceInterface
		case UEIPAddressIEType:
			ueIPAddressHeader := Header{
				Type:   IEType(currentIEType),
				Length: currentIELength,
			}
			tempUeIPAddress, err := DeserializeUEIPAddress(currentIEValue)
			if err != nil {
				return PDI{}, fmt.Errorf("failed to deserialize UE IP Address: %v", err)
			}
			ueIPAddress, ok := tempUeIPAddress.SetHeader(ueIPAddressHeader).(UEIPAddress)
			if !ok {
				return PDI{}, fmt.Errorf("type assertion to FarID failed")
			}
			pdi.UEIPAddress = ueIPAddress
		}
		index += 4 + int(currentIELength)

	}

	return pdi, nil
}
