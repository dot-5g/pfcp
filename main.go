package main

import (
	"log"
	"os"

	"github.com/dot-5g/pfcp/pfcp"
)

func main() {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	pfcpClient := pfcp.NewPfcp("1.2.3.4:8805")
	err = pfcpClient.SendHeartbeatRequest()
	if err != nil {
		log.Println(err)
		return
	}
}
