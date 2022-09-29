package id

import (
	"math/rand"
	"testing"
)

// Returns empty kademlia id.
func emptyKademliaID() *KademliaID {
	result := KademliaID{}
	for i := 0; i < IDLength; i++ {
		result[i] = 0
	}
	return &result
}

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

// Test if random kademliaID gives the same result.
func TestNewRandomKademliaID(t *testing.T) {
	id1 := NewRandomKademliaID()
	id2 := NewRandomKademliaID()
	if id1.Equals(id2) {
		t.FailNow()
	}
}

// Test if random kademlia id with same seed gives the same result.
func TestNewRandomKademliaIDSeed(t *testing.T) {
	rand.Seed(1)
	id1 := NewRandomKademliaID()
	rand.Seed(1)
	id2 := NewRandomKademliaID()
	if !id1.Equals(id2) {
		t.FailNow()
	}
}

// Verify Less by comparing n with n+1.
func TestLessTrue(t *testing.T) {
	kademliaID1 := NewKademliaID("id")
	kademliaID2 := NewKademliaID("id")
	kademliaID2[IDLength-1] += 1
	if !kademliaID1.Less(kademliaID2) {
		t.FailNow()
	}
}

// Verify Less by comparing n with n-1.
func TestLessFalse(t *testing.T) {
	kademliaID1 := NewKademliaID("id")
	kademliaID2 := NewKademliaID("id")
	kademliaID2[IDLength-1] -= 1
	if kademliaID1.Less(kademliaID2) {
		t.FailNow()
	}
}

// Verify Less by comparing n with n.
func TestLessEquals(t *testing.T) {
	kademliaID1 := NewKademliaID("id")
	kademliaID2 := NewKademliaID("id")
	if kademliaID1.Less(kademliaID2) {
		t.FailNow()
	}
}

// Verify that two identical ID are true for Equals.
func TestEqualsTrue(t *testing.T) {
	kademliaID1 := NewKademliaID("id")
	kademliaID2 := NewKademliaID("id")
	if !kademliaID1.Equals(kademliaID2) {
		t.FailNow()
	}
}

// Test if Equals fails when one kademlia id is larger.
func TestEqualsLargerID(t *testing.T) {
	kademliaID1 := NewKademliaID("id1")
	kademliaID2 := NewKademliaID("id1")
	kademliaID2[IDLength-1] += 1
	if kademliaID1.Equals(kademliaID2) {
		t.FailNow()
	}
}

// Test if Equals fails when one kademlia id is smaller.
func TestEqualsSmallerID(t *testing.T) {
	kademliaID1 := NewKademliaID("id1")
	kademliaID2 := NewKademliaID("id1")
	kademliaID2[IDLength-1] -= 1
	if kademliaID1.Equals(kademliaID2) {
		t.FailNow()
	}
}

// Test calcDistance on two equal kademlia ids.
func TestCalcDistanceEquals(t *testing.T) {
	kademliaID1 := NewKademliaID("id")
	kademliaID2 := NewKademliaID("id")
	distance := kademliaID1.CalcDistance(kademliaID2)
	if !distance.Equals(emptyKademliaID()) {
		t.FailNow()
	}
}

// Test if calcDistance is 1 when the difference is 1.
func TestCalcDistanceOneDifference(t *testing.T) {
	kademliaID1 := NewKademliaID("id")
	kademliaID2 := NewKademliaID("id")
	kademliaID2[IDLength-1] += 1
	distance := kademliaID1.CalcDistance(kademliaID2)
	result := emptyKademliaID()
	result[IDLength-1] += 1
	if !distance.Equals(result) {
		t.FailNow()
	}
}

// Test if calcDistance is 2 when the difference is 2.
func TestCalcDistanceTwoDifference(t *testing.T) {
	kademliaID1 := NewKademliaID("id")
	kademliaID2 := NewKademliaID("id")
	kademliaID2[IDLength-1] += 2
	distance := kademliaID1.CalcDistance(kademliaID2)
	result := emptyKademliaID()
	result[IDLength-1] += 2
	if !distance.Equals(result) {
		t.FailNow()
	}
}

// Test if kademlia id can be converted to and from string.
func TestString(t *testing.T) {
	kademliaID1 := NewKademliaID("id")
	str := kademliaID1.String()
	kademliaID2, err := String2KademliaID(str)
	if err != nil {
		t.FailNow()
	} else if !kademliaID1.Equals(kademliaID2) {
		t.FailNow()
	}
}
