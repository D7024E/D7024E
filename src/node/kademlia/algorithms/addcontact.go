package algorithms

import (
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/kademlia"
	"fmt"
)

// // Attempt to add a contact to its bucket.
func AddContact(newContact contact.Contact) {
	rt := bucket.GetInstance()
	_, res := rt.AddContact(newContact)
	if res {
		return
	} else {
		// rt.RemoveContact(head)
		// AddContact(newContact)
	}
}

// Attempt to add "n" random contacts to the routing table.
func TestAddContact(n int) {
	kademlia.GetInstance()
	for i := 0; i < n; i++ {
		var newContact = contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: id.NewRandomKademliaID().String(),
		}
		fmt.Println(newContact)

		AddContact(newContact)
	}
}

// Currently fails when retrieving the bucket id since it need "Me".
