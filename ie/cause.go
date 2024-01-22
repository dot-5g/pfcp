package ie

import (
	"bytes"
	"fmt"
)

type Cause struct {
	Value CauseValue
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
		Value: value,
	}, nil
}

func (cause Cause) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octet 5: Value (1 byte)
	buf.WriteByte(uint8(cause.Value))

	return buf.Bytes()
}

func (cause Cause) GetType() IEType {
	return CauseIEType
}

func DeserializeCause(ieValue []byte) (Cause, error) {
	if len(ieValue) != 1 {
		return Cause{}, fmt.Errorf("invalid length for Cause: got %d bytes, want 1", len(ieValue))
	}

	return Cause{
		Value: CauseValue(ieValue[0]),
	}, nil
}
