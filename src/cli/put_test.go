package cli

import (
	"D7024E/node/stored"
	"testing"
)

func nodeStoreMockSuccess(stored.Value) bool {
	return true
}

func nodeStoreMockFail(stored.Value) bool {
	return false
}

func refreshAlgorithm(stored.Value) {}

// Verify that a hash is returned on success.
func TestPutSuccess(t *testing.T) {
	var input string = "This is a test where a hash should be returned."
	var res string = Put(input, nodeStoreMockSuccess, refreshAlgorithm)
	if res == "" {
		t.FailNow()
	}
}

// Verify correct fail behaviour on correct input.
func TestPutValidFail(t *testing.T) {
	var input string = "asdjasdpaoisjmmnijdsfa0"
	var res string = Put(input, nodeStoreMockFail, refreshAlgorithm)
	if res != "" {
		t.FailNow()
	}
}
