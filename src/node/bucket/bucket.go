package bucket

import (
	"D7024E/node/contact"
	"D7024E/node/id"
	"container/list"
)

const bucketSize = 20

// bucket definition
// contains a List
type bucket struct {
	list *list.List
}

// newBucket returns a new instance of a bucket
func newBucket() *bucket {
	bucket := &bucket{}
	bucket.list = list.New()

	return bucket
}

// AddContact adds the Contact to the front of the bucket
// or moves it to the front of the bucket if it already existed
func (bucket *bucket) AddContact(newContact contact.Contact) (oldContact contact.Contact, res bool) {
	var element *list.Element
	var i interface{} = bucket.list.Front().Value
	head := i.(contact.Contact)

	// Loops through the bucket
	for e := bucket.list.Front(); e != nil; e = e.Next() {
		nodeID := e.Value.(contact.Contact).ID

		if (newContact).ID.Equals(nodeID) {
			element = e
		}
	}
	// The new contact was not found in the bucket.
	if element == nil {
		// If the bucket is not full we can just add the new contact as the most recently seen.
		if bucket.list.Len() < bucketSize {
			bucket.list.PushFront(newContact)
			// Return dummy head and res = true since the operation went through.
			return head, true
		} else {
			// If the bucket is full, return the head and res = false as the operation failed.
			return head, false
		}
	} else {
		// New contact was found in the bucket, move the contact to be the most recently seen.
		bucket.list.MoveToFront(element)
		// Return dummy head and res = true since the operation went through.
		return head, true
	}
}

// GetContactAndCalcDistance returns an array of Contacts where
// the distance has already been calculated
func (bucket *bucket) GetContactAndCalcDistance(target *id.KademliaID) []contact.Contact {
	var contacts []contact.Contact

	for elt := bucket.list.Front(); elt != nil; elt = elt.Next() {
		contact := elt.Value.(contact.Contact)
		contact.CalcDistance(target)
		contacts = append(contacts, contact)
	}

	return contacts
}

// Len return the size of the bucket
func (bucket *bucket) Len() int {
	return bucket.list.Len()
}
