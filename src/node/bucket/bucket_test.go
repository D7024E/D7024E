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

func TestBucketLen(t *testing.T) {
	bucket := newBucket()
	len := bucket.Len()
	if len != 0 {
		t.FailNow()
	}
}
