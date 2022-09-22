package id

import (
	"fmt"
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

func TestEqualsSucces(t *testing.T) {
	idMe := NewRandomKademliaID()

	if !idMe.Equals(idMe) {
		t.FailNow()
	}
}

func TestEqualsFail(t *testing.T) {
	idMe := NewRandomKademliaID()
	idTarget := NewRandomKademliaID()

	if idMe.Equals(idTarget) {
		t.FailNow()
	}
}

func TestLessSucces(t *testing.T) {
	idMe := NewRandomKademliaID()
	idTarget := NewRandomKademliaID()

	fmt.Println("idTarget is " + idTarget.String())
	fmt.Println("idMe is     " + idMe.String())

	fmt.Println(idTarget.Less(idMe))
	fmt.Println(idMe.Less(idTarget))

	if idTarget.Less(idMe) {
		t.FailNow()
	}
}

func TestLessFail(t *testing.T) {
	idMe := NewRandomKademliaID()

	fmt.Println("idMe is     " + idMe.String())

	fmt.Println(idMe.Less(idMe))

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
