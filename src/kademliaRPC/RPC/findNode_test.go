package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"errors"
	"net"
	"testing"
)

// UDPSender mockup that simulates a successful response.
func senderFindNodeMockSuccess(_ net.IP, _ int, message []byte) ([]byte, error) {
	var request rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &request)
	request.KNodes = []contact.Contact{testContact()}
	rpcmarshal.RpcMarshal(request, &message)
	return message, nil
}

// UDPSender mockup that simulates no response.
func senderFindNodeMockFail(_ net.IP, _ int, _ []byte) ([]byte, error) {
	return nil, errors.New("sender error")
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
	if KNodes != nil {
		t.Fatalf("received kNodes when no response was given")
	} else if err == nil {
		t.Fatalf("received nil, expected error")
	}
}

// Test FindNodeResponse sends correct message.
func TestFindNodeResponse(t *testing.T) {
	bucket.GetInstance().AddContact(testContact())
	response := FindNodeResponse(*id.NewRandomKademliaID(), testContact())

	var rpcResponse rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(response, &rpcResponse)
	if !rpcResponse.Equals(
		&rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
			KNodes:  []contact.Contact{testContact()},
		}) {
		t.Fatalf("invalid response")
	}
}
