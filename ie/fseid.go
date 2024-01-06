package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type FSEID struct {
	IEType uint16
	Length uint16
	V4     bool
	V6     bool
	SEID   uint64
	IPv4   []byte
	IPv6   []byte
}

func NewFSEID(seid uint64, ipv4Address string, ipv6Address string) (FSEID, error) {
	fseid := FSEID{
		IEType: uint16(FSEIDIEType),
		SEID:   seid,
	}
	var length uint16 = 9

	ipv4 := net.ParseIP(ipv4Address)
	ipv6 := net.ParseIP(ipv6Address)
	fseid.IPv4 = ipv4.To4()
	fseid.IPv6 = ipv6.To16()

	fmt.Print("IPv4: ", fseid.IPv4, "\n")
	fmt.Print("IPv6: ", fseid.IPv6, "\n")

	if fseid.IPv4 != nil {
		fseid.V4 = true
		length += 4
	}
	if fseid.IPv6 != nil {
		fseid.V6 = true
		length += 16
	}
	fseid.Length = length
	return fseid, nil
}

func (fseid FSEID) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(fseid.IEType))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(fseid.Length))

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

func DeserializeFSEID(ieType uint16, ieLength uint16, ieValue []byte) (FSEID, error) {
	fmt.Printf("ieType: %v\n", ieType)

	fmt.Printf("ieValue: %v\n", ieValue)
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
		IEType: ieType,
		Length: ieLength,
		V4:     v4,
		V6:     v6,
		SEID:   seid,
		IPv4:   ipv4,
		IPv6:   ipv6,
	}, nil
}

func (fseid FSEID) IsZeroValue() bool {
	return fseid.Length == 0
}
