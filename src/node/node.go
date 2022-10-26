package node

import (
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/kademlia/algorithms"
	"fmt"
	"strings"
	"time"
)

func StartKademliaNode() {
	kClosest := algorithms.NodeLookup(*contact.GetInstance().ID)
	for _, c := range kClosest {
		bucket.GetInstance().AddContact(c)
	}

	go func() {
		time.Sleep(time.Minute)
		kClosest := bucket.GetInstance().FindClosestContacts(contact.GetInstance().ID, 100)
		kAddress := []string{}
		for _, c := range kClosest {
			kAddress = append(kAddress, c.Address)
		}
		res := "\n" + strings.Join(kAddress, "          \n")
		fmt.Println("[NODE] -", contact.GetInstance().Address, "ROUTING TABLE: ["+res+"]")
	}()
}
