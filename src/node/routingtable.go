package node

import (
	err "D7024E/error"
	"D7024E/id"
	"D7024E/log"
	"sync"
)

type routingTable struct {
	me      Contact
	buckets [id.IDLength * 8]*bucket
}

const bucketSize = 20

var instance *routingTable // Singleton instance of routing table
var lock = &sync.Mutex{}   // mutex lock for singleton

func CreateInstance(me Contact) {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		instance = &routingTable{}
		for i := 0; i < id.IDLength*8; i++ {
			instance.buckets[i] = newBucket()
		}
		instance.me = me
	} else {
		log.WARN("Attempted creation of already existing instance of routing table")
	}
}

func GetInstance() (*routingTable, error) {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			return nil, &err.InstanceNotCreated{}
		}
	}
	return instance, nil
}

func (rt *routingTable) AddContact(contact Contact) {
	bucketIndex := rt.getBucketIndex(contact.ID)
	bucket := rt.buckets[bucketIndex]
	bucket.AddContact(contact)
}

func (rt *routingTable) FindClosestContacts(target *id.KademliaID, count int) []Contact {
	var candidates ContactCandidates
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
