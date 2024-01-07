package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Cause struct {
	IEType uint16
	Length uint16
	Value  CauseValue
}

type CauseValue uint8

const (
	RequestAccepted              CauseValue = 1
	MoreUsageReportToSend        CauseValue = 2
	RequestRejected              CauseValue = 64
	SessionContextNotFound       CauseValue = 65
	MandatoryIEMissing           CauseValue = 66
	ConditionalIEMissing         CauseValue = 67
	InvalidLength                CauseValue = 68
	MandatoryIEIncorrect         CauseValue = 69
	InvalidForwardingPolicy      CauseValue = 70
	InvalidFTeidAllocation       CauseValue = 71
	NoEstablishedPFCPAssociation CauseValue = 72
	RuleCreationFailure          CauseValue = 73
	PFCPEntityInCongestion       CauseValue = 74
	NoResourcesAvailable         CauseValue = 75
	ServiceNotSupported          CauseValue = 76
	SystemFailure                CauseValue = 77
	RedirectionRequested         CauseValue = 78
)

func NewCause(value CauseValue) (Cause, error) {
	// Validate that value is in the range of supported values
	if value < 1 || value > 78 {
		return Cause{}, fmt.Errorf("invalid value for Cause: %d", value)
	}

	return Cause{
		IEType: uint16(CauseIEType),
		Length: 1,
		Value:  value,
	}, nil
}

func (cause Cause) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(cause.IEType))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(cause.Length))

	// Octet 5: Value (1 byte)
	buf.WriteByte(uint8(cause.Value))

	return buf.Bytes()
}

func (cause Cause) IsZeroValue() bool {
	return cause.Length == 0
}

func DeserializeCause(ieType uint16, ieLength uint16, ieValue []byte) (Cause, error) {
	var cause Cause

	if len(ieValue) != 1 {
		return cause, fmt.Errorf("invalid length for Cause: got %d bytes, want 1", len(ieValue))
	}

	if ieType != uint16(CauseIEType) {
		return cause, fmt.Errorf("invalid IE type: expected %d, got %d", CauseIEType, ieType)
	}

	if ieLength != 1 {
		return cause, fmt.Errorf("invalid length field for Cause: expected 1, got %d", ieLength)
	}

	return Cause{
		IEType: ieType,
		Length: ieLength,
		Value:  CauseValue(ieValue[0]),
	}, nil
}
