package id

import (
	"testing"
)

// Validate that the hashing is consistent.
func TestNewKademliaIDSuccess(t *testing.T) {
	kademliaID := NewKademliaID("data")
	equal := kademliaID.Equals(NewKademliaID("data"))
	if !equal {
		t.FailNow()
	}
}

// Validate that the hashing is consistent.
func TestNewKademliaIDSuccess2(t *testing.T) {
	kademliaID := NewKademliaID("abcdefgh")
	equal := kademliaID.Equals(NewKademliaID("abcdefgh"))
	if !equal {
		t.FailNow()
	}
}

// Validate that different strings result in different hash.
func TestNewKademliaIDFail(t *testing.T) {
	kademliaID := NewKademliaID("data1")
	equal := kademliaID.Equals(NewKademliaID("data2"))
	if equal {
		t.FailNow()
	}
}

// Validate that different strings result in different hash.
func TestNewKademliaIDFail2(t *testing.T) {
	kademliaID := NewKademliaID("dausidoqw9812edqw90usjosjd")
	equal := kademliaID.Equals(NewKademliaID("12tgeydwoiuuaosidj"))
	if equal {
		t.FailNow()
	}
}

// Validate that different strings result in different hash.
func TestNewKademliaIDFail3(t *testing.T) {
	kademliaID := NewKademliaID("Data")
	equal := kademliaID.Equals(NewKademliaID("data"))
	if equal {
		t.FailNow()
	}
}

// func TestNewRandomKademliaID(t *testing.T) {
// 	NewRandomKademliaID()
// 	rand.Seed(1)
// 	id1 := NewRandomKademliaID()
// 	rand.Seed(1)
// 	id2 := NewRandomKademliaID()

// }
