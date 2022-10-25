package server

import (
	"D7024E/kademliaRPC/handler"
	"D7024E/log"
	"net"
)

// Listener of the udp messages.
// Establish udp4 listner by address reccived from ip and port.
// Read from udp connection into buffer to reccive whole message.
func UDPListener(ip net.IP, port int) {
	connection, err := net.ListenUDP("udp4", &net.UDPAddr{IP: ip, Port: port})
	if err != nil {
		log.ERROR("There was an error: %v", err)
	} else {
		log.INFO("Setup for listening to udp over %v:%v", ip, port)
	}
	defer connection.Close()

	for {
		buffer := make([]byte, 4096)
		n, addr, _ := connection.ReadFromUDP(buffer)
		res := handler.HandleCMD((buffer[0:n]))
		connection.WriteToUDP(res, addr)

	}
}
