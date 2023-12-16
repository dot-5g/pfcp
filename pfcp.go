package pfcp

import (
	"log"

	"github.com/dot-5g/pfcp/network"
)

type Pfcp struct {
	ServerAddr string
	Udp        network.UdpSender
}

func NewPfcp(serverAddr string) *Pfcp {

	udpClient, err := network.NewUdp(serverAddr)
	if err != nil {
		log.Printf("Failed to initialize UDP client: %v\n", err)
		return nil
	}

	return &Pfcp{ServerAddr: serverAddr, Udp: udpClient}
}
