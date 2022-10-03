package algorithms

import (
	"D7024E/node/bucket"
	"D7024E/node/contact"
)

// Attempt to add a contact to its buckets.
func AddContact(newContact contact.Contact, ping pingRPC) {
	rt := bucket.GetInstance()
	head, res := rt.AddContact(newContact)
	if res {
		return
	} else {
		resp := ping(*contact.GetInstance(), head)
		if !resp {
			AddContact(newContact, ping)
		}
	}
}
