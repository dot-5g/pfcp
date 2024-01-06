package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type ApplyAction struct {
	IEType uint16
	Length uint16
	DFRT   bool
	IPMD   bool
	IPMA   bool
	DUPL   bool
	NOCP   bool
	BUFF   bool
	FORW   bool
	DROP   bool
	DDPN   bool
	BDPN   bool
	EDRT   bool
}

func NewApplyAction(dfrt, ipmd, ipma, dupl, nocp, buff, forw, drop, ddpn, bdpn, edrt bool) ApplyAction {
	return ApplyAction{
		IEType: uint16(ApplyActionIEType),
		Length: 2,
		DFRT:   dfrt,
		IPMD:   ipmd,
		IPMA:   ipma,
		DUPL:   dupl,
		NOCP:   nocp,
		BUFF:   buff,
		FORW:   forw,
		DROP:   drop,
		DDPN:   ddpn,
		BDPN:   bdpn,
		EDRT:   edrt,
	}
}

func (applyaction ApplyAction) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(applyaction.IEType))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(applyaction.Length))

	// Octet 5: DFRT (bit 8), IPMD (bit 7), IPMA (bit 6), DUPL (bit 5), NOCP (bit 4), BUFF (bit 3), FORW (bit 2), DROP (bit 1)
	var byte5 byte
	if applyaction.DFRT {
		byte5 |= 1 << 7
	}
	if applyaction.IPMD {
		byte5 |= 1 << 6
	}
	if applyaction.IPMA {
		byte5 |= 1 << 5
	}
	if applyaction.DUPL {
		byte5 |= 1 << 4
	}
	if applyaction.NOCP {
		byte5 |= 1 << 3
	}
	if applyaction.BUFF {
		byte5 |= 1 << 2
	}
	if applyaction.FORW {
		byte5 |= 1 << 1
	}
	if applyaction.DROP {
		byte5 |= 1
	}
	buf.WriteByte(byte5)

	// Octet 6: Spare (bits 8 to 4), DDPN (bit 3), BDPN (bit 2), EDRT (bit 1)
	var byte6 byte
	if applyaction.DDPN {
		byte6 |= 1 << 2
	}
	if applyaction.BDPN {
		byte6 |= 1 << 1
	}
	if applyaction.EDRT {
		byte6 |= 1
	}
	buf.WriteByte(byte6)

	return buf.Bytes()
}

func (applyaction ApplyAction) IsZeroValue() bool {
	return applyaction.Length == 0
}

func DeserializeApplyAction(ieType uint16, ieLength uint16, ieValue []byte) (ApplyAction, error) {
	var applyaction ApplyAction

	if ieType != uint16(ApplyActionIEType) {
		return applyaction, fmt.Errorf("invalid IE type: expected %d, got %d", ApplyActionIEType, ieType)
	}

	if ieLength != 2 {
		return applyaction, fmt.Errorf("invalid length field for ApplyAction: expected 2, got %d", ieLength)
	}

	if len(ieValue) != 2 {
		return applyaction, fmt.Errorf("invalid length for ApplyAction: got %d bytes, want 2", len(ieValue))
	}

	applyaction.IEType = ieType
	applyaction.Length = ieLength

	// Deserialize the first byte (Octet 5)
	byte5 := ieValue[0]
	applyaction.DFRT = byte5&(1<<7) != 0
	applyaction.IPMD = byte5&(1<<6) != 0
	applyaction.IPMA = byte5&(1<<5) != 0
	applyaction.DUPL = byte5&(1<<4) != 0
	applyaction.NOCP = byte5&(1<<3) != 0
	applyaction.BUFF = byte5&(1<<2) != 0
	applyaction.FORW = byte5&(1<<1) != 0
	applyaction.DROP = byte5&1 != 0

	// Deserialize the second byte (Octet 6)
	byte6 := ieValue[1]
	applyaction.DDPN = byte6&(1<<2) != 0
	applyaction.BDPN = byte6&(1<<1) != 0
	applyaction.EDRT = byte6&1 != 0

	return applyaction, nil
}
