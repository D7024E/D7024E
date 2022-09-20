package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/network/requestHandler"
	"D7024E/network/sender"
	"D7024E/node/contact"
	"D7024E/node/id"
	"net"
)

// PING rpc.
// Ping target node, if there is a response return true, otherwise false.
func Ping(me contact.Contact, target contact.Contact) bool {
	// Pointer to request Handler instance.
	requestInstance := requestHandler.GetInstance()

	// Request new request id.
	var reqID string
	var err error
	for {
		reqID = id.NewRandomKademliaID().String()
		err = requestInstance.NewRequest(reqID)
		if err != nil {
			break
		}
	}

	// Create a rpc struct.
	rpcMessage := rpcmarshal.RPC{
		Cmd:     "PING",
		Contact: me,
		ReqID:   reqID,
	}

	// Marshal rpc and send.
	var msg *[]byte
	rpcmarshal.RpcMarshal(rpcMessage, msg)
	sender.UDPSender(net.ParseIP(target.Address), 4001, *msg)

	// Wait for response.
	err2 := requestInstance.ReadResponse(reqID, msg)
	if err2 != nil {
		return false
	} else {
		return true
	}
}

// PONG rpc.
// Ping target node, if there is a response return true, otherwise false.
func Pong(me contact.Contact, target contact.Contact, reqID string) {

	// Create a rpc struct.
	rpcMessage := rpcmarshal.RPC{
		Cmd:     "PONG",
		Contact: me,
		ReqID:   reqID,
	}

	// Marshal rpc and send.
	var msg *[]byte
	rpcmarshal.RpcMarshal(rpcMessage, msg)
	sender.UDPSender(net.ParseIP(target.Address), 4001, *msg)
}

func testPing(me contact.Contact) bool {

	// Pointer to request Handler instance.
	requestInstance := requestHandler.GetInstance()

	// Request new request id.
	var reqID string
	var err error
	for {
		reqID = id.NewRandomKademliaID().String()
		err = requestInstance.NewRequest(reqID)
		if err != nil {
			break
		}
	}

	// Create a rpc struct.
	rpcMessage := rpcmarshal.RPC{
		Cmd:     "PING",
		Contact: me,
		ReqID:   reqID,
	}

	// Marshal rpc and send.
	var msg *[]byte
	rpcmarshal.RpcMarshal(rpcMessage, msg)

	err2 := requestInstance.ReadResponse(reqID, msg)
	if err2 != nil {
		return false
	} else {
		return true
	}
}
