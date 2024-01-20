package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

const ntpEpochOffset = 2208988800 // Offset between Unix and NTP epoch (seconds)

type RecoveryTimeStamp struct {
	Header Header
	Value  int64 // Seconds since 1900
}

func NewRecoveryTimeStamp(value time.Time) (RecoveryTimeStamp, error) {
	ieHeader := Header{
		Type:   RecoveryTimeStampIEType,
		Length: 4,
	}
	return RecoveryTimeStamp{
		Header: ieHeader,
		Value:  value.Unix() + ntpEpochOffset,
	}, nil
}

func (rt RecoveryTimeStamp) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 4: Header
	buf.Write(rt.Header.Serialize())

	// Octets 5 to 8: Value
	binary.Write(buf, binary.BigEndian, uint32(rt.Value))

	return buf.Bytes()
}

func (rt RecoveryTimeStamp) IsZeroValue() bool {
	return rt.Value == 0
}

func (rt RecoveryTimeStamp) SetHeader(ieHeader Header) InformationElement {
	rt.Header = ieHeader
	return rt
}

func DeserializeRecoveryTimeStamp(ieValue []byte) (RecoveryTimeStamp, error) {

	if len(ieValue) < 4 {
		return RecoveryTimeStamp{}, fmt.Errorf("invalid length for RecoveryTimeStamp value: expected at least 4 bytes, got %d", len(ieValue))
	}

	rt := RecoveryTimeStamp{
		Value: int64(binary.BigEndian.Uint32(ieValue)),
	}
	return rt, nil
}
