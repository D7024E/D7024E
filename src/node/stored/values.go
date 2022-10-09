package stored

import (
	err "D7024E/error"
	"D7024E/node/id"
	"sync"
	"time"
)

type Value struct {
	Data   string        `json:"data"`   // json data as string.
	ID     id.KademliaID `json:"id"`     // json id as kademlia id.
	Ttl    time.Duration `json:"ttl"`    // json time-to-live.
	DeadAt time.Time     `json:"deadAt"` // json time where value is dead.
}

// Equals for two Value, true if equal otherwise false.
func (v1 *Value) Equals(v2 *Value) bool {
	res := v1.Data == v2.Data
	res = res && v1.ID.Equals(&v2.ID)
	res = res && (v1.Ttl.String() == v2.Ttl.String())
	res = res && v1.DeadAt.Equal(v2.DeadAt)
	return res
}

type Stored struct {
	values []Value
}

var lock = &sync.Mutex{}
var instance *Stored

func GetInstance() *Stored {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Stored{}
		}
	}
	return instance
}

// Store a value within stored values if values id is not already within stored values.
func (stored *Stored) Store(val Value) {
	lock.Lock()
	defer lock.Unlock()
	_, err := instance.FindValue(val.ID)
	if err == nil {
		stored.values = append(stored.values, val)
	}
}

// Find a value within stored values.
func (stored *Stored) FindValue(id id.KademliaID) (Value, error) {
	for _, item := range stored.values {
		if id.Equals(&item.ID) {
			return item, nil
		}
	}
	return Value{}, &err.ValueNotFound{}
}

// Delete a value with id in stored values.
func (stored *Stored) DeleteValue(id id.KademliaID) bool {
	lock.Lock()
	defer lock.Unlock()
	index, err := stored.FindValueIndex(id)
	if err == nil {
		stored.values = append(stored.values[:index], stored.values[index+1:]...)
		return true
	} else {
		return false
	}
}

// FindValue Index which is used to determine the Index of the value that is to
// be removed from stored.
func (stored *Stored) FindValueIndex(id id.KademliaID) (int, error) {
	var index int
	for _, item := range stored.values {
		if id.Equals(&item.ID) {
			return index, nil
		}
		index++
	}
	return 0, &err.ValueNotFound{}
}
