package network

import (
	"log"
	"net"
)

type UdpServer struct {
	Handler func([]byte, net.Addr)
}

func NewUdpServer() *UdpServer {
	return &UdpServer{}
}

func (s *UdpServer) Run(address string) error {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return err
	}
	log.Printf("Listening on %s\n", addr)

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}
	log.Printf("Listening on %s\n", addr)

	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		length, remoteAddr, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Printf("Error reading from UDP: %v", err)
			continue
		}

		if s.Handler != nil {
			go s.Handler(buffer[:length], remoteAddr)
		}
	}
}
