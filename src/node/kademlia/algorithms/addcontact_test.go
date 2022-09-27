package algorithms

import (
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"fmt"
	"testing"
)

func TestAddContact(t *testing.T) {
	rt := bucket.GetInstance()
	for i := 0; i < 100; i++ {
		var newContact = contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: id.NewRandomKademliaID().String(),
		}
		fmt.Println(newContact)
		AddContact(newContact)
		foundContact := rt.FindClosestContacts(newContact.ID, 1)
		if foundContact[0] != newContact {
			t.FailNow()
		}
	}
}
