package ie

import (
	"encoding/binary"
	"fmt"
	"time"
)

const ntpEpochOffset = 2208988800 // Offset between Unix and NTP epoch (seconds)

type RecoveryTimeStamp struct {
	IEType uint16
	Length uint16
	Value  int64 // Seconds since 1900
}

func NewRecoveryTimeStamp(value time.Time) RecoveryTimeStamp {
	return RecoveryTimeStamp{
		IEType: uint16(RecoveryTimeStampIEType),
		Length: 4,
		Value:  value.Unix() + ntpEpochOffset,
	}
}

func (rt RecoveryTimeStamp) Serialize() []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint16(bytes[0:2], uint16(rt.IEType))
	binary.BigEndian.PutUint16(bytes[2:4], uint16(rt.Length))
	binary.BigEndian.PutUint32(bytes[4:8], uint32(rt.Value))
	return bytes
}

func (rt RecoveryTimeStamp) IsZeroValue() bool {
	return rt.Length == 0
}

func DeserializeRecoveryTimeStamp(ieType uint16, ieLength uint16, ieValue []byte) (RecoveryTimeStamp, error) {
	var rt RecoveryTimeStamp

	if ieType != uint16(RecoveryTimeStampIEType) {
		return rt, fmt.Errorf("invalid IE type for RecoveryTimeStamp: expected %d, got %d", RecoveryTimeStampIEType, ieType)
	}
	if ieLength != 4 {
		return rt, fmt.Errorf("invalid length for RecoveryTimeStamp: expected 4, got %d", ieLength)
	}

	if len(ieValue) < 4 {
		return rt, fmt.Errorf("invalid length for RecoveryTimeStamp value: expected at least 4 bytes, got %d", len(ieValue))
	}

	rt.IEType = ieType
	rt.Length = ieLength
	rt.Value = int64(binary.BigEndian.Uint32(ieValue))

	return rt, nil
}
