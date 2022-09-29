package main

import (
	"D7024E/cli"
	"D7024E/network/server"
	"D7024E/node"
	"D7024E/node/kademlia"
	"net"
)

func main() {
	kademlia.GetInstance()
	go server.UDPListener(net.ParseIP(kademlia.GetInstance().Me.Address), 4001)
	node.StartKademliaNode()
	cli.CliListener()
}
