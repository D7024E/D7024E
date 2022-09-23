package rpc

import (
	"D7024E/network/requestHandler"
	"errors"
	"testing"
)

// Test to verify that a valid request id is created.
func TestNewValidRequestIDSuccess(t *testing.T) {
	requestInstance := requestHandler.GetInstance()
	reqID := newValidRequestID()
	err := requestInstance.WriteRespone(reqID, []byte("this is response"))
	if err != nil {
		t.FailNow()
	}
}

// Test that IsError returns a true on error.
func TestIsErrorTrue(t *testing.T) {
	res := isError(errors.New("this is error"))
	if !res {
		t.FailNow()
	}
}

// Test that IsError returns false on nil.
func TestIsErrorFalse(t *testing.T) {
	res := isError(nil)
	if res {
		t.FailNow()
	}
}

// Test if parse ip creates valid ip.
func TestParseIPSuccess(t *testing.T) {
	ip := parseIP("127.0.0.1")
	if ip == nil {
		t.FailNow()
	}
}

func TestParseIPFail(t *testing.T) {
	ip := parseIP("45789")
	if ip != nil {
		t.FailNow()
	}
}
