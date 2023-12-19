package server_test

import (
	"fmt"
	"testing"

	"github.com/dot-5g/pfcp/server"
)

func HandleHeartbeatRequest(h server.HeartbeatRequest) {
	fmt.Printf("We here boys")
}
func TestGivenHandleHeartbeatRequestWhenRunThenHeartbeatRequestHandled(t *testing.T) {
	pfcpServer := server.New()
	pfcpServer.HeartbeatRequest(HandleHeartbeatRequest)

	pfcpServer.Run("127.0.0.1:8805")
}
