package rpc

import (
	"D7024E/network/sender"
	"D7024E/node/bucket"
	"D7024E/node/contact"
)

type pingRPC func(contact.Contact, UDPSender) bool

// Attempt to add a contact to its buckets.
func AddContact(newContact contact.Contact, ping pingRPC) {
	rt := bucket.GetInstance()
	head, res := rt.AddContact(newContact)
	if res {
		return
	} else {
		resp := ping(head, sender.UDPSender)
		if !resp {
			AddContact(newContact, ping)
		}
	}
}
