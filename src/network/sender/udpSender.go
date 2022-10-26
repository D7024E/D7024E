package sender

import (
	"D7024E/log"
	"net"
	"strconv"
	"time"
)

// Should be able to replace read response now since it returns the marshalled response?
// Still need to fix the listener side? Maybe just rework how the handler responds to data?

/**
 * Establish udp4 connection with given address created from ip and port.
 * Send message over connection.
 */
func UDPSender(ip net.IP, port int, message []byte) ([]byte, error) {
	addr := ip.String() + ":" + strconv.Itoa(port)
	conn, err := net.Dial("udp4", addr)
	if err != nil {
		log.ERROR("Sender - [%v]", err)
		return nil, err
	}
	defer conn.Close()

	_, err = conn.Write(message)
	if err != nil {
		log.ERROR("Sender - [%v]", err)
		return nil, err
	}

	res := make([]byte, 4096)
	err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.ERROR("Sender - [%v]", err)
		return nil, err
	}
	_, err = conn.Read(res)
	if err != nil {
		log.ERROR("Sender - [%v]", err)
		return nil, err
	}

	return res, nil
}
