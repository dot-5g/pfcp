package main

import (
	"log"
	"time"

	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/server"
)

func main() {
	pfcpClient := client.New("1.2.3.4:8805")
	recoveryTimeStamp := ie.NewRecoveryTimeStamp(time.Now())
	sequenceNumber := uint32(21)

	_, err := pfcpClient.SendHeartbeatRequest(recoveryTimeStamp, sequenceNumber)
	if err != nil {
		log.Fatalf("SendHeartbeatRequest failed: %v", err)
	}
}

func RunServer() {
	pfcpServer := server.New("localhost:8805")
	pfcpServer.HeartbeatRequest(HandleHeartbeatRequest)
	pfcpServer.HeartbeatResponse(HandleHeartbeatResponse)
	pfcpServer.Run()
}

func HandleHeartbeatRequest(sequenceNumber uint32, recoveryTimeStamp ie.RecoveryTimeStamp) {

}

func HandleHeartbeatResponse(sequenceNumber uint32, recoveryTimeStamp ie.RecoveryTimeStamp) {

}
