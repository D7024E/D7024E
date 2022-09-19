package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/network/requestHandler"
	"D7024E/node/id"
	"D7024E/node/kademlia"
)

func FindValueRecciver(kademliaID id.KademliaID) {
	kademliaInstance := kademlia.GetInstance()
	requestInstance := requestHandler.GetInstance()

	value, err := kademliaInstance.Values.FindValue(kademliaID)

	var reqID string
	for {
		reqID = id.NewRandomKademliaID().String()
		err2 := requestInstance.NewRequest(reqID)
		if err2 == nil {
			break
		}
	}

	var message *[]byte

	if err != nil {
		rpcmarshal.RpcMarshal(rpcmarshal.RPC{
			Cmd:     "AVIF",
			Contact: kademliaInstance.Me,
			ReqID:   reqID,
		}, message)
	} else {
		rpcmarshal.RpcMarshal(rpcmarshal.RPC{
			Cmd:     "AVIF",
			Contact: kademliaInstance.Me,
			ReqID:   reqID,
			Content: value,
		}, message)
	}

	requestInstance.WriteRespone(reqID, *message)
}
