package messages

import (
	"encoding/binary"
	"time"
)

type RecoveryTimeStamp struct {
	Type   int
	Length int
	Value  int64 // Seconds since 1900
}

type HeartbeatRequest struct {
	RecoveryTimeStamp RecoveryTimeStamp
	SequenceNumber    uint32
}

type HeartbeatResponse struct {
	RecoveryTimeStamp RecoveryTimeStamp
	SequenceNumber    uint32
}

func NewRecoveryTimeStamp(value time.Time) RecoveryTimeStamp {
	return RecoveryTimeStamp{
		Type:   96,
		Length: 8,
		Value:  value.Unix() + ntpEpochOffset,
	}
}

const ntpEpochOffset = 2208988800 // Offset between Unix and NTP epoch (seconds)

func (rt RecoveryTimeStamp) ToBytes() []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint16(bytes[0:2], uint16(rt.Type))
	binary.BigEndian.PutUint16(bytes[2:4], uint16(rt.Length))
	binary.BigEndian.PutUint32(bytes[4:8], uint32(rt.Value))
	return bytes
}

func FromBytes(b []byte) RecoveryTimeStamp {
	return RecoveryTimeStamp{
		Type:   int(binary.BigEndian.Uint16(b[0:2])),
		Length: int(binary.BigEndian.Uint16(b[2:4])),
		Value:  int64(binary.BigEndian.Uint32(b[4:8])),
	}
}
