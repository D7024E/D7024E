package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/network/requestHandler"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"net"
	"testing"
)

// UDPSender mockup that simulates a successful response.
func senderFindNodeMockSuccess(_ net.IP, _ int, message []byte) {
	var request rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &request)
	request.KNodes = []contact.Contact{testContact()}
	rpcmarshal.RpcMarshal(request, &message)

	requestHandler.GetInstance().WriteRespone(
		request.ReqID,
		message)
}

// UDPSender mockup that simulates no response.
func senderFindNodeMockFail(_ net.IP, _ int, _ []byte) {}

// UDPSender mockup that simulates a response.
func senderFindNodeMock(_ net.IP, _ int, message []byte) {
	var request rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &request)

	requestHandler.GetInstance().WriteRespone(
		request.ReqID,
		message)
}

// Test FindNodeRequest on valid response.
func TestFindNodeRequestValidResponse(t *testing.T) {
	KNodes, err := FindNodeRequest(testContact(), *id.NewRandomKademliaID(), senderFindNodeMockSuccess)
	if err != nil {
		t.Fatalf("received %v, expected nil", err)
	} else if !testContact().ID.Equals(KNodes[0].ID) {
		t.Fatalf("received wrong response from FindNode when given valid response")
	}
}

// Test FindNodeRequest on no response.
func TestFindNodeRequestNoResponse(t *testing.T) {
	KNodes, err := FindNodeRequest(testContact(), *id.NewRandomKademliaID(), senderFindNodeMockFail)
	if err == nil {
		t.Fatalf("received nil, expected error")
	} else if KNodes != nil {
		t.Fatalf("received kNodes when no response was given")
	}
}

// Test FindNodeResponse sends correct message.
func TestFindNodeResponse(t *testing.T) {
	bucket.GetInstance().AddContact(testContact())
	reqID := newValidRequestID()
	FindNodeResponse(reqID, *id.NewRandomKademliaID(), testContact(), senderFindNodeMock)

	var response []byte
	err := requestHandler.GetInstance().ReadResponse(reqID, &response)
	if err != nil {
		t.Fatalf("did not read response")
	}

	var rpcResponse rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(response, &rpcResponse)
	if !rpcResponse.Equals(
		&rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
			ReqID:   reqID,
			KNodes:  []contact.Contact{testContact()},
		}) {
		t.Fatalf("invalid response")
	}
}
