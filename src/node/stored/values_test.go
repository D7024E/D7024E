package stored

import (
	"D7024E/node/id"
	"testing"
)

// Test to Store one value.
func TestStoreValueSuccess(t *testing.T) {
	list := GetInstance()
	value := Value{
		Data: "Erik",
		ID:   *id.NewRandomKademliaID(),
	}
	list.Store(value)

	if list == nil {
		t.FailNow()
	}
}

// Test to find the stored value.
func TestFindValueSuccess(t *testing.T) {

	list := GetInstance()
	value := Value{
		Data: "Erik",
		ID:   *id.NewRandomKademliaID(),
	}
	list.Store(value)
	_, err := list.FindValue(value.ID)

	if !(err == nil) {
		t.FailNow()
	}
}

// Test to find value on a empty list.
func TestFindValueFail(t *testing.T) {
	list := GetInstance()
	id := id.NewRandomKademliaID()

	_, err := list.FindValue(*id)

	if !(err != nil) {
		t.FailNow()
	}
}

// Test success on Equal method for values by comparing the same value.
func TestValueEqualsTrue(t *testing.T) {

	list := GetInstance()
	value := Value{
		Data: "Erik",
		ID:   *id.NewRandomKademliaID(),
	}

	list.Store(value)

	if !(value.Equals(&value) == true) {
		t.FailNow()
	}

}

// Test failure on Equal method for values by comparing two different values.
func TestFindValueEqualsFalse(t *testing.T) {

	list := GetInstance()
	value1 := Value{
		Data: "Erik",
		ID:   *id.NewRandomKademliaID(),
	}

	value2 := Value{
		Data: "Dennis",
		ID:   *id.NewRandomKademliaID(),
	}
	list.Store(value1)
	list.Store(value2)

	if !(value1.Equals(&value2) == false) {
		t.FailNow()
	}

}

// Test to delete one and the only element of the storedList.
func TestDeleteValueSuccessWithOneElementStored(t *testing.T) {
	list := GetInstance()
	value := Value{
		Data: "Erik",
		ID:   *id.NewRandomKademliaID(),
	}
	list.Store(value)

	boolean := list.DeleteValue(value.ID)

	if boolean == false {
		t.FailNow()
	}
}

// Test to check so that the correct element is deleted from a list containing three elements.
func TestDeleteValueSuccessWithThreeElementStored(t *testing.T) {
	list := GetInstance()
	value1 := Value{
		Data: "Erik",
		ID:   *id.NewRandomKademliaID(),
	}

	value2 := Value{
		Data: "Dennis",
		ID:   *id.NewRandomKademliaID(),
	}

	value3 := Value{
		Data: "Anders",
		ID:   *id.NewRandomKademliaID(),
	}

	list.Store(value1)
	list.Store(value2)
	list.Store(value3)

	res := list.DeleteValue(value2.ID)

	if res == false {
		t.FailNow()
	}
}

// Test to delete an existing value which is not part of the storedList which in this test is empty.
func TestDeleteValueOnEmptyList(t *testing.T) {
	list := GetInstance()
	value := Value{
		Data: "Erik",
		ID:   *id.NewRandomKademliaID(),
	}

	boolean := list.DeleteValue(value.ID)

	if boolean == true {
		t.FailNow()
	}
}

// Test to delete an value which does not exist in an non-empty storedList.
func TestDeleteValueInAnNonEmptyListFail(t *testing.T) {
	list := GetInstance()
	value1 := Value{
		Data: "Erik",
		ID:   *id.NewRandomKademliaID(),
	}

	value2 := Value{
		Data: "Dennis",
		ID:   *id.NewRandomKademliaID(),
	}

	list.Store(value1)
	boolean := list.DeleteValue(value2.ID)

	if boolean == true {
		t.FailNow()
	}
}
