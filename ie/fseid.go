package ie

import (
	"bytes"
	"encoding/binary"
	"net"
)

type FSEID struct {
	Header IEHeader
	V4     bool
	V6     bool
	SEID   uint64
	IPv4   []byte
	IPv6   []byte
}

func NewFSEID(seid uint64, ipv4Address string, ipv6Address string) (FSEID, error) {
	var length uint16 = 9
	var v4 bool
	var v6 bool
	var ipv4Addr []byte
	var ipv6Addr []byte

	ipv4 := net.ParseIP(ipv4Address)
	ipv6 := net.ParseIP(ipv6Address)
	ipv4Addr = ipv4.To4()
	ipv6Addr = ipv6.To16()

	if ipv4Addr != nil {
		v4 = true
		length += 4
	}
	if ipv6Addr != nil {
		v6 = true
		length += 16
	}

	ieHeader := IEHeader{
		Type:   FSEIDIEType,
		Length: length,
	}

	fseid := FSEID{
		Header: ieHeader,
		V4:     v4,
		V6:     v6,
		SEID:   seid,
		IPv4:   ipv4Addr,
		IPv6:   ipv6Addr,
	}

	return fseid, nil
}

func (fseid FSEID) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 4: Header
	buf.Write(fseid.Header.Serialize())

	// Octet 5: Spare (6 bits) + V4 (1 bit) + V6 (1 bit)
	var flags byte
	if fseid.V4 {
		flags |= 1 << 1 // Set the second bit from the right if V4 is true
	}
	if fseid.V6 {
		flags |= 1 << 0 // Set the first bit from the right if V6 is true
	}
	buf.WriteByte(flags)

	// Octets 6 13: SEID
	binary.Write(buf, binary.BigEndian, fseid.SEID)

	// Octet m to (m+3) IPv4 address
	if fseid.V4 {
		buf.Write(fseid.IPv4)
	}

	// Octet p  to (p+15): IPv6 address
	if fseid.V6 {
		buf.Write(fseid.IPv6)
	}

	return buf.Bytes()
}

func DeserializeFSEID(ieHeader IEHeader, ieValue []byte) (FSEID, error) {
	v4 := ieValue[0]&0x02 > 0
	v6 := ieValue[0]&0x01 > 0
	seid := binary.BigEndian.Uint64(ieValue[1:9])
	var ipv4 []byte
	var ipv6 []byte

	v4StartByte := 9
	v6StartByte := 9

	if v4 {
		ipv4 = ieValue[v4StartByte : v4StartByte+4]
		v6StartByte += 4
	} else {
		ipv4 = nil
	}

	if v6 {
		ipv6 = ieValue[v6StartByte : v6StartByte+16]
	} else {
		ipv6 = nil
	}

	return FSEID{
		Header: ieHeader,
		V4:     v4,
		V6:     v6,
		SEID:   seid,
		IPv4:   ipv4,
		IPv6:   ipv6,
	}, nil
}

func (fseid FSEID) IsZeroValue() bool {
	return fseid.Header.Length == 0
}
