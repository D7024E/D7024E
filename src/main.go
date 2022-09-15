package main

import (
	"D7024E/network"
	"time"
)

func main() {
	// network.Start()
	go network.Receiver()
	time.Sleep(1 * time.Second)
	network.Sender("172.0.0.1", "this is udp test")
	for {

	}
}
