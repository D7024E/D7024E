package server

import (
	rpc "D7024E/kademliaRPC"
	"D7024E/log"
	"net"
)

// Listener of the udp messages.
// Establish udp4 listner by address reccived from ip and port.
// Read from udp connection into buffer to reccive whole message.
func UDPListener(ip net.IP, port int) {
	connection, err := net.ListenUDP("udp4", &net.UDPAddr{IP: ip, Port: port})
	if err != nil {
		log.ERROR("There was an error:", err)
	} else {
		log.INFO("Setup for listning to udp over %v:%v", ip, port)
	}
	defer connection.Close()
	for {
		buffer := make([]byte, 4096)
		n, addr, _ := connection.ReadFromUDP(buffer)
		go rpc.InitiateCMD((buffer[0:n]))
		log.INFO("Received \"%s\" from %v", string(buffer[0:n]), addr)

	}
}
