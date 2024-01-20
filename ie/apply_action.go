package ie

import (
	"bytes"
	"fmt"
)

type ApplyAction struct {
	Header Header
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

type ApplyActionFlag int
type ApplyActionExtraFlag int

const (
	DROP ApplyActionFlag = iota
	FORW
	BUFF
	IPMA
	IPMD
)

const (
	NOCP ApplyActionExtraFlag = iota
	BDPN
	DDPN
	DUPL
	DFRT
	EDRT
)

func contains(flags []ApplyActionExtraFlag, flag ApplyActionExtraFlag) bool {
	for _, f := range flags {
		if f == flag {
			return true
		}
	}
	return false
}

func NewApplyAction(flag ApplyActionFlag, extraFlags []ApplyActionExtraFlag) (ApplyAction, error) {
	var dfrt bool
	var ipmd bool
	var ipma bool
	var dupl bool
	var nocp bool
	var buff bool
	var forw bool
	var drop bool
	var ddpn bool
	var bdpn bool
	var edrt bool

	ieHeader := Header{
		Type:   IEType(ApplyActionIEType),
		Length: 2,
	}

	if (contains(extraFlags, NOCP) || contains(extraFlags, BDPN) || contains(extraFlags, DDPN)) && flag != BUFF {
		return ApplyAction{}, fmt.Errorf("the NOCP flag, BDPN and DDPN flag may only be set if the BUFF flag is set")
	}

	if contains(extraFlags, DUPL) && flag == IPMA {
		return ApplyAction{}, fmt.Errorf("the DUPL flag may be set with any of the DROP, FORW, BUFF and NOCP flags")
	}

	if contains(extraFlags, DFRT) && flag != FORW {
		return ApplyAction{}, fmt.Errorf("the DFRT flag may only be set if the FORW flag is set")
	}

	if contains(extraFlags, EDRT) && flag != FORW {
		return ApplyAction{}, fmt.Errorf("the EDRT flag may only be set if the FORW flag is set")
	}

	switch flag {
	case DROP:
		drop = true
		if contains(extraFlags, DUPL) {
			dupl = true
		}
	case FORW:
		forw = true
		if contains(extraFlags, DUPL) {
			dupl = true
		}
		if contains(extraFlags, DFRT) {
			dfrt = true
		}
		if contains(extraFlags, EDRT) {
			edrt = true
		}
	case BUFF:
		buff = true
		if contains(extraFlags, DUPL) {
			dupl = true
		}
		if contains(extraFlags, NOCP) {
			nocp = true
		}
		if contains(extraFlags, BDPN) {
			bdpn = true
		}
		if contains(extraFlags, DDPN) {
			ddpn = true
		}
	case IPMA:
		ipma = true
	case IPMD:
		ipmd = true
		if contains(extraFlags, DUPL) {
			dupl = true
		}
	}

	return ApplyAction{
		Header: ieHeader,
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
	}, nil
}

func (applyaction ApplyAction) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 4: Header
	buf.Write(applyaction.Header.Serialize())

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
	return applyaction.Header.Length == 0
}

func (applyaction ApplyAction) SetHeader(header Header) InformationElement {
	applyaction.Header = header
	return applyaction
}

func DeserializeApplyAction(ieValue []byte) (ApplyAction, error) {
	var applyaction ApplyAction

	if len(ieValue) != 2 {
		return ApplyAction{}, fmt.Errorf("invalid length for ApplyAction: got %d bytes, want 2", len(ieValue))
	}

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
