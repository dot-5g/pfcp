package main

import (
	"fmt"

	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/server"
)

func main() {
	RunServer()
}

func RunServer() {
	pfcpServer := server.New("localhost:8805")
	pfcpServer.HeartbeatRequest(HandleHeartbeatRequest)
	pfcpServer.HeartbeatResponse(HandleHeartbeatResponse)
	go pfcpServer.Run()
}

func HandleHeartbeatRequest(sequenceNumber uint32, msg messages.HeartbeatRequest) {
	fmt.Printf("Received Heartbeat Request - Recovery TimeStamp: %v", msg.RecoveryTimeStamp)
}

func HandleHeartbeatResponse(sequenceNumber uint32, msg messages.HeartbeatResponse) {
	fmt.Printf("Received Heartbeat Response - Recovery TimeStamp: %v", msg.RecoveryTimeStamp)
}
