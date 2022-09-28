package bucket

import (
	"D7024E/node/contact"
	"D7024E/node/id"
	"testing"
)

func GenreateAContact() contact.Contact {
	contact := contact.Contact{
		ID:      id.NewRandomKademliaID(),
		Address: "THIS IS ADDRESS",
	}
	return contact
}

func TestNewBucket(t *testing.T) {
	bucket := newBucket()
	if bucket == nil {
		t.FailNow()
	}
}

// Adds a single Contact
func TestAddContact(t *testing.T) {
	bucket := newBucket()
	contact := GenreateAContact()
	bucket.AddContact(contact)

	if bucket.Len() == 0 {
		t.FailNow()
	}
}

// Tries to add a new contact to an already full bucket which will return bool=false
func TestAddContactBucketFullNewContactNotAdded(t *testing.T) {
	bucket := newBucket()
	for i := 0; i < 20; i++ {
		contact := GenreateAContact()
		bucket.AddContact(contact)
	}

	contact := GenreateAContact()

	_, isAdded := bucket.AddContact(contact)

	if isAdded == true {
		t.FailNow()
	}

}

// Test to add the same contact with the same id twice which will give length == 1 and isAdded == false
func TestAddSameContactTwice(t *testing.T) {
	bucket := newBucket()
	contact := GenreateAContact()
	bucket.AddContact(contact)
	_, isAdded := bucket.AddContact(contact)

	if bucket.Len() != 1 && isAdded == true {
		t.FailNow()
	}
}

func TestRemoveContact(t *testing.T) {
	bucket := newBucket()
	contact := GenreateAContact()
	bucket.AddContact(contact)
	bucket.RemoveContact(contact)
	if bucket.Len() != 0 {
		t.FailNow()
	}
}

// List of contacts with calculated lists is equal to the length of the bucket
func TestCalcDistance(t *testing.T) {
	bucket := newBucket()
	contact1 := GenreateAContact()
	contact2 := GenreateAContact()
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
