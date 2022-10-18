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
	res := value.refresh()
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
	res := value.refresh()
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
	if err != nil {
		t.FailNow()
	}
}

// Test to find value on a empty list.
func TestFindValueFail(t *testing.T) {
	list := GetInstance()
	id := id.NewRandomKademliaID()

	_, err := list.FindValue(*id)
	if err == nil {
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

// Test if dead values are cleaned by function correctly.
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

// Test whether value id is added to empty refreshed.
func TestAddRefreshAddToEmpty(t *testing.T) {
	valueID := *id.NewRandomKademliaID()
	stored := Stored{refreshed: []id.KademliaID{}}
	stored.AddRefresh(valueID)
	if len(stored.refreshed) != 1 {
		t.Fatalf("invalid length of refreshed values")
	} else if !stored.refreshed[0].Equals(&valueID) {
		t.Fatalf("invalid refreshed id")
	}
}

// Test whether value id is added to none empty refreshed.
func TestAddRefreshNoneEmpty(t *testing.T) {
	valueID := *id.NewRandomKademliaID()
	stored := Stored{refreshed: []id.KademliaID{*id.NewRandomKademliaID()}}
	stored.AddRefresh(valueID)
	if len(stored.refreshed) != 2 {
		t.Fatalf("invalid length of refreshed values")
	} else if !stored.refreshed[1].Equals(&valueID) {
		t.Fatalf("invalid refreshed id")
	}
}

// Test wether duplicate values id are handled.
func TestAddRefreshDuplicate(t *testing.T) {
	valueID := *id.NewRandomKademliaID()
	stored := Stored{refreshed: []id.KademliaID{valueID}}
	stored.AddRefresh(valueID)
	if len(stored.refreshed) != 1 {
		t.Fatalf("invalid length of refreshed values")
	} else if !stored.refreshed[0].Equals(&valueID) {
		t.Fatalf("invalid refreshed id")
	}
}

// Test if IsRefreshed returns the correct bool if valueID is within stored.refreshed.
func TestIsRefreshedTrue(t *testing.T) {
	valueID := *id.NewRandomKademliaID()
	stored := Stored{refreshed: []id.KademliaID{valueID}}
	stored.refreshed = append(stored.refreshed, valueID)
	res := stored.IsRefreshed(valueID)
	if !res {
		t.Fatalf("refreshed value is not within refreshed, where it was added")
	}
}

// Test IsRefreshed if value is not within empty stored.refreshed.
func TestIsRefreshedFalseEmpty(t *testing.T) {
	stored := Stored{}
	res := stored.IsRefreshed(*id.NewRandomKademliaID())
	if res {
		t.Fatalf("value is within refreshed when it was never added")
	}
}

// Test IsRefreshed if value is not within none empty stored.refreshed.
func TestIsRefreshedFalseNoneEmpty(t *testing.T) {
	stored := Stored{refreshed: []id.KademliaID{*id.NewRandomKademliaID()}}
	res := stored.IsRefreshed(*id.NewRandomKademliaID())
	if res {
		t.Fatalf("value is within refreshed when it was never added")
	}
}

// Test whether if the value id is deleted from stored.refreshed.
func TestStopRefresh(t *testing.T) {
	valueID := *id.NewRandomKademliaID()
	stored := Stored{refreshed: []id.KademliaID{valueID}}
	stored.StopRefresh(valueID)
	if len(stored.refreshed) != 0 {
		t.Fatalf("value was note deleted from refreshed values")
	}
}

// Test whether if the value id is not deleted from stored.refreshed,
// when another id is inserted.
func TestStopRefreshDifferentID(t *testing.T) {
	valueID := *id.NewRandomKademliaID()
	stored := Stored{refreshed: []id.KademliaID{valueID}}
	stored.StopRefresh(*id.NewRandomKademliaID())
	if len(stored.refreshed) != 1 {
		t.Fatalf("value was deleted, when stop refresh attempted to delete another value")
	}
}
