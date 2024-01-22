package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

const ntpEpochOffset = 2208988800 // Offset between Unix and NTP epoch (seconds)

type RecoveryTimeStamp struct {
	Value int64 // Seconds since 1900
}

func NewRecoveryTimeStamp(value time.Time) (RecoveryTimeStamp, error) {

	return RecoveryTimeStamp{
		Value: value.Unix() + ntpEpochOffset,
	}, nil
}

func (rt RecoveryTimeStamp) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 5 to 8: Value
	binary.Write(buf, binary.BigEndian, uint32(rt.Value))

	return buf.Bytes()
}

func (rt RecoveryTimeStamp) GetType() IEType {
	return RecoveryTimeStampIEType
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
