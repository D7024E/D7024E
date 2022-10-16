package rpc

import (
	"D7024E/network/requestHandler"
	"D7024E/node/id"
	"net"
)

type UDPSender func(net.IP, int, []byte)

// Returns a new valid requestID.
func newValidRequestID() string {
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

	return reqID
}

// Verify if it is an error.
func isError(err error) bool {
	return err != nil
}

// Parse string to net ip.
func parseIP(ip string) net.IP {
	return net.ParseIP(ip)
}
