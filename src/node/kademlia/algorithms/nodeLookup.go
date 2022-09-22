package kademlia

import (
	"D7024E/config"
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/log"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
)

func NodeLookup(destNode id.KademliaID, batch []contact.Contact) (closest []contact.Contact) {
	rt := bucket.GetInstance()
	//node := rt.Me

	if len(batch) == 0 {
		batch = rt.FindClosestContacts(&destNode, config.Alpha)
	}
	var newBatch [][]contact.Contact
	// For each of the alpha nodes in "batch", send a findNode RPC and append the result to "newBatch"
	for i := 0; i < len(batch); i++ {
		var kN []contact.Contact
		kN, err := rpc.FindNode(batch[i], destNode)
		if err != nil {
			log.ERROR("%v", err)
		} else {
			newBatch = append(newBatch, kN)
		}
	}

	// Convert the contact batch into a single slice.
	batch = mergeBatch(newBatch)

	//

	return
}

// Note that this merge do not take duplicates into account.
func mergeBatch(batch [][]contact.Contact) []contact.Contact {
	var mergedBatch []contact.Contact
	for i := 0; i < len(batch); i++ {
		mergedBatch = append(mergedBatch, batch[i]...)
	}
	return mergedBatch
}
