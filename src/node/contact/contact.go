package contact

import (
	"D7024E/node/id"
	"fmt"
	"sort"
)

// Contact definition
// stores the KademliaID, the ip address and the distance
type Contact struct {
	ID       *id.KademliaID
	Address  string
	distance *id.KademliaID
}

func (contact *Contact) GetDistance() *id.KademliaID {
	return contact.distance
}

func (contact *Contact) SetDistance(dist *id.KademliaID) {
	contact.distance = dist
}

func (c1 *Contact) Equals(c2 *Contact) bool {
	res := c1.ID.Equals(c2.ID)
	res = res && (c1.Address == c2.Address)
	return res
}

// CalcDistance calculates the distance to the target and
// fills the contacts distance field
func (c *Contact) CalcDistance(target *id.KademliaID) {
	c.distance = c.ID.CalcDistance(target)
}

// Less returns true if contact.distance < otherContact.distance
func (c *Contact) Less(otherContact *Contact) bool {
	return c.distance.Less(otherContact.distance)
}

// String returns a simple string representation of a Contact
func (c *Contact) String() string {
	return fmt.Sprintf(`contact("%s", "%s")`, c.ID, c.Address)
}

// ContactCandidates definition
// stores an array of Contacts
type ContactCandidates struct {
	contacts []Contact
}

// Append an array of Contacts to the ContactCandidates
func (candidates *ContactCandidates) Append(contacts []Contact) {
	candidates.contacts = append(candidates.contacts, contacts...)
}

// GetContacts returns the first count number of Contacts
func (candidates *ContactCandidates) GetContacts(count int) []Contact {
	return candidates.contacts[:count]
}

// Sort the Contacts in ContactCandidates
func (candidates *ContactCandidates) Sort() {
	sort.Sort(candidates)
}

// Len returns the length of the ContactCandidates
func (candidates *ContactCandidates) Len() int {
	return len(candidates.contacts)
}ucket.GetInstance()
me := rt.Me

if len(batch) == 0 {
	batch = rt.FindClosestContacts(&destNode, config.Alpha)
}
var newBatch [][]contact.Contact
// For each of the alpha nodes in "batch", send a findNode RPC and append the result to "newBatch"
for i := 0; i < len(batch); i++ {
	var kN []contact.Contact
	kN, err := rpc.FindNode(batch[i], destNode)
	if err != nil {
		log.ERROR("%v", err)
	} else {
		newBatch = append(newBatch, kN)
	}
}

// Convert the contact batch into a single slice.
batch = mergeBatch(newBatch)

// Calculate the distance to each node in the batch and remove duplicates.
distBatch := getAllDistances(*me.ID, batch)
cleanedBatch := removeDuplicates(distBatch)
fmt.Print(cleanedBatch)

// Sort the cleaned batch and extract the alpha closest nodes.
sortedBatch := kademliaSort.SortContacts(cleanedBatch)
alphaNodes := removeDeadNodes(sortedBatch)[:3]

return alphaNodes
} smaller than
// the Contact at index j
func (candidates *ContactCandidates) Less(i, j int) bool {
	return candidates.contacts[i].Less(&candidates.contacts[j])
}
