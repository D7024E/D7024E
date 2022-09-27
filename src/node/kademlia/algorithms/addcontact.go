package algorithms

import (
	"D7024E/node/bucket"
	"D7024E/node/contact"
)

// // Attempt to add a contact to its bucket.
func AddContact(newContact contact.Contact) {
	rt := bucket.GetInstance()
	_, res := rt.AddContact(newContact)
	if res {
		return
	} else {
		// rt.RemoveContact(head)
		// AddContact(newContact)
	}
}
