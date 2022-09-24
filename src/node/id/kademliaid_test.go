package id

import (
	"fmt"
	"testing"
)

func TestCalcDistanceSuccess(t *testing.T) {
	idMe := NewRandomKademliaID()
	idTarget := NewRandomKademliaID()
	// Call calcDistance function.
	result1 := idMe.CalcDistance(idTarget)

	// Count distance with a for loop on the same values.
	result2 := KademliaID{}
	for i := 0; i < IDLength; i++ {
		result2[i] = idMe[i] ^ idTarget[i]
	}

	if !result1.Equals(&result2) {
		t.FailNow()
	}
}

func TestCalcDistanceFail(t *testing.T) {
	idMe := NewRandomKademliaID()
	idTarget := NewRandomKademliaID()
	idTarget2 := NewRandomKademliaID()

	// Call calcDistance function.
	result1 := idMe.CalcDistance(idTarget)

	// Count distance with a different target ID which will cause it to fail.
	result2 := KademliaID{}
	for i := 0; i < IDLength; i++ {
		result2[i] = idMe[i] ^ idTarget2[i]
	}

	if result1.Equals(&result2) {
		t.FailNow()
	}
}

func TestEqualsSuccess(t *testing.T) {
	idMe := NewRandomKademliaID()

	// Checking if idMe == idMe which it should be.
	if !idMe.Equals(idMe) {
		t.FailNow()
	}
}

func TestEqualsFail(t *testing.T) {
	idMe := NewRandomKademliaID()
	idTarget := NewRandomKademliaID()

	// Checking if idMe == idTarget which it shouldn't be.
	if idMe.Equals(idTarget) {
		t.FailNow()
	}
}

func TestLessSucces(t *testing.T) {
	idMe := NewRandomKademliaID()
	idTarget := NewRandomKademliaID()

	// The same id is always generated.
	// idMe = 210fc7bb818639ac48a4c6afa2f1581a8b9525e2.
	// idTarger = 0fda68927f2b2ff836f73578db0fa54c29f7fd92.

	// Check if idMe is less than idTarget which it should be
	if idMe.Less(idTarget) {
		t.FailNow()
	}
}

func TestLessFail(t *testing.T) {
	idMe := NewRandomKademliaID()
	idTarget := NewRandomKademliaID()

	// The same id is always generated.
	// idMe = 210fc7bb818639ac48a4c6afa2f1581a8b9525e2
	// idTarger = 0fda68927f2b2ff836f73578db0fa54c29f7fd92

	// Check if idMe is less than idTarget which it shouldn't be
	if !idTarget.Less(idMe) {
		t.FailNow()
	}
}

func TestLessWithSameID(t *testing.T) {
	idMe := NewRandomKademliaID()
	// The same id is always generated.
	// idMe = 210fc7bb818639ac48a4c6afa2f1581a8b9525e2.

	fmt.Println(idMe)

	// Check if idMe is less than idMe which it shouldn't be because the ID's are the same.
	if idMe.Less(idMe) {
		t.FailNow()
	}
}

func TestNewKademliaID(t *testing.T) {
	id := NewRandomKademliaID()
	idString := id.String()
	id2 := NewKademliaID(idString)
	if !id.Equals(id2) {
		t.FailNow()
	}
}
