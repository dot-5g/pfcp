package network

import (
	"log"
	"net"
	"sync"
)

type UdpServer struct {
	conn    *net.UDPConn
	closeCh chan struct{}
	wg      sync.WaitGroup
	Handler func([]byte, net.Addr)
}

func (udpServer *UdpServer) SetHandler(handler func([]byte, net.Addr)) {
	udpServer.Handler = handler
}

func NewUdpServer() *UdpServer {
	return &UdpServer{
		closeCh: make(chan struct{}),
	}
}

func (udpServer *UdpServer) Run(address string) error {
	var err error
	addr, _ := net.ResolveUDPAddr("udp", address)
	udpServer.conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}

	udpServer.wg.Add(1)
	go func() {
		defer udpServer.wg.Done()
		for {
			select {
			case <-udpServer.closeCh:
				return
			default:
				buffer := make([]byte, 1024)
				length, remoteAddr, err := udpServer.conn.ReadFrom(buffer)
				if err != nil {
					log.Printf("Error reading from UDP: %v", err)
					continue
				}

				if udpServer.Handler != nil {
					go udpServer.Handler(buffer[:length], remoteAddr)
				}
			}
		}
	}()

	log.Printf("Listening on %s\n", addr)
	return nil
}

func (udpServer *UdpServer) Close() {
	close(udpServer.closeCh)
	udpServer.conn.Close()
	udpServer.wg.Wait()
}
