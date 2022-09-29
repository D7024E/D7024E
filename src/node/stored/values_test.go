package stored

import (
	"D7024E/node/id"
	"fmt"
	"testing"
)

func TestStoreValueSuccess(t *testing.T) {
	list := GetInstance()
	value := Value{
		Data: "Erik",
		ID:   *id.NewRandomKademliaID(),
	}
	list.Store(append(list.values, value))
	if list == nil {
		t.FailNow()
	}
}

func TestFindValueSuccess(t *testing.T) {

	list := GetInstance()
	value := Value{
		Data: "Erik",
		ID:   *id.NewRandomKademliaID(),
	}
	list.Store(append(list.values, value))
	_, err := list.FindValue(value.ID)

	if !(err == nil) {
		t.FailNow()
	}
}

func TestFindValueFail(t *testing.T) {
	list := GetInstance()
	id := id.NewRandomKademliaID()

	_, err := list.FindValue(*id)

	fmt.Println(err)

	if !(err != nil) {
		t.FailNow()
	}
}

func TestFindValueEqualsTrue(t *testing.T) {

	list := GetInstance()
	value := Value{
		Data: "Erik",
		ID:   *id.NewRandomKademliaID(),
	}

	list.Store(append(list.values, value))

	fmt.Println(value.Equals(&value))

	if !(value.Equals(&value) == true) {
		t.FailNow()
	}

}

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
	list.Store(append(list.values, value1))
	list.Store(append(list.values, value2))

	fmt.Println(value1.Equals(&value2))

	if !(value1.Equals(&value2) == false) {
		t.FailNow()
	}

}
