# PFCP

[![GoDoc](https://godoc.org/github.com/dot-5g/pfcp?status.svg)](https://godoc.org/github.com/dot-5g/pfcp)


A Go library for using the PFCP protocol in 5G networks as defined in the [ETSI TS 29.244 specification](https://www.etsi.org/deliver/etsi_ts/129200_129299/129244/16.04.00_60/ts_129244v160400p.pdf). 

## Usage

### Client

```go
package main

import (
	"log"
	"time"

	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/messages"
)

func main() {
	pfcpClient := client.New("1.2.3.4:8805")
	recoveryTimeStamp := ie.NewRecoveryTimeStamp(time.Now())
	sequenceNumber := uint32(21)
	heartbeatRequestMsg := messages.HeartbeatRequest{
		RecoveryTimeStamp: recoveryTimeStamp,
	}

	err := pfcpClient.SendHeartbeatRequest(heartbeatRequestMsg, sequenceNumber)
	if err != nil {
		log.Fatalf("SendHeartbeatRequest failed: %v", err)
	}
}
```

### Server


```go
package main

import (
	"fmt"

	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/server"
)

func main() {
	pfcpServer := server.New("localhost:8805")
	pfcpServer.HeartbeatRequest(HandleHeartbeatRequest)
	pfcpServer.Run()
}

func HandleHeartbeatRequest(sequenceNumber uint32, msg messages.HeartbeatRequest) {
	fmt.Printf("Received Heartbeat Request - Recovery TimeStamp: %v", msg.RecoveryTimeStamp)
}

```

## Procedures

### Node

- [x] Heartbeat
- [ ] Load Control (Optional)
- [ ] Overload Control (Optional)
- [ ] PFCP PFD Management (Optional)
- [x] PFCP Association Setup
- [x] PFCP Association Update
- [x] PFCP Association Release
- [x] PFCP Node Report

### Session

- [ ] PFCP Session Establishment
- [ ] PFCP Session Modification
- [ ] PFCP Session Deletion
- [ ] PFCP Session Report
