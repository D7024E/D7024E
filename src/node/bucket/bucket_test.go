package bucket

import (
	"D7024E/node/contact"
	"D7024E/node/id"
	"testing"
)

func TestNewBucket(t *testing.T) {
	bucket := newBucket()
	if bucket == nil {
		t.FailNow()
	}
}

// Validate contact to be added is in fact added.
func TestAddContact(t *testing.T) {
	bucket := newBucket()
	newContact := generateContact()
	bucket.AddContact(newContact)
	if bucket.list.Front().Value.(contact.Contact) != newContact {
		t.FailNow()
	}
}

// Tries to add a new contact to an already full bucket which will return bool=false
func TestAddContactBucketFullNewContactNotAdded(t *testing.T) {
	bucket := newBucket()
	for i := 0; i < 20; i++ {
		contact := generateContact()
		bucket.AddContact(contact)
	}

	contact := generateContact()

	_, isAdded := bucket.AddContact(contact)

	if isAdded == true {
		t.FailNow()
	}
}

// Test to add the same contact with the same id twice which will give length == 1 and isAdded == false
func TestAddSameContactTwice(t *testing.T) {
	bucket := newBucket()
	contact := generateContact()
	bucket.AddContact(contact)
	_, isAdded := bucket.AddContact(contact)

	if bucket.Len() != 1 && isAdded == true {
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

func TestRemoveContact(t *testing.T) {
	bucket := newBucket()
	contact := generateContact()
	bucket.AddContact(contact)
	bucket.RemoveContact(contact)
	if bucket.Len() != 0 {
		t.FailNow()
	}
}

// List of contacts with calculated lists is equal to the length of the bucket
func TestCalcDistance(t *testing.T) {
	bucket := newBucket()
	contact1 := generateContact()
	contact2 := generateContact()
	bucket.AddContact(contact1)
	bucket.AddContact(contact2)
	contacts := bucket.GetContactAndCalcDistance(contact1.ID)

	if len(contacts) != bucket.Len() {
		t.FailNow()
	}

}

func TestBucketLen(t *testing.T) {
	bucket := newBucket()
	len := bucket.Len()
	if len != 0 {
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
