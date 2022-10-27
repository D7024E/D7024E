package algorithms

import (
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"math/rand"
	"testing"
	"time"
)

// NodeLookup mockup
func nodeLookupMock(_ id.KademliaID) []contact.Contact {
	return []contact.Contact{
		{ID: id.NewKademliaID("DATA"), Address: "172.21.0.3"},
		{ID: id.NewKademliaID("DATA1"), Address: "172.21.0.4"},
		{ID: id.NewKademliaID("DATA2"), Address: "172.21.0.5"},
		{ID: id.NewKademliaID("DATA3"), Address: "172.21.0.6"},
		{ID: id.NewKademliaID("DATA4"), Address: "172.21.0.7"},
	}
}

// Store rpc mockup, will always succeed.
func storeSuccess(contact.Contact, stored.Value, rpc.UDPSender) bool {
	return true
}

// Store rpc mockup, will return true 50% of the time.
func store50RandomSuccess(contact.Contact, stored.Value, rpc.UDPSender) bool {
	if rand.Intn(2) == 1 {
		return true
	} else {
		return false
	}
}

// Test AlphaNodeStoreRec if rpc always succeed.
func TestKNodeStoreRecSuccess(t *testing.T) {
	value := stored.Value{
		Data:   "DATA",
		ID:     *id.NewKademliaID("DATA"),
		TTL:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}
	success := AlphaNodeStoreRec(value, storeSuccess, nodeLookupMock)
	if !success {
		t.FailNow()
	}
}

// Test AlphaNodeStoreRec if it will eventually succeed if rpc return true 50% of the time.
func TestKNodeStoreRecRandomSuccess(t *testing.T) {
	value := stored.Value{
		Data:   "DATA",
		ID:     *id.NewKademliaID("DATA"),
		TTL:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}
	success := AlphaNodeStoreRec(value, store50RandomSuccess, nodeLookupMock)
	if !success {
		t.FailNow()
	}
}
