package rpc

import (
	"D7024E/environment"
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
)

// FindNode RPC request
// Retrieve k contacts from target node, return error if request timeout.
func FindNodeRequest(target contact.Contact, kademliaID id.KademliaID, sender UDPSender) ([]contact.Contact, error) {
	var message []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "FINO",
			Contact: *contact.GetInstance(),
			ID:      kademliaID,
		},
		&message,
	)
	resMessage, err := sender(parseIP(target.Address), environment.Port, message)
	if isError(err) || resMessage == nil {
		return nil, err
	}
	var rpcMessage rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(resMessage, &rpcMessage)
	return rpcMessage.KNodes, nil
}

// Creates a response RPC struct and populates it with the K, (K = 20), closest nodes to the destination node.
// Which is then sent back to the sender.
func FindNodeResponse(kademliaID id.KademliaID, target contact.Contact) []byte {
	rpcMessage := rpcmarshal.RPC{
		Cmd:     "RESP",
		Contact: *contact.GetInstance(),
	}
	rpcMessage.KNodes = bucket.GetInstance().FindClosestContacts(&kademliaID, 20)
	var message []byte
	rpcmarshal.RpcMarshal(rpcMessage, &message)
	return message
}
