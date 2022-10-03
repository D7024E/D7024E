package bucket

import (
	"D7024E/log"
	"D7024E/node/contact"
	"D7024E/node/id"
	"sync"
)

type RoutingTable struct {
	buckets [id.IDLength * 8]*bucket
}

var instance *RoutingTable // Singleton instance of routing table
var lock = &sync.Mutex{}   // mutex lock for singleton

func GetInstance() *RoutingTable {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			log.INFO("New routing table created")
			instance = &RoutingTable{}
			for i := 0; i < id.IDLength*8; i++ {
				instance.buckets[i] = newBucket()
			}
		}
	}
	return instance
}

// Attempt to add a contact to its bucket.
func (rt *RoutingTable) AddContact(newContact contact.Contact) (head contact.Contact, res bool) {
	lock.Lock()
	defer lock.Unlock()
	bucketIndex := rt.getBucketIndex(newContact.ID)
	bucket := rt.buckets[bucketIndex]
	head, res = bucket.AddContact(newContact)
	return head, res
}

func (rt *RoutingTable) RemoveContact(target contact.Contact) {
	lock.Lock()
	defer lock.Unlock()
	bucketIndex := rt.getBucketIndex(target.ID)
	bucket := rt.buckets[bucketIndex]
	bucket.RemoveContact(target)
}

func (rt *RoutingTable) FindClosestContacts(target *id.KademliaID, count int) []contact.Contact {
	var candidates contact.ContactCandidates
	bucketIndex := rt.getBucketIndex(target)
	bucket := rt.buckets[bucketIndex]

	candidates.Append(bucket.GetContactAndCalcDistance(target))

	for i := 1; (bucketIndex-i >= 0 || bucketIndex+i < id.IDLength*8) && candidates.Len() < count; i++ {
		if bucketIndex-i >= 0 {
			bucket = rt.buckets[bucketIndex-i]
			candidates.Append(bucket.GetContactAndCalcDistance(target))
		}
		if bucketIndex+i < id.IDLength*8 {
			bucket = rt.buckets[bucketIndex+i]
			candidates.Append(bucket.GetContactAndCalcDistance(target))
		}
	}

	candidates.Sort()

	if count > candidates.Len() {
		count = candidates.Len()
	}

	return candidates.GetContacts(count)
}

func (rt *RoutingTable) getBucketIndex(contactId *id.KademliaID) int {
	me := contact.GetInstance()
	distance := contactId.CalcDistance(me.ID)
	for i := 0; i < id.IDLength; i++ {
		for j := 0; j < 8; j++ {
			if (distance[i]>>uint8(7-j))&0x1 != 0 {
				return i*8 + j
			}
		}
	}

	return id.IDLength*8 - 1
}
