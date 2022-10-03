package bucket

import (
	"D7024E/node/contact"
	"D7024E/node/id"
	"testing"
)

// Validate instance consistency.
func TestGetInstance(t *testing.T) {
	instance := GetInstance()
	for i := 0; i < 100; i++ {
		newInstance := GetInstance()
		if newInstance != instance {
			t.FailNow()
		}
	}
}

// Verify that AddContact is called correctly.
func TestAddContactRouting(t *testing.T) {
	rt := GetInstance()
	newContact := generateContact()
	bucketIndex := rt.getBucketIndex(newContact.ID)
	bucket := rt.buckets[bucketIndex]
	rt.AddContact(newContact)
	if bucket.list.Front().Value.(contact.Contact) != newContact {
		t.FailNow()
	}
}

// Verify that duplicate contacts can't be added.
func TestAddContactRoutingDuplicates(t *testing.T) {
	rt := GetInstance()
	newContact := generateContact()
	bucketIndex := rt.getBucketIndex(newContact.ID)
	bucket := rt.buckets[bucketIndex]
	rt.AddContact(newContact)
	rt.AddContact(newContact)
	if bucket.Len() != 1 {
		t.FailNow()
	}
}

// Verify that a remove contact is called correctly.
func TestRemoveContactRouting(t *testing.T) {
	rt := GetInstance()
	newContact := generateContact()
	bucketIndex := rt.getBucketIndex(newContact.ID)
	bucket := rt.buckets[bucketIndex]
	rt.AddContact(newContact)
	rt.RemoveContact(newContact)
	if bucket.Len() != 0 {
		t.FailNow()
	}
}

// Verify that the closest returned contact is correct.
func TestFindClosestContactsKnownTarget(t *testing.T) {
	rt := GetInstance()
	keyContact := generateContact()
	rt.AddContact(keyContact)
	for i := 0; i < 100; i++ {
		rt.AddContact(generateContact())
	}
	res := rt.FindClosestContacts(keyContact.ID, 10)

	if res[0].ID != keyContact.ID {
		t.FailNow()
	}
}

// Verify that FindClosestContacts does not return more contacts than exist the routing table
func TestFindClosestContactsFewerContacts(t *testing.T) {
	rt := GetInstance()
	keyContact := generateContact()
	rt.AddContact(keyContact)
	for i := 0; i < 4; i++ {
		rt.AddContact(generateContact())
	}
	res := rt.FindClosestContacts(keyContact.ID, 10)
	if len(res) != 5 {
		t.FailNow()
	}
}

// Verify that FindClosestContacts doesn't break on a empty routing table.
func TestFindClosestContactsEmpty(t *testing.T) {
	rt := GetInstance()
	keyContact := generateContact()
	res := rt.FindClosestContacts(keyContact.ID, 10)
	if len(res) > 0 {
		t.FailNow()
	}
}

// Verify that the returned bucket index is consistent.
func TestGetBucketIndex(t *testing.T) {
	rt := GetInstance()
	keyContact := id.NewRandomKademliaID()
	bucketIndex := rt.getBucketIndex(keyContact)
	for i := 0; i < 100; i++ {
		newBucketIndex := rt.getBucketIndex(keyContact)
		if bucketIndex != newBucketIndex {
			t.FailNow()
		}
	}
}
