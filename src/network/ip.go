package network

import (
	"D7024E/log"
	"net"
)

/**
 * Get of ip and port currently used.
 */
func GetAddress() net.Addr {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.FATAL("Failed to retrieve own IP")
	}

	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr)
}
