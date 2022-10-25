package node

import (
	"D7024E/log"
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
	if contact.GetInstance().Address == "172.21.0.2" {
		contact.GetInstance().ID = id.NewKademliaID("172.21.0.2")
	}
	log.INFO("[NODE] - CONNECTED - [KADEMLIA NETWORK]")
	go func() {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		kClosest := algorithms.NodeLookup(*id.NewKademliaID("172.21.0.2"))
		kID := []string{}
		for _, c := range kClosest {
			kID = append(kID, c.ID.String())
		}
		res := "\n" + strings.Join(kID, "          \n")
		fmt.Println("[NODE] - TABLE: [" + res + "]")
	}()
}
