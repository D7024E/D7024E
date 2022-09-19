package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/network/requestHandler"
	"D7024E/network/sender"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/kademlia"
	"D7024E/node/stored"
	"errors"
	"net"
)

func FindValueRequest(valueID id.KademliaID, node contact.Contact) (stored.Value, error) {
	kademliaInstance := kademlia.GetInstance()
	requestInstance := requestHandler.GetInstance()

	value, err := kademliaInstance.Values.FindValue(valueID)

	if err != nil {
		var reqID string
		var err2 error
		for {
			reqID = id.NewRandomKademliaID().String()
			err2 = requestInstance.NewRequest(reqID)
			if err2 == nil {
				break
			}
		}

		var message *[]byte
		rpcmarshal.RpcMarshal(rpcmarshal.RPC{
			Cmd:     "FIVA",
			Contact: kademliaInstance.Me,
			ReqID:   reqID,
			ID:      valueID,
		}, message)

		var rpcMessage *[]byte
		sender.UDPSender(net.ParseIP(node.Address), 4001, *message)

		err3 := requestInstance.ReadResponse(reqID, rpcMessage)
		if err3 != nil {
			return stored.Value{}, errors.New("timeout of request")
		}

		var rpc rpcmarshal.RPC
		rpcmarshal.RpcUnmarshal(*rpcMessage, &rpc)

		return rpc.Content, nil
	} else {
		return value, nil
	}
}

func FindValueResponse(reqID string, kademliaID id.KademliaID) {
	kademliaInstance := kademlia.GetInstance()
	requestInstance := requestHandler.GetInstance()

	value, err := kademliaInstance.Values.FindValue(kademliaID)

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
