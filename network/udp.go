package network

import (
	"log"
	"net"
)

type Udp struct {
	address *net.UDPAddr
}

type UdpSender interface {
	Send(message []byte) error
}

func NewUdp(address string) (*Udp, error) {
	udpAddress, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Printf("Error resolving UDP address: %s\n", err)
		return nil, err
	}
	return &Udp{
		address: udpAddress,
	}, nil
}

func (udp *Udp) Send(message []byte) error {

	conn, err := net.DialUDP("udp", nil, udp.address)

	if err != nil {
		log.Printf("Error dialing UDP: %s\n", err)
		return err
	}
	defer conn.Close()

	_, err = conn.Write(message)
	if err != nil {
		log.Printf("Error sending message: %s\n", err)
		return err
	}

	return nil
}
