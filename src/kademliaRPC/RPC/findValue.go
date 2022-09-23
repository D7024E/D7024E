package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/network/requestHandler"
	"D7024E/network/sender"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"errors"
	"net"
)

// FIND_VALUE RPC
// Attempt to find valueID in values in given target. If the given target does not
// have value or that the request timeout, return error. Otherwise return the
// value ID valueID.
func FindValueRequest(me contact.Contact, valueID id.KademliaID, target contact.Contact) (stored.Value, error) {
	reqID := newValidRequestID()
	message := findValueRequestMessage(me, reqID, valueID)
	ip := parseIP(target.Address)
	sender.UDPSender(ip, 4001, message)
	err := requestHandler.GetInstance().ReadResponse(reqID, &message)
	return findValueRequestReturn(message, err)
}

// Create FindValue message by marshaling RPC.
func findValueRequestMessage(me contact.Contact, reqID string, valueID id.KademliaID) []byte {
	var message []byte
	rpcmarshal.RpcMarshal(rpcmarshal.RPC{
		Cmd:     "FIVA",
		Contact: me,
		ReqID:   reqID,
		ID:      valueID,
	}, &message)
	return message
}

// Handle the request response, by throwing errors and or unmarshaling.
func findValueRequestReturn(message []byte, err error) (stored.Value, error) {
	if isError(err) {
		return stored.Value{}, err
	} else {
		var rpcMessage rpcmarshal.RPC
		rpcmarshal.RpcUnmarshal(message, &rpcMessage)
		if (stored.Value{} == rpcMessage.Content) {
			return stored.Value{}, errors.New("target node does not have value stored")
		} else {
			return rpcMessage.Content, nil
		}
	}
}

// FIND_VALUE RPC Response
// Checks own stored values for Value with ID valueID, if found add value to
// rpc response message, otherwise send message without a value
// (thereby Content will be nil).
func FindValueResponse(me contact.Contact, target contact.Contact, reqID string, valueID id.KademliaID) {
	message := findValueResponseMessage(me, reqID, valueID)
	sender.UDPSender(net.ParseIP(target.Address), 4001, message)
}

// Create a find value response message and store the value.
func findValueResponseMessage(me contact.Contact, reqID string, valueID id.KademliaID) []byte {
	// Create rpc which will be sent.
	rpcMessage := rpcmarshal.RPC{
		Cmd:     "RESP",
		Contact: me,
		ReqID:   reqID,
	}

	// Check if value is stored within node, if so add it to the rpcMessage.
	value, err := stored.GetInstance().FindValue(valueID)
	if err == nil {
		rpcMessage.Content = value
	}

	// Marshal rpc message.
	var message []byte
	rpcmarshal.RpcMarshal(rpcMessage, &message)
	return message
}
