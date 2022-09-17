package main

import (
	"D7024E/contact"
	"D7024E/id"
	"D7024E/log"
	"D7024E/node"
	"D7024E/node/stored"
)

func main() {
	// ip := net.IPv4(127, 0, 0, 1)
	// port := 4001
	// go network.UDPListener(ip, port)
	// network.UDPSender(ip, port, "this is the message")
	rt := node.GetInstance()
	rt.SetMe(contact.Contact{ID: id.NewRandomKademliaID(), Address: "this is address"})
	log.INFO("%v", rt.GetMe())
	val := stored.GetInstance()
	i := *id.NewRandomKademliaID()
	val.Store([]stored.Value{{Data: "this is the data", ID: i}})
	res, err := val.FindValue(i)
	if err != nil {
		log.ERROR("%v", err)
	} else {
		log.INFO("%v", res)
	}
}
