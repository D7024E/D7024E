package contact

import (
	"D7024E/node/id"
	"fmt"
	"testing"
)

// Generates a Contact instance.
func GenreateAContact() Contact {
	contact := Contact{
		ID:      id.NewRandomKademliaID(),
		Address: "THIS IS ADDRESS",
	}
	return contact
}

// Returns empty kademlia id.
func emptyKademliaID() *id.KademliaID {
	result := id.KademliaID{}
	for i := 0; i < id.IDLength; i++ {
		result[i] = 0
	}
	return &result
}

// Generates a Candiates instance with two Contacts in it.
func GenerateACandidateList() ContactCandidates {

	var contactList []Contact
	contact1 := GenreateAContact()
	contact2 := GenreateAContact()
	contactList = append(contactList, contact1)
	contactList = append(contactList, contact2)

	candidates := ContactCandidates{
		contacts: []Contact{},
	}

	candidates.Append(contactList)
	return candidates
}

func TestEqualsSuccess(t *testing.T) {
	contact := GenreateAContact()

	if !(contact.Equals(&contact) == true) {
		t.FailNow()
	}
}

func TestEqualsFail(t *testing.T) {
	contact1 := GenreateAContact()

	contact2 := GenreateAContact()

	if !(contact1.Equals(&contact2) == false) {
		t.FailNow()
	}
}

func TestContactToStringSuccess(t *testing.T) {
	contact := GenreateAContact()
	var conString interface{} = contact.String()
	if _, ok := conString.(string); ok {
		fmt.Println("do nothing")
	} else {
		t.FailNow()
	}
}

func TestSetAndGetDistanceSuccess(t *testing.T) {
	contact := GenreateAContact()
	contact.SetDistance(id.NewRandomKademliaID())

	fmt.Println(contact.GetDistance())
	if contact.GetDistance() == nil {
		t.FailNow()
	}
}

func TestSetAndGetDistanceFail(t *testing.T) {
	contact := GenreateAContact()

	fmt.Println(contact.GetDistance())
	if contact.GetDistance() != nil {
		t.FailNow()
	}
}

func TestContactAppend(t *testing.T) {

	candidates := GenerateACandidateList()

	fmt.Println(&candidates)

	if !(candidates.Len() != 0) {
		t.FailNow()
	}
}

func TestGetCandidates(t *testing.T) {
	candidates := GenerateACandidateList()

	candidates.GetContacts(2)
	if !(candidates.GetContacts(1) != nil) {
		t.FailNow()
	}
}

func TestCandidatesLen(t *testing.T) {
	// GenerateACandidateList generates a list of length 2.
	candidates := GenerateACandidateList()
	if !(candidates.Len() == 2) {
		t.FailNow()
	}
}

func TestCandidatesLess(t *testing.T) {
	candidates := GenerateACandidateList()
	dist := emptyKademliaID()
	candidates.contacts[0].SetDistance(dist)
	dist[id.IDLength-1] += 1
	candidates.contacts[1].SetDistance(dist)
	if candidates.Less(0, 1) {
		t.FailNow()
	}
}
