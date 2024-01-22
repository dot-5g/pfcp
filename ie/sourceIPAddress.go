package ie

import (
	"bytes"
	"net"
)

type SourceIPAddress struct {
	MPL              bool
	V4               bool
	V6               bool
	IPv4Address      []byte
	IPv6Address      []byte
	MaskPrefixLength uint8
}

func NewSourceIPAddress(ipv4Address string, ipv6Address string) (SourceIPAddress, error) {
	var v4 bool
	var v6 bool
	var mpl bool
	var maskPrefixLength uint8
	var ipv4Addr []byte
	var ipv6Addr []byte
	ipv4, ipv4net, _ := net.ParseCIDR(ipv4Address)
	ipv6, ipv6net, _ := net.ParseCIDR(ipv6Address)

	if ipv4.To4() != nil {
		v4 = true
		ipv4Addr = ipv4.To4()
		mpl = true
		ones, _ := ipv4net.Mask.Size()
		maskPrefixLength = uint8(ones)
	}

	if ipv6.To16() != nil {
		v6 = true
		ipv6Addr = ipv6.To16()
		mpl = true
		ones, _ := ipv6net.Mask.Size()
		maskPrefixLength = uint8(ones)
	}

	sourceIPAddress := SourceIPAddress{
		MPL:              mpl,
		V4:               v4,
		V6:               v6,
		IPv4Address:      ipv4Addr,
		IPv6Address:      ipv6Addr,
		MaskPrefixLength: maskPrefixLength,
	}

	return sourceIPAddress, nil
}

func (sourceIPAddress SourceIPAddress) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octet 5: Spare, Spare, Spare, MPL, V4, V6
	var octet5 byte
	if sourceIPAddress.MPL {
		octet5 |= 1 << 7
	}
	if sourceIPAddress.V4 {
		octet5 |= 1 << 6
	}
	if sourceIPAddress.V6 {
		octet5 |= 1 << 5
	}
	buf.WriteByte(octet5)

	// Octets 6 to 9: IPv4 Address
	if sourceIPAddress.V4 {
		buf.Write(sourceIPAddress.IPv4Address)
		buf.WriteByte(sourceIPAddress.MaskPrefixLength)
	}

	// Octets 6 to 21: IPv6 Address
	if sourceIPAddress.V6 {
		buf.Write(sourceIPAddress.IPv6Address)
		buf.WriteByte(sourceIPAddress.MaskPrefixLength)
	}

	return buf.Bytes()
}

func (sourceIPAddress SourceIPAddress) GetType() IEType {
	return SourceIPAddressIEType
}

func DeserializeSourceIPAddress(ieValue []byte) (SourceIPAddress, error) {
	var mpl bool
	var v4 bool
	var v6 bool
	var ipv4Address []byte
	var ipv6Address []byte
	var maskPrefixLength uint8

	if ieValue[0]&0x80 == 0x80 {
		mpl = true
	}
	if ieValue[0]&0x40 == 0x40 {
		v4 = true
		ipv4Address = ieValue[1:5]
		if mpl {
			maskPrefixLength = ieValue[5]
		}
	}
	if ieValue[0]&0x20 == 0x20 {
		v6 = true
		ipv6Address = ieValue[1:17]
		if mpl {
			maskPrefixLength = ieValue[17]
		}
	}

	return SourceIPAddress{
		MPL:              mpl,
		V4:               v4,
		V6:               v6,
		IPv4Address:      ipv4Address,
		IPv6Address:      ipv6Address,
		MaskPrefixLength: maskPrefixLength,
	}, nil
}
