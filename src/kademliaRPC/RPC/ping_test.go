package rpc

import (
	"D7024E/node/contact"
	"D7024E/node/id"
	"fmt"
	"testing"
)

func TestPing(t *testing.T) {

	idMe := id.NewRandomKademliaID()

	me := contact.Contact{
		ID: idMe,
	}

	fmt.Printf("String id are: " + idMe.String())

	result := testPing(me)

	if result != true {
		t.Errorf("Ping(me, target) FAILED. Expected %t, got %t", true, result)
	} else {
		t.Logf("Ping(me, target) PASSED. Expected %t, got %t", true, result)
	}
}
