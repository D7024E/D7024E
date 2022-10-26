package node

import (
	"D7024E/log"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/kademlia/algorithms"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func StartKademliaNode() {
	log.INFO("[NODE] - CONNECTED - [KADEMLIA NETWORK]")
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	kClosest := algorithms.NodeLookup(*contact.GetInstance().ID)
	kAddress := []string{}
	for _, c := range kClosest {
		kAddress = append(kAddress, c.Address)
	}
	res := "\n" + strings.Join(kAddress, "          \n")
	fmt.Println("[NODE] - TABLE: [" + res + "]")
	fmt.Println(contact.GetInstance().Address)

	go func() {
		time.Sleep(time.Minute)
		kClosest := bucket.GetInstance().FindClosestContacts(id.NewRandomKademliaID(), 100)
		kAddress := []string{}
		for _, c := range kClosest {
			kAddress = append(kAddress, c.Address)
		}
		res := "\n" + strings.Join(kAddress, "          \n")
		fmt.Println("[NODE] -", contact.GetInstance().Address, "TABLE: ["+res+"]")
	}()
}
