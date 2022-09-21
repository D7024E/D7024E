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
}

// Swap the position of the Contacts at i and j
// WARNING does not check if either i or j is within range
func (candidates *ContactCandidates) Swap(i, j int) {
	candidates.contacts[i], candidates.contacts[j] = candidates.contacts[j], candidates.contacts[i]
}

// Less returns true if the Contact at index i is smaller than
// the Contact at index j
func (candidates *ContactCandidates) Less(i, j int) bool {
	return candidates.contacts[i].Less(&candidates.contacts[j])
}
