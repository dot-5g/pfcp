package main

import (
	"fmt"
	"time"

	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/messages"
)

func main() {
	pfcpClient := client.New("localhost:8805")
	nodeID, err := ie.NewNodeID("1.2.3.4")
	if err != nil {
		fmt.Printf("Error creating NodeID: %v", err)
	}
	recoveryTimeStamp, err := ie.NewRecoveryTimeStamp(time.Now())
	if err != nil {
		fmt.Printf("Error creating Recovery Time Stamp IE: %v", err)
	}
	message := messages.PFCPAssociationSetupRequest{
		NodeID:            nodeID,
		RecoveryTimeStamp: recoveryTimeStamp,
	}
	sequenceNumber := uint32(1)
	err = pfcpClient.SendPFCPAssociationSetupRequest(message, sequenceNumber)
	if err != nil {
		fmt.Printf("Error sending Heartbeat Request: %v", err)
	}
	fmt.Printf("Heartbeat Request sent successfully to %s.\n", pfcpClient.ServerAddress)
}
