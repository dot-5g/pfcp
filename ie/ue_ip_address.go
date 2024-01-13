package ie

import (
	"bytes"
	"errors"
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

func NewUEIPAddress(ipv4CIDR string, ipv6CIDR string, sd SourceDestination, ipv6PrefixDelegationBits uint8) (UEIPAddress, error) {
	var ipv4Address, ipv6Address []byte
	var ipv6PrefixLength uint8
	var length uint16 = 1

	if ipv4CIDR != "" {
		ip, _, err := net.ParseCIDR(ipv4CIDR)
		if err != nil {
			return UEIPAddress{}, err
		}
		ipv4Address = ip.To4()
		if ipv4Address == nil {
			return UEIPAddress{}, errors.New("invalid IPv4 address")
		}
		length += 4
	}

	if ipv6CIDR != "" {
		ip, ipv6Network, err := net.ParseCIDR(ipv6CIDR)
		if err != nil {
			return UEIPAddress{}, err
		}
		ipv6Address = ip.To16()
		if ipv6Address == nil {
			return UEIPAddress{}, errors.New("invalid IPv6 address")
		}
		ones, _ := ipv6Network.Mask.Size()
		ipv6PrefixLength = uint8(ones)
		length += 18
	}

	var sourceDestination bool
	if sd.Source {
		sourceDestination = false
	} else if sd.Destination {
		sourceDestination = true
	}

	ieHeader := Header{
		Type:   UEIPAddressIEType,
		Length: length,
	}

	return UEIPAddress{
		Header:                   ieHeader,
		IP6PL:                    ipv6PrefixLength != 0,
		CHV6:                     ipv6Address == nil,
		CHV4:                     ipv4Address == nil,
		IPv6D:                    ipv6PrefixDelegationBits != 0,
		SD:                       sourceDestination,
		V4:                       ipv4Address != nil,
		V6:                       ipv6Address != nil,
		IPv4Address:              ipv4Address,
		IPv6Address:              ipv6Address,
		IPv6PrefixDelegationBits: ipv6PrefixDelegationBits,
		IPv6PrefixLength:         ipv6PrefixLength,
	}, nil
}

func (ueIPaddress *UEIPAddress) Serialize() []byte {
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
	if ueIPaddress.IPv6D {
		buf.WriteByte(ueIPaddress.IPv6PrefixLength)
	}

	return buf.Bytes()

}
