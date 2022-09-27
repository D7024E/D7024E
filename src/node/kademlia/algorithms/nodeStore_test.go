package algorithms

import (
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"math/rand"
	"testing"
	"time"
)

// Store rpc mockup, will always succeed.
func storeSuccess(_ contact.Contact, _ contact.Contact, _ stored.Value) bool {
	return true
}

// Store rpc mockup, will return true 50% of the time.
func store50RandomSuccess(_ contact.Contact, _ contact.Contact, _ stored.Value) bool {
	if rand.Intn(2) == 1 {
		return true
	} else {
		return false
	}
}

// Test AlphaNodeStoreRec if rpc always succeed.
func TestAlphaNodeStoreRecSuccess(t *testing.T) {
	value := stored.Value{
		Data:   "DATA",
		ID:     *id.NewKademliaID("DATA"),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}
	success := AlphaNodeStoreRec(value, storeSuccess)
	if !success {
		t.FailNow()
	}
}

// Test AlphaNodeStoreRec if it will eventually succeed if rpc return true 50% of the time.
func TestAlphaNodeStoreRecRandomSuccess(t *testing.T) {
	value := stored.Value{
		Data:   "DATA",
		ID:     *id.NewKademliaID("DATA"),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}
	success := AlphaNodeStoreRec(value, store50RandomSuccess)
	if !success {
		t.FailNow()
	}
}
