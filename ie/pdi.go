package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type PDI struct {
	SourceInterface SourceInterface // Mandatory
	UEIPAddress     UEIPAddress     // Optional
}

func NewPDI(sourceInterface SourceInterface, ueIPAddress UEIPAddress) (PDI, error) {
	return PDI{
		SourceInterface: sourceInterface,
		UEIPAddress:     ueIPAddress,
	}, nil
}

func (pdi PDI) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 5 to n: Source Interface

	serializedSourceInterface := pdi.SourceInterface.Serialize()
	sourceInterfaceLength := uint16(len(serializedSourceInterface))
	sourceInterfaceHeader := Header{
		Type:   pdi.SourceInterface.GetType(),
		Length: sourceInterfaceLength,
	}
	buf.Write(sourceInterfaceHeader.Serialize())
	buf.Write(serializedSourceInterface)

	// Octets (n+1) to (n+m): UE IP Address
	serializedUEIPAddress := pdi.UEIPAddress.Serialize()
	ueIpAddressLength := uint16(len(serializedUEIPAddress))
	ueIPAddressHeader := Header{
		Type:   pdi.UEIPAddress.GetType(),
		Length: ueIpAddressLength,
	}
	buf.Write(ueIPAddressHeader.Serialize())
	buf.Write(serializedUEIPAddress)

	return buf.Bytes()

}

func (pdi PDI) GetType() IEType {
	return PDIIEType
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
			sourceInterface, err := DeserializeSourceInterface(currentIEValue)
			if err != nil {
				return PDI{}, fmt.Errorf("failed to deserialize Source Interface: %v", err)
			}
			pdi.SourceInterface = sourceInterface
		case UEIPAddressIEType:
			ueIPAddress, err := DeserializeUEIPAddress(currentIEValue)
			if err != nil {
				return PDI{}, fmt.Errorf("failed to deserialize UE IP Address: %v", err)
			}
			pdi.UEIPAddress = ueIPAddress
		}
		index += 4 + int(currentIELength)
	}

	return pdi, nil
}
