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
func senderMockSuccess(_ net.IP, _ int, message []byte) {
	var request rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &request)

	var response []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
			ReqID:   request.ReqID,
		},
		&response)

	requestHandler.GetInstance().WriteRespone(
		request.ReqID,
		response)
}

// UDPSender mockup that simulates no response.
func senderMockFail(_ net.IP, _ int, _ []byte) {}

// RefreshRequest that receives a valid response.
func TestRefreshRequestSuccess(t *testing.T) {
	valueID := *id.NewRandomKademliaID()
	target := contact.Contact{
		ID:      id.NewRandomKademliaID(),
		Address: "ADDRESS",
	}
	res := RefreshRequest(valueID, target, senderMockSuccess)
	if !res {
		t.FailNow()
	}
}

// RefreshRequest that receives no response.
func TestRefreshRequestFail(t *testing.T) {
	valueID := *id.NewRandomKademliaID()
	target := contact.Contact{
		ID:      id.NewRandomKademliaID(),
		Address: "ADDRESS",
	}
	res := RefreshRequest(valueID, target, senderMockFail)
	if res {
		t.FailNow()
	}
}

// Test RefreshResponse that it responds and that the response is correct.
func TestRefreshResponse(t *testing.T) {
	valueID := *id.NewRandomKademliaID()
	target := contact.Contact{
		ID:      id.NewRandomKademliaID(),
		Address: "ADDRESS",
	}
	reqID := newValidRequestID()
	stored.GetInstance().Store(stored.Value{
		Data:   "DATA",
		ID:     valueID,
		Ttl:    time.Hour,
		DeadAt: time.Now().Add(time.Hour)})

	RefreshResponse(valueID, target, reqID, senderMockSuccess)

	var response []byte
	err := requestHandler.GetInstance().ReadResponse(reqID, &response)
	if err != nil {
		t.FailNow()
	}

	var rpcResponse rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(response, &rpcResponse)
	if !rpcResponse.Equals(&rpcmarshal.RPC{
		Cmd:     "RESP",
		Contact: *contact.GetInstance(),
		ReqID:   reqID}) {
		t.Errorf("wrong rpc response")
	}

}
