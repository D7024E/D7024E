package network

import (
	"D7024E/log"
	"net"
)

func Receiver(ip net.IP, port int) {
	connection, err := net.ListenUDP("udp4", &net.UDPAddr{IP: ip, Port: port})
	if err != nil {
		log.ERROR("There was an error:", err)
	} else {
		log.INFO("Setup for listning to udp")
	}
	defer connection.Close()
	buffer := make([]byte, 4096)
	for {
		n, addr, _ := connection.ReadFromUDP(buffer)
		log.INFO("Received ", string(buffer[0:n]), " from ", addr)

	}

}
