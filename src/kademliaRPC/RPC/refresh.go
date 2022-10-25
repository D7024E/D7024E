package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/log"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
)

// Send a refresh request to target.
func RefreshRequest(valueID id.KademliaID, target contact.Contact, sender UDPSender) bool {
	var message []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESH",
			Contact: *contact.GetInstance(),
			ID:      valueID,
		},
		&message)

	resMessage, err := sender(parseIP(target.Address), 4001, message)
	if isError(err) || resMessage == nil {
		log.ERROR("Error when sending rpc")
		return false
	}

	var rpcMessage rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(resMessage, &rpcMessage)
	return rpcMessage.Acknowledge
}

// Respond to a refresh unless the node does not hold the value.
func RefreshResponse(valueID id.KademliaID, target contact.Contact) []byte {
	_, err := stored.GetInstance().FindValue(valueID)
	rpcMessage := rpcmarshal.RPC{
		Cmd:         "RESP",
		Contact:     *contact.GetInstance(),
		Acknowledge: !isError(err),
	}

	var message []byte
	rpcmarshal.RpcMarshal(rpcMessage, &message)
	return message
}
