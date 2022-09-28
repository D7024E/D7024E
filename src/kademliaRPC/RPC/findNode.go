package rpc

import (
	"D7024E/environment"
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/network/requestHandler"
	"D7024E/network/sender"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
)

// FindNode RPC request
// Retrieve k contacts from target node, return error if request timeout.
func FindNodeRequest(me contact.Contact, target contact.Contact, kademliaID id.KademliaID) ([]contact.Contact, error) {
	requestInstance := requestHandler.GetInstance()
	reqID := newValidRequestID()
	message := findNodeRequestMessage(me, reqID, kademliaID)
	ip := parseIP(target.Address)
	sender.UDPSender(ip, environment.Port, message)
	err := requestInstance.ReadResponse(reqID, &message)
	return findNodeRequestReturn(message, err)
}

// Create FindNodeRequest message by marshaling.
func findNodeRequestMessage(me contact.Contact, reqID string, kademliaID id.KademliaID) []byte {
	var message []byte
	rpcmarshal.RpcMarshal(rpcmarshal.RPC{
		Cmd:     "FINO",
		Contact: me,
		ReqID:   reqID,
		ID:      kademliaID,
	}, &message)
	return message
}

// The return response from the request.
func findNodeRequestReturn(message []byte, err error) ([]contact.Contact, error) {
	if isError(err) {
		return nil, err
	}
	var rpcMessage rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &rpcMessage)
	return rpcMessage.KNodes, nil
}

// Creates a response RPC struct and populates it with the K, (K = 20), closest nodes to the destination node.
// Which is then sent back to the sender.
func FindNodeResponse(me contact.Contact, reqID string, kademliaID id.KademliaID, target contact.Contact) {
	message := findNodeResponseMessage(me, reqID, kademliaID)
	ip := parseIP(target.Address)
	sender.UDPSender(ip, environment.Port, message)
}

// Create the response message.
func findNodeResponseMessage(me contact.Contact, reqID string, kademliaID id.KademliaID) []byte {
	rpcMessage := rpcmarshal.RPC{
		Cmd:     "RESP",
		Contact: me,
		ReqID:   reqID,
	}
	rpcMessage.KNodes = bucket.GetInstance().FindClosestContacts(&kademliaID, 20)
	var message []byte
	rpcmarshal.RpcMarshal(rpcMessage, &message)
	return message
}
