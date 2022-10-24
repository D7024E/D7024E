package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/log"
	"D7024E/node/contact"
	"D7024E/node/stored"
)

// STORE RPC
// Attempt to store value into target node. If successful return true otherwise
// return false.
func StoreRequest(target contact.Contact, value stored.Value, sender UDPSender) bool {
	reqID := newValidRequestID()
	var message []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "STRE",
			Contact: *contact.GetInstance(),
			ReqID:   reqID,
			Content: value,
		},
		&message,
	)
	resMessage, err := sender(parseIP(target.Address), 4001, message)
	if err != nil {
		log.ERROR("Error when sending rpc")
		return false
	}

	var rpcMessage rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(resMessage, &rpcMessage)

	return !isError(err)

}

// STORE RPC Response
// Stores the given value. Then return a rpc message to inform the requesting
// node that the value is stored.
func StoreResponse(target contact.Contact, reqID string, value stored.Value, sender UDPSender) {
	stored.GetInstance().Store(value)
	var message []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
			ReqID:   reqID,
		},
		&message,
	)
	sender(parseIP(target.Address), 4001, message)
}
