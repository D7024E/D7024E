package stored

import (
	"D7024E/node/id"
	"fmt"
	"testing"
	"time"
)

// Test success on Equal method for values by comparing the same value.
func TestValueEqualsTrue(t *testing.T) {
	value := Value{
		Data:   "Erik",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}

	if !(value.Equals(&value) == true) {
		t.FailNow()
	}

}

// Test failure on Equal method for values by comparing two different values.
func TestFindValueEqualsFalse(t *testing.T) {
	value1 := Value{
		Data:   "Erik",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}

	value2 := Value{
		Data:   "Dennis",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}
	if !(value1.Equals(&value2) == false) {
		t.FailNow()
	}

}

// See if deadAt is changed when value is not refreshed, since it is dead.
func TestRefreshDeadValue(t *testing.T) {
	t0 := time.Now()
	value := Value{ID: *id.NewRandomKademliaID(), DeadAt: t0}
	res := value.Refresh()
	if res {
		t.FailNow()
	} else if !t0.Equal(value.DeadAt) {
		t.FailNow()
	}
}

// See if value is refreshed when value is still alive.
func TestRefreshAliveValue(t *testing.T) {
	t0 := time.Now().Add(time.Minute)
	value := Value{ID: *id.NewRandomKademliaID(), Ttl: time.Hour, DeadAt: t0}
	res := value.Refresh()
	if !res {
		t.FailNow()
	} else if !t0.Before(value.DeadAt) {
		t.FailNow()
	}
}

// Check if isDead confirmed that value is dead, if deadAt is past.
func TestIsDeadTrue(t *testing.T) {
	value := Value{
		ID:     *id.NewRandomKademliaID(),
		DeadAt: time.Now(),
	}
	res := value.isDead()
	if !res {
		t.FailNow()
	}
}

// Check if isDead confirms that value is not dead.
func TestIsDeadFalse(t *testing.T) {
	value := Value{
		ID:     *id.NewRandomKademliaID(),
		DeadAt: time.Now().Add(time.Hour),
	}
	res := value.isDead()
	if res {
		t.FailNow()
	}
}

// Test to Store one value.
func TestStoreValueSuccess(t *testing.T) {
	stored := Stored{}
	value := Value{
		Data:   "Erik",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}
	err := stored.Store(value)
	if err != nil {
		t.FailNow()
	}
}

// Check if adding duplicate values is accepted.
func TestStoreValueDuplicate(t *testing.T) {
	stored := Stored{}
	value := Value{
		Data:   "Erik",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}
	err := stored.Store(value)
	if err != nil {
		t.FailNow()
	}
	err = stored.Store(value)
	if err == nil {
		fmt.Println(err)
		t.FailNow()
	}
}

// Test to find the stored value.
func TestFindValueSuccess(t *testing.T) {
	stored := Stored{values: []Value{{
		Data:   "Erik",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}}}

	_, err := stored.FindValue(stored.values[0].ID)
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

// Test to delete one and the only element of the storedList.
func TestDeleteValueSuccessWithOneElementStored(t *testing.T) {
	stored := Stored{values: []Value{{
		Data:   "Erik",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}}}

	err := stored.deleteValue(stored.values[0].ID)
	if err != nil {
		t.FailNow()
	}
}

// Test to check so that the correct element is deleted from a list containing three elements.
func TestDeleteValueSuccessWithThreeElementStored(t *testing.T) {
	values := []Value{{
		Data:   "Erik",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}, {
		Data:   "Dennis",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}, {
		Data:   "Anders",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}}

	stored := Stored{values: values}
	err := stored.deleteValue(values[1].ID)
	if err != nil {
		t.FailNow()
	}
}

// Test to delete an existing value which is not part of the storedList which in this test is empty.
func TestDeleteValueOnEmptyList(t *testing.T) {
	err := (&Stored{}).deleteValue(*id.NewRandomKademliaID())
	if err == nil {
		t.FailNow()
	}
}

// Test to delete an value which does not exist in an non-empty storedList.
func TestDeleteValueInAnNonEmptyListFail(t *testing.T) {
	stored := Stored{values: []Value{{
		Data:   "Erik",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}}}

	value2 := Value{
		Data:   "Dennis",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}

	err := stored.deleteValue(value2.ID)
	if err == nil {
		t.FailNow()
	}
}

func TestCleaningDeadValues(t *testing.T) {
	stored := Stored{}
	values := []Value{{ID: *id.NewRandomKademliaID()}}
	stored.values = values
	go stored.cleaningDeadValues(time.Duration(0))
	time.Sleep(1)
	if len(stored.values) != 0 {
		t.FailNow()
	}
}
