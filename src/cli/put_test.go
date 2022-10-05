package cli

import (
	"D7024E/node/stored"
	"testing"
)

// Verify correct fail behaviour on correct input.
func TestPutValidFail(t *testing.T) {
	var input string = "put asdjasdpaoisjmmnijdsfa0"
	var res string = Put(input, NodeStoreMockFail)
	if res != "" {
		t.FailNow()
	}
}

func NodeStoreMockSuccess(stored.Value) bool {
	return true
}

func NodeStoreMockFail(stored.Value) bool {
	return false
}
