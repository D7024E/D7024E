package bucket

import (
	"D7024E/node/contact"
	"D7024E/node/id"
	"container/list"
)

const BucketSize = 20

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
	for e := bucket.list.Front(); e != nil; e = e.Next() {
		nodeID := e.Value.(contact.Contact).ID

		if (newContact).ID.Equals(nodeID) {
			element = e
		}
	}

	var head contact.Contact
	if bucket.list.Len() == BucketSize {
		head = contact.Contact(bucket.list.Back().Value.(contact.Contact))
		bucket.list.Remove(bucket.list.Back())
	} else {
		head = contact.Contact{}
	}

	if element == nil {
		if bucket.list.Len() < BucketSize {
			bucket.list.PushFront(newContact)
			return head, true
		} else {
			return head, false
		}
	} else {
		bucket.list.MoveToFront(element)
		return head, true
	}
}

func (bucket *bucket) RemoveContact(target contact.Contact) {
	for e := bucket.list.Front(); e != nil; e = e.Next() {
		nodeID := e.Value.(contact.Contact).ID
		if target.ID.Equals(nodeID) {
			bucket.list.Remove(e)
			break
		}
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
