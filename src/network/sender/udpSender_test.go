package sender

import (
	"D7024E/kademliaRPC/rpcmarshal"

	"net"
	"testing"
	"time"
)

func TestUDPSender(t *testing.T) {
	var testData []byte
	rpc := rpcmarshal.RPC{
		Cmd: "THIS IS A TEST",
	}
	rpcmarshal.RpcMarshal(rpc, &testData)
	buffer := make([]byte, 4096)

	go func() {
		dummyAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:4001")
		if err != nil {
			println("dummyAddr failed")
		}
		connection, err2 := net.ListenUDP("udp4", dummyAddr)
		if err2 != nil {
			println("There was an error:", err)
		}
		defer connection.Close()

		for {
			_, _, err3 := connection.ReadFromUDP(buffer)
			if err3 != nil {
				println("ReadFromUDP failed")
			}
		}
	}()

	destIP, err4 := net.ResolveIPAddr("udp4", "127.0.0.1")
	if err4 != nil {
		println("ResolveIPAddr failed")
	}
	UDPSender(destIP.IP, 4001, testData)
	time.Sleep(3 * time.Second)
	if buffer == nil {
		t.FailNow()
	}
}
