package ie

import (
	"bytes"
	"fmt"
	"net"
)

type UEIPAddress struct {
	Header                   Header
	IP6PL                    bool
	CHV6                     bool
	CHV4                     bool
	IPv6D                    bool
	SD                       bool
	V4                       bool
	V6                       bool
	IPv4Address              []byte
	IPv6Address              []byte
	IPv6PrefixDelegationBits uint8
	IPv6PrefixLength         uint8
}

type SourceDestination struct {
	Source      bool
	Destination bool
}

func NewUEIPAddress(ipv4Address string, ipv6Address string, sd SourceDestination, ipv6PrefixDelegationBits uint8, ipv6PrefixLength uint8, chooseV4 bool, chooseV6 bool) (UEIPAddress, error) {
	var sourceDestination bool
	var ipv4Bytes []byte
	var ipv6Bytes []byte
	var length uint16 = 1
	var v4 bool
	var v6 bool
	var ipv6d bool
	var ip6pl bool

	if chooseV4 && ipv4Address != "" {
		return UEIPAddress{}, fmt.Errorf("cannot choose IPv4 and provide IPv4 address")
	}

	if chooseV6 && ipv6Address != "" {
		return UEIPAddress{}, fmt.Errorf("cannot choose IPv6 and provide IPv6 address")
	}

	if ipv6PrefixDelegationBits != 0 {
		if ipv6Address == "" && !chooseV6 {
			return UEIPAddress{}, fmt.Errorf("cannot provide IPv6 prefix delegation bits without IPv6 Address or choosing IPv6")
		}
		ipv6d = true
		length += 1
	}

	if ipv6PrefixLength != 0 {
		if ipv6Address == "" && !chooseV6 {
			return UEIPAddress{}, fmt.Errorf("cannot provide IPv6 prefix length without IPv6 Address or choosing IPv6")
		}
		if ipv6d {
			return UEIPAddress{}, fmt.Errorf("cannot provide IPv6 prefix length with IPv6 prefix delegation bits")
		}
		ip6pl = true
		length += 1
	}

	if ipv4Address != "" {
		ipv4Bytes = net.ParseIP(ipv4Address).To4()
		if ipv4Bytes == nil {
			return UEIPAddress{}, fmt.Errorf("invalid IPv4 address")
		}
		v4 = true
		length += 4
	}

	if ipv6Address != "" {
		ipv6Bytes = net.ParseIP(ipv6Address).To16()
		if ipv6Bytes == nil {
			return UEIPAddress{}, fmt.Errorf("invalid IPv6 address")
		}
		v6 = true
		length += 16
	}

	ieHeader := Header{
		Type:   UEIPAddressIEType,
		Length: length,
	}

	return UEIPAddress{
		Header:                   ieHeader,
		IP6PL:                    ip6pl,
		CHV6:                     chooseV6,
		CHV4:                     chooseV4,
		IPv6D:                    ipv6d,
		SD:                       sourceDestination,
		V4:                       v4,
		V6:                       v6,
		IPv4Address:              ipv4Bytes,
		IPv6Address:              ipv6Bytes,
		IPv6PrefixDelegationBits: ipv6PrefixDelegationBits,
		IPv6PrefixLength:         ipv6PrefixLength,
	}, nil
}

func (ueIPaddress UEIPAddress) IsZeroValue() bool {
	return ueIPaddress.Header.Length == 0
}

func (ueIPaddress UEIPAddress) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 4: Header
	buf.Write(ueIPaddress.Header.Serialize())

	// Octet 5: Bit 1: V6, Bit 2: V4, Bit 3: S/D, Bit 4: IPv6D, Bit 5: CHV4, Bit 6: CHV6, Bit 7: IP6PL, Bit 8: Spare
	var octet5 byte
	if ueIPaddress.IP6PL {
		octet5 |= 1 << 7
	}
	if ueIPaddress.CHV6 {
		octet5 |= 1 << 6
	}
	if ueIPaddress.CHV4 {
		octet5 |= 1 << 5
	}
	if ueIPaddress.IPv6D {
		octet5 |= 1 << 4
	}
	if ueIPaddress.SD {
		octet5 |= 1 << 3
	}
	if ueIPaddress.V4 {
		octet5 |= 1 << 2
	}
	if ueIPaddress.V6 {
		octet5 |= 1 << 1
	}
	buf.WriteByte(octet5)

	// Octet m to (m+3): IPv4 Address
	if ueIPaddress.V4 {
		buf.Write(ueIPaddress.IPv4Address)
	}

	// Octet p to (p+15): IPv6 Address
	if ueIPaddress.V6 {
		buf.Write(ueIPaddress.IPv6Address)
	}

	// Octet r: IPv6 Delegation Bits
	if ueIPaddress.IPv6D {
		buf.WriteByte(ueIPaddress.IPv6PrefixDelegationBits)
	}

	// Octet s: IPv6 Prefix Length
	if ueIPaddress.IP6PL {
		buf.WriteByte(ueIPaddress.IPv6PrefixLength)
	}

	return buf.Bytes()
}

func (ueIPaddress UEIPAddress) SetHeader(ieHeader Header) InformationElement {
	ueIPaddress.Header = ieHeader
	return ueIPaddress
}

func DeserializeUEIPAddress(ieValue []byte) (UEIPAddress, error) {
	if len(ieValue) < 1 {
		return UEIPAddress{}, fmt.Errorf("invalid length for UEIPAddress")
	}

	ueIPAddress := UEIPAddress{}

	octet5 := ieValue[0]
	ueIPAddress.IP6PL = octet5&(1<<7) > 0
	ueIPAddress.CHV6 = octet5&(1<<6) > 0
	ueIPAddress.CHV4 = octet5&(1<<5) > 0
	ueIPAddress.IPv6D = octet5&(1<<4) > 0
	ueIPAddress.SD = octet5&(1<<3) > 0
	ueIPAddress.V4 = octet5&(1<<2) > 0
	ueIPAddress.V6 = octet5&(1<<1) > 0

	index := 1

	if ueIPAddress.V4 {
		if len(ieValue[index:]) < 4 {
			return UEIPAddress{}, fmt.Errorf("invalid length for IPv4 address")
		}
		ueIPAddress.IPv4Address = ieValue[index : index+4]
		index += 4
	}

	if ueIPAddress.V6 {
		if len(ieValue[index:]) < 16 {
			return UEIPAddress{}, fmt.Errorf("invalid length for IPv6 address")
		}
		ueIPAddress.IPv6Address = ieValue[index : index+16]
		index += 16
	}

	if ueIPAddress.IPv6D {
		if len(ieValue[index:]) < 1 {
			return UEIPAddress{}, fmt.Errorf("invalid length for IPv6 prefix delegation bits")
		}
		ueIPAddress.IPv6PrefixDelegationBits = ieValue[index]
		index += 1
	}

	if ueIPAddress.IP6PL {
		if len(ieValue[index:]) < 1 {
			return UEIPAddress{}, fmt.Errorf("invalid length for IPv6 prefix length")
		}
		ueIPAddress.IPv6PrefixLength = ieValue[index]
		index += 1
	}

	return ueIPAddress, nil
}
