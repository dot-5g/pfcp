# PFCP

A Go library for using the PFCP protocol in 5G networks.

## Usage

### Client

```go
pfcpClient := pfcp.New("1.2.3.4:8805")
err := pfcpClient.SendHeartbeatRequest()
if err != nil {
    t.Errorf("SendHeartbeatRequest failed: %v", err)
}
```
