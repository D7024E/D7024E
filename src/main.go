package main

import (
	"D7024E/network"
	"net"
)

func main() {
	ip := net.IPv4(127, 0, 0, 1)
	port := 4001
	go network.Receiver(ip, port)
	// time.Sleep(2 * time.Second)
	network.Sender(ip, port, "this is the message")
}
