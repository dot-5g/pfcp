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
