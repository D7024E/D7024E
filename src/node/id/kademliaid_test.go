package id

import (
	"testing"
)

func TestCalcDistanceSucces(t *testing.T) {
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

func TestCalcDistanceFail(t *testing.T) {
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
