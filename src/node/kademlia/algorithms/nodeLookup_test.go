package algorithms

import (
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"fmt"
	"testing"
)

func nodeLookupPingMockSuccess(contact.Contact, rpc.UDPSender) bool {
	return true
}

func nodeLookupPingMockTimeout(contact.Contact, rpc.UDPSender) bool {
	return false
}

func TestAddContacts(t *testing.T) {
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
	fmt.Println(res)
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
