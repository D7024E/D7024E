package kademliaSort

import (
	"D7024E/node/contact"

	"sort"
)

type ContactList struct {
	Contacts []contact.Contact
}

func SortContacts(batch []contact.Contact) []contact.Contact {
	var toSort ContactList
	toSort.Contacts = batch
	sort.Sort(&toSort)
	return toSort.Contacts
}

func (list *ContactList) AddContact(contact contact.Contact) {
	list.Contacts = append(list.Contacts, contact)
}

// Functions for sort to function...

func (list *ContactList) Len() int {
	return len(list.Contacts)
}

func (list *ContactList) Swap(i int, j int) {
	list.Contacts[i], list.Contacts[j] = list.Contacts[j], list.Contacts[i]
}

func (list *ContactList) Less(i int, j int) bool {
	return list.Contacts[i].Less(&list.Contacts[j])
}
