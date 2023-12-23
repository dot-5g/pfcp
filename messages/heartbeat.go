package messages

import (
	"github.com/dot-5g/pfcp/ie"
)

type HeartbeatRequest struct {
	RecoveryTimeStamp ie.RecoveryTimeStamp
	SequenceNumber    uint32
}

type HeartbeatResponse struct {
	RecoveryTimeStamp ie.RecoveryTimeStamp
	SequenceNumber    uint32
}
