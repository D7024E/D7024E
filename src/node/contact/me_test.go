package contact

import "testing"

// Retrieves the "me" instance from "GetInstance()", then retrieves it another
// 100 times. Each retrieval is checked against the first one, if any of them differ
// the test fails.
func TestMe(t *testing.T) {
	initVal := GetInstance()
	for i := 0; i < 100; i++ {
		testVal := GetInstance()
		if testVal != initVal {
			t.FailNow()
		}
	}
}

// Retrieves the nodes address from "getAddress()", then retrieves it another
// 100 times. Each retrieval is checked against the first one, if any of them differ
// the test fails.
func TestGetAddress(t *testing.T) {
	initVal := getAddress()
	for i := 0; i < 100; i++ {
		testVal := getAddress()
		if testVal != initVal {
			t.FailNow()
		}
	}
}
