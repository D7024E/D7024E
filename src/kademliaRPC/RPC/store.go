package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/network/requestHandler"
	"D7024E/network/sender"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"net"
)

// STORE RPC
// Attempt to store value into target node. If successful return true otherwise
// return false.
func StoreRequest(me contact.Contact, target contact.Contact, value stored.Value) bool {
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
		Cmd:     "STRE",
		Contact: me,
		ReqID:   reqID,
		Content: value,
	}, message)
	sender.UDPSender(net.ParseIP(target.Address), 4001, *message)

	// Attempt read the response, if successful return true, otherwise false.
	err2 := requestInstance.ReadResponse(reqID, message)
	if err2 != nil {
		return false
	} else {
		return true
	}

}

// STORE RPC Response
// Stores the given value. Then return a rpc message to inform the requesting
// node that the value is stored.
func StoreRespond(me contact.Contact, target contact.Contact, reqID string, value stored.Value) {
	// Create rpc which will be sent.
	rpcMessage := rpcmarshal.RPC{
		Cmd:     "ERTS",
		Contact: me,
		ReqID:   reqID,
	}

	// Store the value within node.
	stored.GetInstance().Store([]stored.Value{value})

	// Marshal rpc and send it to target.
	var message *[]byte
	rpcmarshal.RpcMarshal(rpcMessage, message)
	sender.UDPSender(net.ParseIP(target.Address), 4001, *message)
}
