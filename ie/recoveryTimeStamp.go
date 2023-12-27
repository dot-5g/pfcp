package ie

import (
	"encoding/binary"
	"time"
)

const ntpEpochOffset = 2208988800 // Offset between Unix and NTP epoch (seconds)

type RecoveryTimeStamp struct {
	IEtype uint16
	Length uint16
	Value  int64 // Seconds since 1900
}

func NewRecoveryTimeStamp(value time.Time) RecoveryTimeStamp {
	return RecoveryTimeStamp{
		IEtype: 96,
		Length: 4,
		Value:  value.Unix() + ntpEpochOffset,
	}
}

func (rt RecoveryTimeStamp) Serialize() []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint16(bytes[0:2], uint16(rt.IEtype))
	binary.BigEndian.PutUint16(bytes[2:4], uint16(rt.Length))
	binary.BigEndian.PutUint32(bytes[4:8], uint32(rt.Value))
	return bytes
}

func DeserializeRecoveryTimeStamp(ieType uint16, ieLength uint16, ieValue []byte) RecoveryTimeStamp {
	return RecoveryTimeStamp{
		IEtype: ieType,
		Length: ieLength,
		Value:  int64(binary.BigEndian.Uint32(ieValue)),
	}
}
