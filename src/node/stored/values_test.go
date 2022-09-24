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
