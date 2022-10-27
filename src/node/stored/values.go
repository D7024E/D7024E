package stored

import (
	"D7024E/errors"
	"D7024E/node/id"
	"fmt"
	"reflect"
	"sync"
	"time"
)

type Value struct {
	Data   string        `json:"data"` // json data as string.
	ID     id.KademliaID `json:"-"`    // json id as kademlia id.
	Ttl    time.Duration `json:"ttl"`  // json time-to-live.
	DeadAt time.Time     `json:"-"`    // json time where value is dead.
}

var lock = &sync.Mutex{}

// Equals for two Value, true if equal otherwise false.
func (v1 *Value) Equals(v2 *Value) bool {
	if reflect.ValueOf(*v2).IsZero() && reflect.ValueOf(*v1).IsZero() {
		return true
	} else if reflect.ValueOf(*v2).IsZero() {
		return false
	} else if reflect.ValueOf(*v1).IsZero() {
		return false
	}

	lock.Lock()
	defer lock.Unlock()
	res := v1.Data == v2.Data
	res = res && v1.ID.Equals(&v2.ID)
	res = res && (v1.Ttl.String() == v2.Ttl.String())
	return res
}

// Checks if value is dead otherwise update the values time to live.
func (value *Value) refresh() bool {
	stored := GetInstance()
	lock.Lock()
	defer lock.Unlock()
	if !value.isDead() {
		value.DeadAt = time.Now().Add(value.Ttl)
		return true
	} else {
		stored.deleteValue(value.ID)
		return false
	}
}

// isDead function to check if value is dead, meaning that the deadAt is past.
func (value *Value) isDead() bool {
	if value.DeadAt.After(time.Now()) {
		return false
	} else {
		return true
	}
}

type Stored struct {
	values    []Value
	refreshed []id.KademliaID
}

var instance *Stored

func GetInstance() *Stored {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Stored{}
			go instance.cleaningDeadValues(15 * time.Second)
		}
	}
	return instance
}

// Store a value within stored values if values id is not already within stored values.
func (stored *Stored) Store(val Value) error {
	val.ID = *id.NewKademliaID(val.Data)
	_, err := stored.FindValue(val.ID)
	lock.Lock()
	defer lock.Unlock()
	val.DeadAt = time.Now().Add(val.Ttl)
	if err != nil {
		fmt.Println("[VALUES] - storing value with id: ", val.ID.String())
		stored.values = append(stored.values, val)
		return nil
	} else {
		fmt.Println("[VALUES] - attempt to store duplicate value with id: ", val.ID.String())
		return &errors.ValueAlreadyExist{}
	}
}

// Find a value within stored values.
func (stored *Stored) FindValue(valueID id.KademliaID) (Value, error) {
	lock.Lock()
	defer lock.Unlock()
	for i, item := range stored.values {
		if valueID.Equals(&item.ID) {
			go stored.values[i].refresh()
			fmt.Println("[VALUES] - find value: ", item)
			if !item.isDead() {
				return item, nil
			} else {
				return Value{}, &errors.ValueTimeout{}
			}
		}
	}
	return Value{}, &errors.ValueNotFound{}
}

// Delete a value with id in stored values.
func (stored *Stored) deleteValue(valueID id.KademliaID) error {
	fmt.Println("[VALUES] - delete value with id: ", valueID.String())
	for i, val := range stored.values {
		if val.ID.Equals(&valueID) {
			stored.values = append(stored.values[:i], stored.values[i+1:]...)
			return nil
		}
	}
	return &errors.ValueNotFound{}
}

// Cleaning of dead values.
func (stored *Stored) cleaningDeadValues(sleepTime time.Duration) {
	time.Sleep(sleepTime)
	lock.Lock()
	deleteID := []id.KademliaID{}
	for _, val := range stored.values {
		if val.isDead() {
			deleteID = append(deleteID, val.ID)
		}
	}
	for _, valueID := range deleteID {
		stored.deleteValue(valueID)
	}
	lock.Unlock()
	go stored.cleaningDeadValues(sleepTime)
}

// Add new value id to refresh.
func (stored *Stored) AddRefresh(valueID id.KademliaID) {
	lock.Lock()
	defer lock.Unlock()
	for _, refreshedValueID := range stored.refreshed {
		if refreshedValueID.Equals(&valueID) {
			return
		}
	}
	stored.refreshed = append(stored.refreshed, valueID)
}

// Is refresh checks if valueID is and should be refreshed.
func (stored *Stored) IsRefreshed(valueID id.KademliaID) bool {
	lock.Lock()
	defer lock.Unlock()
	for _, refreshedValueID := range stored.refreshed {
		if refreshedValueID.Equals(&valueID) {
			return true
		}
	}
	return false
}

// Stop the refreshing of value by removing it from refreshed slice.
func (stored *Stored) StopRefresh(valueID id.KademliaID) {
	lock.Lock()
	defer lock.Unlock()
	for i, refreshedValueID := range stored.refreshed {
		if refreshedValueID.Equals(&valueID) {
			stored.refreshed = append(stored.refreshed[:i], stored.refreshed[i+1:]...)
		}
	}
}
