package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"errors"
	"net"
	"testing"
	"time"
)

// UDPSender mockup that simulates a successful response.
func senderStoreMockSuccess(_ net.IP, _ int, message []byte) ([]byte, error) {
	var request rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &request)

	var response []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
		},
		&response,
	)
	return response, nil
}

// UDPSender mockup that simulates no response.
func senderStoreMockFail(_ net.IP, _ int, _ []byte) ([]byte, error) {
	return nil, errors.New("sender error")
}

// Test StoreRequest when valid response is given.
func TestStoreRequestWithResponse(t *testing.T) {
	target := contact.Contact{
		ID:      id.NewRandomKademliaID(),
		Address: "ADDRESS",
	}
	value := stored.Value{
		Data:   "Data",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}
	res := StoreRequest(target, value, senderStoreMockSuccess)
	if !res {
		t.Fatalf("no response when response was given")
	}
}

// Test Store Request when no response is given.
func TestStoreRequestWithoutResponse(t *testing.T) {
	target := contact.Contact{
		ID:      id.NewRandomKademliaID(),
		Address: "ADDRESS",
	}
	value := stored.Value{
		Data:   "Data",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}
	res := StoreRequest(target, value, senderStoreMockFail)
	if res {
		t.Fatalf("perceived response when none given")
	}

}

// Test if StoreResponse stores value correctly.
func TestStoreResponseSuccess(t *testing.T) {
	target := contact.Contact{
		ID:      id.NewRandomKademliaID(),
		Address: "ADDRESS",
	}
	value := stored.Value{
		Data:   "DATA",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}

	response := StoreResponse(target, value)

	_, err := stored.GetInstance().FindValue(value.ID)
	if err != nil {
		t.Errorf("value not stored")
	}

	var rpcResponse rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(response, &rpcResponse)
	if !rpcResponse.Equals(
		&rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
		}) {
		t.Fatalf("invalid response message sent")
	}
}
