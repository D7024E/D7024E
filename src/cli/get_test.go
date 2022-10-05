package cli

import (
	"D7024E/node/id"
	"D7024E/node/stored"
	"errors"
	"testing"
)

// Validate successfull get.
func TestGetSuccess(t *testing.T) {
	var res string = Get(id.NewRandomKademliaID().String(), NodeValueLookupMockSuccess)
	if res == "" {
		t.FailNow()
	}
}

// Validate behaviour on valid fail.
func TestGetValidFail(t *testing.T) {
	var res string = Get(id.NewRandomKademliaID().String(), NodeValueLookupMockValidFail)
	if res != "value not found" {
		t.FailNow()
	}
}

// Validate incorrect hash length rejection
func TestGetInvalidHash(t *testing.T) {
	var shortRes string = Get("Too short hash", NodeValueLookupMockSuccess)
	var longRes string = Get("Too long hash sadjnmaslsdasndoasndoasindoasindoas", NodeValueLookupMockSuccess)
	if shortRes != "invalid length str 2 kademlia" {
		t.FailNow()
	} else if longRes != "invalid length str 2 kademlia" {
		t.FailNow()
	}
}

func NodeValueLookupMockSuccess(valueID id.KademliaID) (stored.Value, error) {
	var mockValue = stored.Value{
		Data: "This is a mock",
		ID:   *id.NewRandomKademliaID(),
	}
	return mockValue, nil
}

func NodeValueLookupMockValidFail(valueID id.KademliaID) (stored.Value, error) {
	var mockValue = stored.Value{
		Data: "This is a fail mock",
		ID:   *id.NewRandomKademliaID(),
	}
	return mockValue, errors.New("value not found")
}
