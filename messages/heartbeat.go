package messages

import (
	"time"
)

type RecoveryTimeStamp time.Time

type HeartbeatRequest struct {
	RecoveryTimeStamp RecoveryTimeStamp
}

type HeartbeatResponse struct {
	RecoveryTimeStamp RecoveryTimeStamp
}
