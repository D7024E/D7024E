package rpc

import (
	"D7024E/errors"
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
)

// FIND_VALUE RPC
// Attempt to find valueID in values in given target. If the given target does not
// have value or that the request timeout, return error. Otherwise return the
// value ID valueID.
func FindValueRequest(valueID id.KademliaID, target contact.Contact, sender UDPSender) (stored.Value, error) {
	var message []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "FIVA",
			Contact: *contact.GetInstance(),
			ID:      valueID,
		},
		&message,
	)

	resMessage, err := sender(parseIP(target.Address), 4001, message)
	if isError(err) {
		return stored.Value{}, err
	}

	var rpcMessage rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(resMessage, &rpcMessage)
	if (stored.Value{} == rpcMessage.Content) {
		return stored.Value{}, &errors.ValueNotFound{}
	} else {
		return rpcMessage.Content, nil
	}
}

// FIND_VALUE RPC Response
// Checks own stored values for Value with ID valueID, if found add value to
// rpc response message, otherwise send message without a value
// (thereby Content will be nil).
func FindValueResponse(target contact.Contact, valueID id.KademliaID, sender UDPSender) []byte {
	rpcMessage := rpcmarshal.RPC{
		Cmd:     "RESP",
		Contact: *contact.GetInstance(),
	}
	value, err := stored.GetInstance().FindValue(valueID)
	if err == nil {
		rpcMessage.Content = value
	}
	var message []byte
	rpcmarshal.RpcMarshal(rpcMessage, &message)
	return message
}
