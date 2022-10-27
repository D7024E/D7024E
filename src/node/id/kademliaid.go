package id

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
	"reflect"
	"time"
)

const IDLength = 20

type KademliaID [IDLength]byte

// Initiates seed using time now.
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Creates new kademlia id by hashing data.
func NewKademliaID(data string) *KademliaID {
	sha := sha256.New()
	hash := sha.Sum([]byte(data))
	newKademliaID := KademliaID{}
	for i := 0; i < IDLength; i++ {
		newKademliaID[i] = hash[i]
	}
	return &newKademliaID
}

// Creates new random kademlia id.
func NewRandomKademliaID() *KademliaID {
	newKademliaID := KademliaID{}
	for i := 0; i < IDLength; i++ {
		newKademliaID[i] = uint8(rand.Intn(256))
	}
	return &newKademliaID
}

// Checks if otherKademliaID is smaller then kademliaID.
func (kademliaID *KademliaID) Less(otherKademliaID *KademliaID) bool {
	for i := 0; i < IDLength; i++ {
		if kademliaID[i] != otherKademliaID[i] {
			return kademliaID[i] < otherKademliaID[i]
		}
	}
	return false
}

// Checks if kademliaID is equal to otherKademliaID.
func (kademliaID *KademliaID) Equals(otherKademliaID *KademliaID) bool {
	if reflect.ValueOf(*kademliaID).IsZero() && reflect.ValueOf(*otherKademliaID).IsZero() {
		return true
	} else if reflect.ValueOf(*kademliaID).IsZero() {
		return false
	} else if reflect.ValueOf(*otherKademliaID).IsZero() {
		return false
	}
	for i := 0; i < IDLength; i++ {
		if kademliaID[i] != otherKademliaID[i] {
			return false
		}
	}
	return true
}

// Calculate distance between two nodes.
func (kademliaID *KademliaID) CalcDistance(target *KademliaID) *KademliaID {
	result := KademliaID{}
	for i := 0; i < IDLength; i++ {
		result[i] = kademliaID[i] ^ target[i]
	}
	return &result
}

// Convert kademlia id to string.
func (kademliaID *KademliaID) String() string {
	return hex.EncodeToString(kademliaID[0:IDLength])
}

// Convert string to kademlia id.
func String2KademliaID(str string) (*KademliaID, error) {
	kademliaID := KademliaID{}
	bytes, _ := hex.DecodeString(str)
	if len(bytes) != IDLength {
		return &kademliaID, errors.New("invalid length str 2 kademlia")
	}
	for i := 0; i < IDLength; i++ {
		kademliaID[i] = bytes[i]
	}
	return &kademliaID, nil
}
