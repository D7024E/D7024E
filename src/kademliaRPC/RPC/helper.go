package rpc

import (
	"net"
)

type UDPSender func(net.IP, int, []byte) ([]byte, error)

// Verify if it is an error.
func isError(err error) bool {
	return err != nil
}

// Parse string to net ip.
func parseIP(ip string) net.IP {
	return net.ParseIP(ip)
}
