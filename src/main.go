package main

import (
	"D7024E/cli"
	"D7024E/network/server"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/kademlia"
	"D7024E/node/kademlia/algorithms"
	"net"
	"time"
)

func main() {
	kademlia.GetInstance()
	go server.UDPListener(net.ParseIP(kademlia.GetInstance().Me.Address), 4001)
	algorithms.AddContact(contact.Contact{ID: id.NewKademliaID("172.21.0.2"), Address: "172.21.0.2"})
	time.Sleep(5 * time.Second)
	if contact.GetInstance().Address == "172.21.0.2" {
		contact.GetInstance().ID = id.NewKademliaID("172.21.0.2")
		time.Sleep(10 * time.Second)
	} else {
		algorithms.NodeLookup(*id.NewKademliaID("172.21.0.2"))
	}
	cli.CliListener()
}
