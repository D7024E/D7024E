package bucket

import (
	"D7024E/node/contact"
	"D7024E/node/id"
	"testing"
)

// Validate contact to be added is in fact added.
func TestAddContact(t *testing.T) {
	bucket := newBucket()
	newContact := generateContact()
	bucket.AddContact(newContact)
	if bucket.list.Front().Value.(contact.Contact) != newContact {
		t.FailNow()
	}
}

// Validate that the test does not give false positives.
func TestAddContactFail(t *testing.T) {
	bucket := newBucket()
	newContact := generateContact()
	bucket.AddContact(newContact)
	if bucket.list.Front().Value.(contact.Contact) == generateContact() {
		t.FailNow()
	}
}

func generateContact() contact.Contact {
	newContact := contact.Contact{
		ID:      id.NewRandomKademliaID(),
		Address: id.NewRandomKademliaID().String(),
	}
	return newContact
}
