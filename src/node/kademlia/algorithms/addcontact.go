package algorithms

import (
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/node/bucket"
	"D7024E/node/contact"
)

// Attempt to add a contact to its buckets.
func AddContact(newContact contact.Contact) {
	rt := bucket.GetInstance()
	head, res := rt.AddContact(newContact)
	if res {
		return
	} else {
		resp := rpc.Ping(*contact.GetInstance(), head)
		if !resp {
			AddContact(newContact)
		}
	}
}
