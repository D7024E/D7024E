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

func TestGetInstance(t *testing.T) {
	rtable := GetInstance()

	if rtable == nil {
		t.FailNow()
	}
}
