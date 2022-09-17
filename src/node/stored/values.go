package stored

import (
	err "D7024E/error"
	"D7024E/id"
	"sync"
	"time"
)

type Value struct {
	Data   string        `json:"name"`   // json data as string.
	ID     id.KademliaID `json:"id"`     // json id as kademlia id.
	Ttl    time.Duration `json:"ttl"`    // json time-to-live.
	DeadAt time.Time     `json:"deadAt"` // json time where value is dead.
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

func (stored *Stored) Store(val []Value) {
	lock.Lock()
	defer lock.Unlock()
	stored.values = append(stored.values, val...)
}

func (stored *Stored) FindValue(id id.KademliaID) (Value, error) {
	for _, item := range stored.values {
		if id.Equals(&item.ID) {
			return item, nil
		}
	}
	return Value{}, &err.ValueNotFound{}
}