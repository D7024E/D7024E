package main

import (
	"D7024E/log"
	"D7024E/network/server"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"net"
)

func main() {
	ip := net.IPv4(127, 0, 0, 1)
	port := 4001
	go server.UDPListener(ip, port)
	server.UDPSender(ip, port, "this is the message")

	rt := bucket.GetInstance()
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
