package server_test

import (
	"testing"

	"github.com/dot-5g/pfcp/server"
)

func TestServer(t *testing.T) {
	t.Run("TestMoreThanOneServer", MoreThanOneServer)
	t.Run("TestServerClosedNoError", ServerClosedNoError)
}

func MoreThanOneServer(t *testing.T) {
	server1 := server.New("127.0.0.1:8805")
	server2 := server.New("127.0.0.1:8805")

	err1 := server1.Run()
	err2 := server2.Run()
	defer server1.Close()
	defer server2.Close()

	if err1 != nil {
		t.Errorf("Expected no error to be returned")
	}

	if err2 == nil {
		t.Errorf("Expected error to be returned")
	}
}

func ServerClosedNoError(t *testing.T) {
	server := server.New("127.0.0.1:8805")
	server.Run()
	server.Close()

	err := server.Run()
	defer server.Close()

	if err != nil {
		t.Errorf("Expected no error to be returned")
	}

}
