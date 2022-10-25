package cli

import (
	"D7024E/node/id"
	"testing"
)

func TestForget(t *testing.T) {
	testHash := id.NewRandomKademliaID().String()
	res := Forget(testHash)
	if res != testHash {
		t.FailNow()
	}
}
