package main

import (
	"log"

	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/server"
)

func main() {
	pfcpClient := client.New("1.2.3.4:8805")
	_, err := pfcpClient.SendHeartbeatRequest()
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

func HandleHeartbeatRequest(h messages.HeartbeatRequest) {

}

func HandleHeartbeatResponse(h messages.HeartbeatResponse) {

}
