package server_test

import (
	"testing"
	"time"

	"github.com/dot-5g/pfcp/server"
)

func TestServer(t *testing.T) {
	// t.Run("TestMoreThanOneServer", MoreThanOneServer)
	// t.Run("TestServerClosedNoError", ServerClosedNoError)
}

func MoreThanOneServer(t *testing.T) {
	address := "127.0.0.1:8805"

	server1 := server.New(address)
	server2 := server.New(address)
	go server1.Run()
	defer server1.Close()

	time.Sleep(time.Second)

	err2 := server2.Run()
	defer server2.Close()

	if err2 == nil {
		t.Errorf("Expected error to be returned when starting server2 on the same address")
	}
}

func ServerClosedNoError(t *testing.T) {
	server := server.New("127.0.0.1:8805")
	go server.Run()
	time.Sleep(time.Second)
	server.Close()

	go func() {
		err := server.Run()
		if err != nil {
			t.Errorf("Expected no error to be returned")
		}
	}()
	defer server.Close()

}
