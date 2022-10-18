package algorithms

import (
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"errors"
	"testing"
	"time"
)

// Mockup function for FindValue rpc, which always succeeds.
func findValueSuccess(valueID id.KademliaID, _ contact.Contact, _ rpc.UDPSender) (stored.Value, error) {
	return stored.Value{
		Data:   "DATA",
		ID:     valueID,
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}, nil
}

// Mockup function for FindValue rpc, which always fails.
func findValueFail(id.KademliaID, contact.Contact, rpc.UDPSender) (stored.Value, error) {
	return stored.Value{}, errors.New("value not found")
}

// Test for AlphaNodesValueLookup which utilizes findValueSuccess mockup which
// should result in the function always succeeding.
func TestAlphaNodeValueLookupSuccess(t *testing.T) {
	valueID := id.NewRandomKademliaID()
	alphaClosest := []contact.Contact{
		{
			ID:      id.NewRandomKademliaID(),
			Address: "127.21.0.2",
		},
		{
			ID:      id.NewRandomKademliaID(),
			Address: "127.21.0.3",
		},
		{
			ID:      id.NewRandomKademliaID(),
			Address: "127.21.0.4",
		},
	}
	_, err := alphaNodeValueLookup(*valueID, alphaClosest, findValueSuccess)
	if err != nil {
		t.FailNow()
	}
}

// Test for AlphaNodesValueLookup which utilizes findValueFail mockup which
// should result in the function always failing.
func TestAlphaNodeValueLookupFail(t *testing.T) {
	valueID := id.NewRandomKademliaID()
	alphaClosest := []contact.Contact{
		{
			ID:      id.NewRandomKademliaID(),
			Address: "127.21.0.2",
		},
		{
			ID:      id.NewRandomKademliaID(),
			Address: "127.21.0.3",
		},
		{
			ID:      id.NewRandomKademliaID(),
			Address: "127.21.0.4",
		}}
	_, err := alphaNodeValueLookup(*valueID, alphaClosest, findValueFail)
	if err == nil {
		t.FailNow()
	}
}
