package ipc

import (
	"testing"
	"time"
)

type EchoServer struct {
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func (server *EchoServer) Handle(method, request string) *Response {
	return &Response{"Echo:", request}
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1,_ := client1.Call("abc", "From Client1")
	resp2,_ := client2.Call("abc", "From Client2")

	if resp1.Code + resp1.Body != "Echo:From Client1" || resp2.Code + resp2.Body != "Echo:From Client2" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2:", resp2)
	}

	client1.Close()
	client2.Close()

	time.Sleep(time.Second)
}
