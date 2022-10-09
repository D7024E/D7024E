package stored

import (
	"D7024E/errors"
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

var lock = &sync.Mutex{}

// Equals for two Value, true if equal otherwise false.
func (v1 *Value) Equals(v2 *Value) bool {
	lock.Lock()
	defer lock.Unlock()
	res := v1.Data == v2.Data
	res = res && v1.ID.Equals(&v2.ID)
	res = res && (v1.Ttl.String() == v2.Ttl.String())
	res = res && v1.DeadAt.Equal(v2.DeadAt)
	return res
}

type Stored struct {
	values []Value
}

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
func (stored *Stored) Store(val Value) error {
	_, err := GetInstance().FindValue(val.ID)
	lock.Lock()
	defer lock.Unlock()
	if err != nil {
		stored.values = append(stored.values, val)
		return nil
	} else {
		return &errors.ValueAlreadyExist{}
	}
}

// Find a value within stored values.
func (stored *Stored) FindValue(id id.KademliaID) (Value, error) {
	lock.Lock()
	defer lock.Unlock()
	for _, item := range stored.values {
		if id.Equals(&item.ID) {
			return item, nil
		}
	}
	return Value{}, &errors.ValueNotFound{}
}

// Delete a value with id in stored values.
func (stored *Stored) deleteValue(valueID id.KademliaID) error {
	values := GetInstance().values
	lock.Lock()
	defer lock.Unlock()
	for i, val := range values {
		if val.ID.Equals(&valueID) {
			stored.values = append(stored.values[:i], stored.values[i+1:]...)
			return nil
		}
	}
	return &errors.ValueNotFound{}
}

// isDead function to check if value is dead, meaning that the deadAt is past.
// If value is dead silently delete it.
func (stored *Stored) isDead(val Value) bool {
	lock.Lock()
	if val.DeadAt.After(time.Now()) {
		lock.Unlock()
		return false
	} else {
		lock.Unlock()
		GetInstance().deleteValue(val.ID)
		return true
	}
}
