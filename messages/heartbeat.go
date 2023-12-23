package messages

import (
	"github.com/dot-5g/pfcp/information_elements"
)

type HeartbeatRequest struct {
	RecoveryTimeStamp information_elements.RecoveryTimeStamp
	SequenceNumber    uint32
}

type HeartbeatResponse struct {
	RecoveryTimeStamp information_elements.RecoveryTimeStamp
	SequenceNumber    uint32
}
