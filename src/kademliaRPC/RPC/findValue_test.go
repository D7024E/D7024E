package rpc

import (
	kademliaErrors "D7024E/errors"
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/node/contact"
	"D7024E/node/stored"
	"errors"
	"fmt"
	"net"
	"testing"
)

// UDPSender mockup that simulates a successful response.
func senderFindValueMockSuccess(_ net.IP, _ int, message []byte) ([]byte, error) {
	var request rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &request)

	var response []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
			Content: testValue(),
		},
		&response)
	return response, nil
}

// UDPSender mockup that simulates value not found.
func senderFindValueMockNotFound(_ net.IP, _ int, message []byte) ([]byte, error) {
	var request rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &request)

	var response []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
			Content: stored.Value{},
		},
		&response)
	return response, nil
}

// UDPSender mockup that simulates no response.
func senderFindValueMockFail(_ net.IP, _ int, _ []byte) ([]byte, error) {
	return nil, errors.New("sender error")
}

// Test FindValueRequest if valid response gets the correct output.
func TestFindValueRequestValidResponse(t *testing.T) {
	val, err := FindValueRequest(testValue().ID, testContact(), senderFindValueMockSuccess)
	if err != nil {
		t.Fatalf("received error from request that got a valid response")
	} else if !val.Equals(&val) {
		t.Fatalf("value does not equal received value from response")
	}
}

// Test FindValueRequest if right output is given when there is no response.
func TestFindValueRequestNoResponse(t *testing.T) {
	val, err := FindValueRequest(testValue().ID, testContact(), senderFindValueMockFail)
	if err == nil {
		t.Fatalf("received no error from request that got a no response")
	} else if !val.Equals(&stored.Value{}) {
		fmt.Println(val)
		fmt.Println(stored.Value{})
		t.Fatalf("received a value containing data, when no response")
	}
}

// Test FindValueRequest if right output is given when value is not found.
func TestFindValueRequestValueNotFound(t *testing.T) {
	val, err := FindValueRequest(testValue().ID, testContact(), senderFindValueMockNotFound)
	if !errors.Is(&kademliaErrors.ValueNotFound{}, err) {
		t.Fatalf("received wrong error")
	} else if !val.Equals(&stored.Value{}) {
		t.Fatalf("received none empty value")
	}
}

// Test FindValueResponse if correct response is given.
func TestFindValueResponseFoundValue(t *testing.T) {
	value := testValue()
	err := stored.GetInstance().Store(value)
	if err != nil {
		t.FailNow()
	}
	response := FindValueResponse(testContact(), value.ID)
	var rpcResponse rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(response, &rpcResponse)
	if !rpcResponse.Equals(
		&rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
			Content: stored.Value{
				Data: value.Data,
				Ttl:  value.Ttl,
			},
		}) {
		t.Fatalf("invalid response")
	}
}
