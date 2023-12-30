package ie

import (
	"fmt"
	"net"
)

type SourceIPAddress struct {
	IEtype           uint16
	Length           uint16
	MPL              bool
	V4               bool
	V6               bool
	IPv4Address      []byte
	IPv6Address      []byte
	MaskPrefixLength uint8
}

func NewSourceIPAddress(cidr string) (SourceIPAddress, error) {
	sourceIPAddress := SourceIPAddress{
		IEtype: 192,
	}

	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return sourceIPAddress, fmt.Errorf("invalid CIDR")
	}

	if ip.To4() != nil {
		sourceIPAddress.V4 = true
		sourceIPAddress.IPv4Address = ip.To4()
		sourceIPAddress.Length = 6
	} else {
		sourceIPAddress.V6 = true
		sourceIPAddress.IPv6Address = ip.To16()
		sourceIPAddress.Length = 18
	}

	if ipnet != nil {
		sourceIPAddress.MPL = true
		ones, _ := ipnet.Mask.Size()
		sourceIPAddress.MaskPrefixLength = uint8(ones)
	}
	return sourceIPAddress, nil
}

func (sourceIPAddress SourceIPAddress) IsZeroValue() bool {
	return sourceIPAddress.Length == 0
}

func (sourceIPAddress SourceIPAddress) Serialize() []byte {
	var length uint16

	if sourceIPAddress.V4 {
		length = 6
	}
	if sourceIPAddress.V6 {
		length = 18
	}
	bytes := make([]byte, 4+length)
	bytes[0] = byte(sourceIPAddress.IEtype >> 8)
	bytes[1] = byte(sourceIPAddress.IEtype)
	bytes[2] = byte(length >> 8)
	bytes[3] = byte(length)
	if sourceIPAddress.MPL {
		bytes[4] = 0x80
	}
	if sourceIPAddress.V4 {
		bytes[4] |= 0x40
	}
	if sourceIPAddress.V6 {
		bytes[4] |= 0x20
	}
	if sourceIPAddress.V4 {
		copy(bytes[5:9], sourceIPAddress.IPv4Address)
		bytes[9] = sourceIPAddress.MaskPrefixLength
	}
	if sourceIPAddress.V6 {
		copy(bytes[5:21], sourceIPAddress.IPv6Address)
		bytes[21] = sourceIPAddress.MaskPrefixLength
	}
	return bytes
}

func DeserializeSourceIPAddress(ieType uint16, ieLength uint16, ieValue []byte) (SourceIPAddress, error) {
	sourceIPAddress := SourceIPAddress{
		IEtype: ieType,
		Length: ieLength,
	}

	if ieValue[0]&0x80 == 0x80 {
		sourceIPAddress.MPL = true
	}
	if ieValue[0]&0x40 == 0x40 {
		sourceIPAddress.V4 = true
		sourceIPAddress.IPv4Address = ieValue[1:5]
		if sourceIPAddress.MPL {
			sourceIPAddress.MaskPrefixLength = ieValue[5]
		}
	}
	if ieValue[0]&0x20 == 0x20 {
		sourceIPAddress.V6 = true
		sourceIPAddress.IPv6Address = ieValue[1:17]
		if sourceIPAddress.MPL {
			sourceIPAddress.MaskPrefixLength = ieValue[17]
		}
	}
	return sourceIPAddress, nil
}
