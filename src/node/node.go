package node

import (
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/log"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/kademlia/algorithms"
	"strings"
	"time"
)

func StartKademliaNode() {
	algorithms.AddContact(contact.Contact{ID: id.NewKademliaID("172.21.0.2"), Address: "172.21.0.2"}, rpc.Ping)
	time.Sleep(5 * time.Second)
	if contact.GetInstance().Address == "172.21.0.2" {
		contact.GetInstance().ID = id.NewKademliaID("172.21.0.2")
	} else {
		kClosest := algorithms.NodeLookup(*id.NewKademliaID("172.21.0.2"))
		kID := []string{}
		for _, c := range kClosest {
			kID = append(kID, c.ID.String())
		}
		res := "\n" + strings.Join(kID, "          \n")
		log.INFO("KCLOSEST NODES - [%v]", res)
	}
	log.INFO("CONNECTED - [KADEMLIA NETWORK]")
}
