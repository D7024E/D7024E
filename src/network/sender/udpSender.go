package sender

import (
	"D7024E/log"
	"fmt"
	"net"
	"strconv"
)

/**
 * Establish udp4 connection with given address created from ip and port.
 * Send message over connection.
 */
func UDPSender(ip net.IP, port int, message []byte) {
	addr := ip.String() + ":" + strconv.Itoa(port)
	connection, err := net.Dial("udp4", addr)
	if err != nil {
		log.ERROR("Reccived error ", err)
	} else {
		log.INFO("Setup for sending udp over %s", addr)
	}
	defer connection.Close()

	sentWords, err := fmt.Fprint(connection, message)
	if err != nil {
		log.ERROR("Something went wrong in the sender...")
	} else {
		log.INFO("Message was sent, it was %v chars long...", sentWords)
	}
}
