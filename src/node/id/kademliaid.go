package id

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"math/rand"
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
	sha := sha1.New()
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

func (kademliaID KademliaID) Less(otherKademliaID *KademliaID) bool {
	for i := 0; i < IDLength; i++ {
		if kademliaID[i] != otherKademliaID[i] {
			return kademliaID[i] < otherKademliaID[i]
		}
	}
	return false
}

func (kademliaID KademliaID) Equals(otherKademliaID *KademliaID) bool {
	for i := 0; i < IDLength; i++ {
		if kademliaID[i] != otherKademliaID[i] {
			return false
		}
	}
	return true
}

func (kademliaID KademliaID) CalcDistance(target *KademliaID) *KademliaID {
	result := KademliaID{}
	for i := 0; i < IDLength; i++ {
		result[i] = kademliaID[i] ^ target[i]
	}
	return &result
}

func (kademliaID *KademliaID) String() string {
	return hex.EncodeToString(kademliaID[0:IDLength])
}

func String2KademliaID(str string) (*KademliaID, error) {
	kademliaID := KademliaID{}
	bytes := []byte(str)
	if len(bytes) != IDLength {
		return &kademliaID, errors.New("invalid length str 2 kademlia")
	}

	for i := 0; i < IDLength; i++ {
		kademliaID[i] = bytes[i]
	}
	return &kademliaID, nil
}
