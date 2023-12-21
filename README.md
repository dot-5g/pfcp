# PFCP

[![GoDoc](https://godoc.org/github.com/dot-5g/pfcp?status.svg)](https://godoc.org/github.com/dot-5g/pfcp)


A Go library for using the PFCP protocol in 5G networks as defined in the [ETSI TS 29.244 specification](https://www.etsi.org/deliver/etsi_ts/129200_129299/129244/16.04.00_60/ts_129244v160400p.pdf). 

## Usage

### Client

```go
package main

import (
	"log"

	"github.com/dot-5g/pfcp/client"
)

func main() {
	pfcpClient := client.New("1.2.3.4:8805")
	_, err := pfcpClient.SendHeartbeatRequest()
	if err != nil {
		log.Fatalf("SendHeartbeatRequest failed: %v", err)
	}
}
```

### Server


```go
package main

import (
	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/server"
)

func main() {
	pfcpServer := server.New("localhost:8805")
	pfcpServer.HeartbeatRequest(HandleHeartbeatRequest)
	pfcpServer.HeartbeatResponse(HandleHeartbeatResponse)
	pfcpServer.Run()
}

func HandleHeartbeatRequest(h *messages.HeartbeatRequest) {
	// Do something
}

func HandleHeartbeatResponse(h *messages.HeartbeatResponse) {
	// Do something
}

```
