package network

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

type UdpServer struct {
	conn    *net.UDPConn
	closeCh chan struct{}
	wg      sync.WaitGroup
	Handler func([]byte)
}

func (udpServer *UdpServer) SetHandler(handler func([]byte)) {
	udpServer.Handler = handler
}

func NewUdpServer() *UdpServer {
	return &UdpServer{
		closeCh: make(chan struct{}),
	}
}

func (udpServer *UdpServer) Run(address string) error {
	var err error
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return fmt.Errorf("failed to resolve UDP address: %w", err)
	}

	udpServer.conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on UDP address: %w", err)
	}

	udpServer.wg.Add(1)
	go udpServer.listen()
	log.Printf("Listening on %s\n", addr)
	return nil
}

func (udpServer *UdpServer) listen() {
	defer udpServer.wg.Done()

	for {
		select {
		case <-udpServer.closeCh:
			return
		default:
			buffer := make([]byte, 1024)
			udpServer.conn.SetReadDeadline(time.Now().Add(time.Millisecond * 500)) // Set a short deadline
			length, _, err := udpServer.conn.ReadFrom(buffer)

			if err != nil {
				if !strings.Contains(err.Error(), "use of closed network connection") {
					log.Printf("Error reading from UDP: %v", err)
				}
				continue
			}

			if udpServer.Handler != nil {
				go udpServer.Handler(buffer[:length])
			}
		}
	}
}

func (udpServer *UdpServer) Close() {
	close(udpServer.closeCh)
	udpServer.conn.Close()
	udpServer.wg.Wait()
}
