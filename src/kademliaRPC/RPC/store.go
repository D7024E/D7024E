package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/network/requestHandler"
	"D7024E/network/sender"
	"D7024E/node/contact"
	"D7024E/node/stored"
)

// STORE RPC
// Attempt to store value into target node. If successful return true otherwise
// return false.
func StoreRequest(me contact.Contact, target contact.Contact, value stored.Value) bool {
	reqID := newValidRequestID()
	message := storeRequestMessage(me, reqID, value)
	ip := parseIP(target.Address)
	sender.UDPSender(ip, 4001, message)
	err := requestHandler.GetInstance().ReadResponse(reqID, &message)
	return !isError(err)

}

// Creates the rpc message for store request utilizing json marshaling.
func storeRequestMessage(me contact.Contact, reqID string, value stored.Value) []byte {
	var message []byte
	rpcmarshal.RpcMarshal(rpcmarshal.RPC{
		Cmd:     "STRE",
		Contact: me,
		ReqID:   reqID,
		Content: value,
	}, &message)
	return message
}

// STORE RPC Response
// Stores the given value. Then return a rpc message to inform the requesting
// node that the value is stored.
func StoreRespond(me contact.Contact, target contact.Contact, reqID string, value stored.Value) {
	stored.GetInstance().Store([]stored.Value{value})
	message := storeRespondMessage(me, reqID)
	ip := parseIP(target.Address)
	sender.UDPSender(ip, 4001, message)
}

// Creates the Store Respond by json marshaling of rpc.
func storeRespondMessage(me contact.Contact, reqID string) []byte {
	var message []byte
	rpcmarshal.RpcMarshal(rpcmarshal.RPC{
		Cmd:     "ERTS",
		Contact: me,
		ReqID:   reqID,
	}, &message)
	return message
}
