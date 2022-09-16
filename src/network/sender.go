package network

import (
	"D7024E/log"
	"fmt"
	"net"
	"strconv"
)

func Sender(ip net.IP, port int, message string) {

	addr := ip.String() + ":" + strconv.Itoa(port)

	connection, err := net.Dial("udp4", addr)
	if err != nil {
		log.ERROR("Reccived error ", err)
	} else {
		log.INFO("Setup for sending udp")
	}

	sentWords, err := fmt.Fprintf(connection, message)
	if err != nil {
		log.ERROR("Something went wrong in the sender...")
	}
	log.INFO("Message was sent, it was", sentWords, "chars long...")

}
