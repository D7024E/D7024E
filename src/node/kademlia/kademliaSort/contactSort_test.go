package kademliaSort

import (
	"D7024E/node/contact"
	"D7024E/node/id"
	"fmt"
	"testing"
)

func TestContactSort(t *testing.T) {
	var contacts []contact.Contact

	contact1 := contact.Contact{
		ID:      id.NewRandomKademliaID(),
		Address: "Erik",
	}

	contact2 := contact.Contact{
		ID:      id.NewRandomKademliaID(),
		Address: "Dennis",
	}

	contact3 := contact.Contact{
		ID:      id.NewRandomKademliaID(),
		Address: "Anders",
	}

	contact1.SetDistance(id.NewRandomKademliaID())
	contact2.SetDistance(id.NewRandomKademliaID())
	contact3.SetDistance(id.NewRandomKademliaID())

	contacts = append(contacts, contact2)
	contacts = append(contacts, contact3)
	contacts = append(contacts, contact1)

	fmt.Println(contacts)

	sorted := SortContacts(contacts)

	fmt.Println(sorted)

	t.FailNow()

}
