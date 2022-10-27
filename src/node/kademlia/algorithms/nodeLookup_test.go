package algorithms

import (
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"errors"
	"testing"
)

func nodeLookupPingMockSuccess(contact.Contact, rpc.UDPSender) bool {
	return true
}

func nodeLookupPingMockTimeout(contact.Contact, rpc.UDPSender) bool {
	return false
}

func nodeLookupTestContacts() []contact.Contact {
	rt := bucket.NewRoutingTable()
	rt.AddContact(contact.Contact{ID: id.NewKademliaID("1"), Address: "0.0.0.1"})
	rt.AddContact(contact.Contact{ID: id.NewKademliaID("2"), Address: "0.0.0.2"})
	rt.AddContact(contact.Contact{ID: id.NewKademliaID("3"), Address: "0.0.0.3"})
	return rt.FindClosestContacts(id.NewKademliaID("1"), 3)
}

func nodeLookupTestContacts2() []contact.Contact {
	rt := bucket.NewRoutingTable()
	rt.AddContact(contact.Contact{ID: id.NewKademliaID("1"), Address: "0.0.0.1"})
	rt.AddContact(contact.Contact{ID: id.NewKademliaID("2"), Address: "0.0.0.2"})
	rt.AddContact(contact.Contact{ID: id.NewKademliaID("3"), Address: "0.0.0.3"})
	return rt.FindClosestContacts(id.NewKademliaID("3"), 3)
}

func nodeLookupFindNodeMockSuccess(contact.Contact, id.KademliaID, rpc.UDPSender) ([]contact.Contact, error) {
	return nodeLookupTestContacts(), nil
}

func nodeLookupFindNodeMockSFail(contact.Contact, id.KademliaID, rpc.UDPSender) ([]contact.Contact, error) {
	return nil, errors.New("this is a error")
}

// Add multiple contacts and verify that they are added to routing table.
func TestAddContactsSuccess(t *testing.T) {
	rt := bucket.NewRoutingTable()
	c := contact.Contact{ID: id.NewKademliaID("0"), Address: "0.0.0.0"}
	rt.AddContact(c)
	batch := []contact.Contact{
		c,
		{ID: id.NewKademliaID("1"), Address: "0.0.0.1"},
		{ID: id.NewKademliaID("2"), Address: "0.0.0.2"},
		{ID: id.NewKademliaID("3"), Address: "0.0.0.3"},
	}

	addContacts(rt, batch, nodeLookupPingMockSuccess)
	res := rt.FindClosestContacts(id.NewKademliaID("0"), len(batch)+1)
	if len(res) != len(batch) {
		t.Fatalf("invalid length of added contacts")
	}

	var found bool
	for _, r := range res {
		found = false
		for _, b := range batch {
			if r.Equals(&b) {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("did not found value")
		}
	}
}

// Attempt to add multiple nodes to routing table, when ping fail so that they
// are never added.
func TestAddContactsFail(t *testing.T) {
	rt := bucket.NewRoutingTable()
	c := contact.Contact{ID: id.NewKademliaID("0"), Address: "0.0.0.0"}
	rt.AddContact(c)
	batch := []contact.Contact{
		c,
		{ID: id.NewKademliaID("1"), Address: "0.0.0.1"},
		{ID: id.NewKademliaID("2"), Address: "0.0.0.2"},
		{ID: id.NewKademliaID("3"), Address: "0.0.0.3"},
	}

	addContacts(rt, batch, nodeLookupPingMockTimeout)
	res := rt.FindClosestContacts(id.NewKademliaID("0"), len(batch)+1)
	if len(res) != 1 {
		t.Fatalf("invalid length of added contacts")
	}
}

// TestNodeLookup and verify that the correct values are returned.
func TestNodeLookupSuccess(t *testing.T) {
	rt := bucket.NewRoutingTable()
	targetID := id.NewKademliaID("1")
	rt.AddContact(contact.Contact{ID: targetID, Address: "0.0.0.1"})
	res := nodeLookup(*targetID, rt, nodeLookupPingMockSuccess, nodeLookupFindNodeMockSuccess)
	if len(nodeLookupTestContacts()) != len(res) {
		t.Fatalf("invalid length")
	}
	for i, c := range nodeLookupTestContacts() {
		if !res[i].Equals(&c) {
			t.Fatalf("invalid contacts, expected %v got %v", c, res[i])
		}
	}
}

// Test node lookup when find node rpc fails, and verify result.
func TestNodeLookupFail(t *testing.T) {
	rt := bucket.NewRoutingTable()
	targetID := id.NewKademliaID("1")
	rt.AddContact(contact.Contact{ID: targetID, Address: "0.0.0.1"})
	res := nodeLookup(*targetID, rt, nodeLookupPingMockSuccess, nodeLookupFindNodeMockSFail)
	if len(res) != 1 {
		t.Fatalf("invalid length, expected 1, got %v", len(res))
	} else if !res[0].Equals(&rt.FindClosestContacts(targetID, 1)[0]) {
		t.Fatalf("invalid contact, expected %v, got %v", rt.FindClosestContacts(targetID, 1)[0], res[0])
	}
}

// Test min with different values.
func TestMin(t *testing.T) {
	a := 1
	b := 0
	if min(a, b) != 0 {
		t.Fatalf("expected %v, got %v", b, min(a, b))
	} else if min(b, a) != 0 {
		t.Fatalf("expected %v, got %v", b, min(b, a))
	}
}

// Test min for two equal values.
func TestMinEqual(t *testing.T) {
	a := 1
	b := 1
	if min(a, b) != 1 {
		t.Fatalf("expected %v, got %v", b, min(a, b))
	} else if min(b, a) != 1 {
		t.Fatalf("expected %v, got %v", b, min(b, a))
	}
}

// Test is same on equal batches.
func TestIsSameEquals(t *testing.T) {
	batch := nodeLookupTestContacts()
	newBatch := nodeLookupTestContacts()
	if !isSame(batch, newBatch) {
		t.Fatalf("not equal when values are equal")
	}
}

// Test is two different length of batches if is same returns correctly.
func TestIsSameInvalidLength(t *testing.T) {
	batch := nodeLookupTestContacts()
	newBatch := nodeLookupTestContacts()
	newBatch = append(newBatch, contact.Contact{ID: id.NewRandomKademliaID()})
	if isSame(batch, newBatch) {
		t.Fatalf("not equal when values are equal")
	}
}

// Test is two different batches.
func TestIsSameNotSame(t *testing.T) {
	batch := nodeLookupTestContacts()
	newBatch := nodeLookupTestContacts2()
	if isSame(batch, newBatch) {
		t.Fatalf("not equal when values are equal")
	}
}
