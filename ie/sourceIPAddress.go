package ie

import (
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

func NewSourceIPAddress(ipv4Address string, ipv6Address string) (SourceIPAddress, error) {
	sourceIPAddress := SourceIPAddress{
		IEtype: uint16(SourceIPAddressIEType),
	}

	length := 2

	ipv4, ipv4net, _ := net.ParseCIDR(ipv4Address)
	ipv6, ipv6net, _ := net.ParseCIDR(ipv6Address)

	if ipv4.To4() != nil {
		sourceIPAddress.V4 = true
		sourceIPAddress.IPv4Address = ipv4.To4()
		length += 4
		sourceIPAddress.MPL = true
		ones, _ := ipv4net.Mask.Size()
		sourceIPAddress.MaskPrefixLength = uint8(ones)
	}

	if ipv6.To16() != nil {
		sourceIPAddress.V6 = true
		sourceIPAddress.IPv6Address = ipv6.To16()
		sourceIPAddress.Length = 18
		length += 16
		sourceIPAddress.MPL = true
		ones, _ := ipv6net.Mask.Size()
		sourceIPAddress.MaskPrefixLength = uint8(ones)
	}
	sourceIPAddress.Length = uint16(length)

	return sourceIPAddress, nil
}

func (sourceIPAddress SourceIPAddress) IsZeroValue() bool {
	return sourceIPAddress.Length == 0
}

func (sourceIPAddress SourceIPAddress) Serialize() []byte {
	bytes := make([]byte, 4+sourceIPAddress.Length)
	bytes[0] = byte(sourceIPAddress.IEtype >> 8)
	bytes[1] = byte(sourceIPAddress.IEtype)
	bytes[2] = byte(sourceIPAddress.Length >> 8)
	bytes[3] = byte(sourceIPAddress.Length)
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
