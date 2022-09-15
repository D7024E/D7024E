package main

import "D7024E/network"

func main() {
	// network.Start()
	go network.Receiver()
	network.Sender("172.0.0.1", "this is udp test")
	for {

	}
}
