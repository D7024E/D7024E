package node

import "time"

type Value struct {
	Data   string        `json:"name"`   // json data as string.
	ID     KademliaID    `json:"id"`     // json id as kademlia id.
	Ttl    time.Duration `json:"ttl"`    // json time-to-live.
	deadAt time.Time     `json:"deadAt"` // json time where value is dead.
}

type Values []Value
