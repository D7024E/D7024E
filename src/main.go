package main

import (
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/log"
	"D7024E/network/server"
	"D7024E/node/contact"
	"D7024E/node/kademlia"
	"math/rand"
	"net"
	"time"
)

func main() {
	// ip := net.IPv4(127, 0, 0, 1)
	// port := 4001
	// go server.UDPListener(ip, port)
	// sender.UDPSender(ip, port, "this is the message")

	// rt := bucket.GetInstance()
	// rt.SetMe(contact.Contact{ID: id.NewRandomKademliaID(), Address: "this is address"})
	// log.INFO("%v", rt.GetMe())
	// val := stored.GetInstance()
	// i := *id.NewRandomKademliaID()
	// val.Store([]stored.Value{{Data: "this is the data", ID: i}})
	// res, err := val.FindValue(i)
	// if err != nil {
	// 	log.ERROR("%v", err)
	// } else {
	// 	log.INFO("%v", res)
	// }
	go server.UDPListener(net.ParseIP(kademlia.GetInstance().Me.Address), 4001)
	min := 2
	max := 6
	time.Sleep(time.Duration(rand.Intn(max-min)+min) * time.Second)
	if kademlia.GetInstance().Me.Address == "172.21.0.2" {
		for {
		}
	}
	res := rpc.Ping(kademlia.GetInstance().Me, contact.Contact{Address: "172.21.0.2"})
	log.INFO("RES: %v", res)
	for {

	}
}
