package network

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type UDPServer struct {
	conn    *net.UDPConn
	closeCh chan struct{}
	Handler func(net.Addr, []byte)
}

func (udpServer *UDPServer) SetHandler(handler func(net.Addr, []byte)) {
	udpServer.Handler = handler
}

func NewUDPServer() *UDPServer {
	return &UDPServer{
		closeCh: make(chan struct{}),
	}
}

func (udpServer *UDPServer) Run(address string) error {
	var err error
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return fmt.Errorf("failed to resolve UDP address: %w", err)
	}

	udpServer.conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on UDP address: %w", err)
	}

	log.Printf("Running PFCP server on on %s\n", addr)

	err = udpServer.listen()

	return err
}

func (udpServer *UDPServer) listen() error {
	for {
		select {
		case <-udpServer.closeCh:
			return nil
		default:
			buffer := make([]byte, 1024)
			length, remoteAddress, err := udpServer.conn.ReadFrom(buffer)
			if err != nil {
				if !strings.Contains(err.Error(), "use of closed network connection") {
					return fmt.Errorf("failed to read from UDP connection: %w", err)
				}
				continue
			}
			if udpServer.Handler != nil {
				udpServer.Handler(remoteAddress, buffer[:length])
			}
		}
	}
}

func (udpServer *UDPServer) Close() error {
	var err error
	select {
	case <-udpServer.closeCh:
	default:
		close(udpServer.closeCh)
	}

	if udpServer.conn != nil {
		err = udpServer.conn.Close()
		log.Printf("Closed PFCP server\n")
	}

	return err
}
