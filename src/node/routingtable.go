package node

import (
	"D7024E/contact"
	"D7024E/id"
	"D7024E/log"
	"sync"
)

type routingTable struct {
	me      contact.Contact
	buckets [id.IDLength * 8]*bucket
}

const bucketSize = 20

var instance *routingTable // Singleton instance of routing table
var lock = &sync.Mutex{}   // mutex lock for singleton

func GetInstance() *routingTable {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			log.INFO("New routing table created")
			instance = &routingTable{}
			for i := 0; i < id.IDLength*8; i++ {
				instance.buckets[i] = newBucket()
			}
		}
	}
	return instance
}

func (rt *routingTable) SetMe(me contact.Contact) {
	rt.me = me
}

func (rt *routingTable) GetMe() contact.Contact {
	return rt.me
}

func (rt *routingTable) AddContact(newContact contact.Contact) {
	lock.Lock()
	defer lock.Unlock()
	bucketIndex := rt.getBucketIndex(newContact.ID)
	bucket := rt.buckets[bucketIndex]
	bucket.AddContact(newContact)
}

func (rt *routingTable) FindClosestContacts(target *id.KademliaID, count int) []contact.Contact {
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

func (rt *routingTable) getBucketIndex(contactId *id.KademliaID) int {
	distance := contactId.CalcDistance(rt.me.ID)
	for i := 0; i < id.IDLength; i++ {
		for j := 0; j < 8; j++ {
			if (distance[i]>>uint8(7-j))&0x1 != 0 {
				return i*8 + j
			}
		}
	}

	return id.IDLength*8 - 1
}
