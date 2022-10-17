package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/network/requestHandler"
	"D7024E/node/contact"
	"D7024E/node/id"
	"net"
	"testing"
)

// UDPSender mockup that simulates a successful response.
func senderFindNodeMockSuccess(_ net.IP, _ int, message []byte) {
	var request rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &request)

	var response []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
			ReqID:   request.ReqID,
			KNodes:  []contact.Contact{testContact()},
		},
		&response)

	requestHandler.GetInstance().WriteRespone(
		request.ReqID,
		response)
}

// UDPSender mockup that simulates no response.
func senderFindNodeMockFail(_ net.IP, _ int, _ []byte) {}

func TestFindNodeRequestValidResponse(t *testing.T) {
	KNodes, err := FindNodeRequest(testContact(), *id.NewRandomKademliaID(), senderFindNodeMockSuccess)
	if err != nil {
		t.Fatalf("received %v, expected nil", err)
	} else if !testContact().ID.Equals(KNodes[0].ID) {
		t.Fatalf("received wrong response from FindNode when given valid response")
	}
}

func TestFindNodeRequestNoResponse(t *testing.T) {
	KNodes, err := FindNodeRequest(testContact(), *id.NewRandomKademliaID(), senderFindNodeMockFail)
	if err == nil {
		t.Fatalf("expected error")
	} else if KNodes != nil {
		t.Fatalf("")
	}
}
