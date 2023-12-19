# PFCP

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
	err := pfcpClient.SendHeartbeatRequest()
	if err != nil {
		log.Fatalf("SendHeartbeatRequest failed: %v", err)
	}
}
```

### Server


```go
func main() {
    pfcpServer := server.New()
    pfcpServer.HeartbeatRequest(HandleHeartbeatRequest)
    pfcpServer.HeartbeatResponse(HandleHeartbeatRequest)
    pfcpServer.Run("localhost:8805")
}

func HandleHeartbeatRequest(h *heartbeatRequest) {
    // Do something
}

func HandleHeartbeatResponse(h *heartbeatRequest) {
    // Do something
}
```