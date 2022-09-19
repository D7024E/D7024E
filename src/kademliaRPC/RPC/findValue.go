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
	// Get pointer to request id instance.
	requestInstance := requestHandler.GetInstance()

	// Find valid request id then proceed.
	var reqID string
	var err error
	for {
		reqID = id.NewRandomKademliaID().String()
		err = requestInstance.NewRequest(reqID)
		if err == nil {
			break
		}
	}

	// Marshal message then send it to node.
	var message *[]byte
	rpcmarshal.RpcMarshal(rpcmarshal.RPC{
		Cmd:     "FIVA",
		Contact: me,
		ReqID:   reqID,
		ID:      valueID,
	}, message)
	sender.UDPSender(net.ParseIP(target.Address), 4001, *message)

	// Attempt read the response, otherwise throw timeout error.
	err2 := requestInstance.ReadResponse(reqID, message)
	if err2 != nil {
		return stored.Value{}, errors.New("timeout of request")
	}

	// Unmarshal rpc from message and throw error if content is nil.
	var rpcMessage rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(*message, &rpcMessage)
	if (stored.Value{} == rpcMessage.Content) {
		return stored.Value{}, errors.New("target node does not have value stored")
	}

	return rpcMessage.Content, nil

}

// FIND_VALUE RPC Response
// Checks own stored values for Value with ID valueID, if found add value to
// rpc response message, otherwise send message without a value
// (thereby Content will be nil).
func FindValueResponse(me contact.Contact, target contact.Contact, reqID string, valueID id.KademliaID) {
	// Create rpc which will be sent.
	rpcMessage := rpcmarshal.RPC{
		Cmd:     "AVIF",
		Contact: me,
		ReqID:   reqID,
	}

	// Check if value is stored within node, if so add it to the rpcMessage.
	value, err := stored.GetInstance().FindValue(valueID)
	if err == nil {
		rpcMessage.Content = value
	}

	// Marshal rpc and send it to target.
	var message *[]byte
	rpcmarshal.RpcMarshal(rpcMessage, message)
	sender.UDPSender(net.ParseIP(target.Address), 4001, *message)
}
