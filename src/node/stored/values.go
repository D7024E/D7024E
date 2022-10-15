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
	return res
}

// Checks if value is dead otherwise update the values time to live.
func (value *Value) Refresh() bool {
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
	values []Value
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
	_, err := stored.FindValue(val.ID)
	lock.Lock()
	defer lock.Unlock()
	val.DeadAt = time.Now().Add(val.Ttl)
	if err != nil {
		stored.values = append(stored.values, val)
		return nil
	} else {
		return &errors.ValueAlreadyExist{}
	}
}

// Find a value within stored values.
func (stored *Stored) FindValue(valueId id.KademliaID) (Value, error) {
	lock.Lock()
	defer lock.Unlock()
	for _, item := range stored.values {
		if valueId.Equals(&item.ID) {
			go item.Refresh()
			if item.isDead() {
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
