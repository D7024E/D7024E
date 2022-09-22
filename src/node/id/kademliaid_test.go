package id

import (
	"testing"
)

func TestCalcDistanceEqualsSucces(t *testing.T) {
	idMe := NewRandomKademliaID()
	idTarget := NewRandomKademliaID()

	result1 := idMe.CalcDistance(idTarget)

	result2 := KademliaID{}
	for i := 0; i < IDLength; i++ {
		result2[i] = idMe[i] ^ idTarget[i]
	}

	if !result1.Equals(&result2) {
		t.FailNow()
	}
}

func TestCalcDistanceEqualsFail(t *testing.T) {
	idMe := NewRandomKademliaID()
	idTarget := NewRandomKademliaID()

	idTarget2 := NewRandomKademliaID()

	result1 := idMe.CalcDistance(idTarget)

	result2 := KademliaID{}
	for i := 0; i < IDLength; i++ {
		result2[i] = idMe[i] ^ idTarget2[i]
	}

	if result1.Equals(&result2) {
		t.FailNow()
	}
}

func TestCalcDistanceLessSucces(t *testing.T) {
	idMe := NewRandomKademliaID()
	idTarget := NewRandomKademliaID()

	result1 := idMe.CalcDistance(idTarget)

	result2 := KademliaID{}
	for i := 0; i < IDLength; i++ {
		result2[i] = idMe[i] ^ idTarget[i]
	}

	if result1.Less(&result2) {
		t.FailNow()
	}
}

func TestCalcDistanceLessFail(t *testing.T) {
	idMe := NewRandomKademliaID()
	idTarget := NewRandomKademliaID()

	idTarget2 := NewRandomKademliaID()

	result1 := idMe.CalcDistance(idTarget)

	result2 := KademliaID{}
	for i := 0; i < IDLength; i++ {
		result2[i] = idMe[i] ^ idTarget2[i]
	}

	if !result1.Less(&result2) {
		t.FailNow()
	}
}

func TestNewKademliaIDSuccess(t *testing.T) {
	id := NewRandomKademliaID()
	idString := id.String()
	id2 := NewKademliaID(idString)
	if !id.Equals(id2) {
		t.FailNow()
	}
}
