package contact

import "testing"

func TestMe(t *testing.T) {
	initVal := GetInstance()
	for i := 0; i < 100; i++ {
		testVal := GetInstance()
		if testVal != initVal {
			t.FailNow()
		}
	}
}

func TestGetAddress(t *testing.T) {
	initVal := getAddress()
	for i := 0; i < 100; i++ {
		testVal := getAddress()
		if testVal != initVal {
			t.FailNow()
		}
	}
}
