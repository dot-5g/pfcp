// Package network contains the network layer of the PFCP protocol.
package network

import (
	"log"
	"net"
)

type UDP struct {
	address *net.UDPAddr
}

type UDPSender interface {
	Send(message []byte) error
}

func NewUDP(address string) (*UDP, error) {
	udpAddress, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Printf("Error resolving UDP address: %s\n", err)
		return nil, err
	}
	return &UDP{
		address: udpAddress,
	}, nil
}

func (udp *UDP) Send(message []byte) error {
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
