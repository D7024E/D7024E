package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/network/requestHandler"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
)

// Send a refresh request to target.
func RefreshRequest(valueID id.KademliaID, target contact.Contact, sender UDPSender) bool {
	reqID := newValidRequestID()
	var message []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESH",
			Contact: *contact.GetInstance(),
			ReqID:   reqID,
		},
		&message)
	sender(parseIP(target.Address), 4001, message)
	err := requestHandler.GetInstance().ReadResponse(reqID, &message)
	return !isError(err)
}

// Respond to a refresh unless the node does not hold the value.
func RefreshResponse(valueID id.KademliaID, target contact.Contact, reqID string, sender UDPSender) {
	_, err := stored.GetInstance().FindValue(valueID)
	if err == nil {
		var message []byte
		rpcmarshal.RpcMarshal(
			rpcmarshal.RPC{
				Cmd:     "RESP",
				Contact: *contact.GetInstance(),
				ReqID:   reqID,
			},
			&message)
		sender(parseIP(target.Address), 4001, message)
	}
}