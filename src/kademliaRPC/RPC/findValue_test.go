package rpc

import (
	kademliaErrors "D7024E/errors"
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/network/requestHandler"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"errors"
	"net"
	"testing"
	"time"
)

// Return a test value.
func testValue() stored.Value {
	return stored.Value{
		Data:   "DATA",
		ID:     *id.NewKademliaID("DATA"),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}
}

// Return a test target contact.
func testContact() contact.Contact {
	return contact.Contact{
		ID:      id.NewRandomKademliaID(),
		Address: "0.0.0.0",
	}
}

// UDPSender mockup that simulates a successful response.
func senderFindValueMockSuccess(_ net.IP, _ int, message []byte) {
	var request rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &request)

	var response []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
			ReqID:   request.ReqID,
			Content: testValue(),
		},
		&response)

	requestHandler.GetInstance().WriteRespone(
		request.ReqID,
		response)
}

// UDPSender mockup that simulates value not found.
func senderFindValueMockNotFound(_ net.IP, _ int, message []byte) {
	var request rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &request)

	var response []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
			ReqID:   request.ReqID,
			Content: stored.Value{},
		},
		&response)

	requestHandler.GetInstance().WriteRespone(
		request.ReqID,
		response)
}

// UDPSender mockup that simulates no response.
func senderFindValueMockFail(_ net.IP, _ int, _ []byte) {}

func TestFindValueRequestValidResponse(t *testing.T) {
	val, err := FindValueRequest(testValue().ID, testContact(), senderFindValueMockSuccess)
	if err != nil {
		t.Fatalf("received error from request that got a valid response")
	} else if !val.Equals(&val) {
		t.Fatalf("value does not equal received value from response")
	}
}

func TestFindValueRequestNoResponse(t *testing.T) {
	val, err := FindValueRequest(testValue().ID, testContact(), senderFindValueMockFail)
	if err == nil {
		t.Fatalf("received no error from request that got a no response")
	} else if !val.Equals(&stored.Value{}) {
		t.Fatalf("received a value containing data, when no response")
	}
}

func TestFindValueRequestValueNotFound(t *testing.T) {
	val, err := FindValueRequest(testValue().ID, testContact(), senderFindValueMockNotFound)
	if !errors.Is(&kademliaErrors.ValueNotFound{}, err) {
		t.Fatalf("received wrong error")
	} else if !val.Equals(&stored.Value{}) {
		t.Fatalf("received none empty value")
	}
}
