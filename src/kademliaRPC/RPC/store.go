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
	var message []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "STRE",
			Contact: *contact.GetInstance(),
			Content: value,
		},
		&message,
	)
	resMessage, err := sender(parseIP(target.Address), 4001, message)
	if isError(err) || resMessage == nil {
		log.ERROR("Error when sending rpc")
		return false
	}

	var rpcMessage rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(resMessage, &rpcMessage)

	return true

}

// STORE RPC Response
// Stores the given value. Then return a rpc message to inform the requesting
// node that the value is stored.
func StoreResponse(target contact.Contact, value stored.Value) []byte {
	stored.GetInstance().Store(value)
	var message []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
		},
		&message,
	)
	return message
}
