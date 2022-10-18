package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/network/requestHandler"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"net"
	"testing"
	"time"
)

// UDPSender mockup that simulates a successful response.
func senderStoreMockSuccess(_ net.IP, _ int, message []byte) {
	var request rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &request)

	var response []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
			ReqID:   request.ReqID,
		},
		&response,
	)

	requestHandler.GetInstance().WriteRespone(
		request.ReqID,
		response)
}

// UDPSender mockup that simulates no response.
func senderStoreMockFail(_ net.IP, _ int, _ []byte) {}

// UDPSender mockup that simulates a response.
func senderStoreMock(_ net.IP, _ int, message []byte) {
	var request rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &request)

	requestHandler.GetInstance().WriteRespone(
		request.ReqID,
		message)
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
	reqID := newValidRequestID()
	value := stored.Value{
		Data:   "DATA",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour),
	}

	StoreResponse(target, reqID, value, senderStoreMock)

	_, err := stored.GetInstance().FindValue(value.ID)
	if err != nil {
		t.Errorf("value not stored")
	}

	var response []byte
	err = requestHandler.GetInstance().ReadResponse(reqID, &response)
	if err != nil {
		t.Fatalf("no response received")
	}

	var rpcResponse rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(response, &rpcResponse)
	if !rpcResponse.Equals(
		&rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
			ReqID:   reqID,
		}) {
		t.Fatalf("invalid response message sent")
	}
}
