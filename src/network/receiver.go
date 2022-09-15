package network

import (
	"encoding/hex"
	"fmt"
	"net"
)

func Receiver() {
	ip := net.IPv4(127, 0, 0, 1)
	udpAddr := net.UDPAddr{}
	udpAddr.IP = ip
	udpAddr.Port = 4001
	udpAddr.Zone = "IPv4"
	// Activate a listener on port 4001.
	conn, err := net.ListenUDP("udp4", &udpAddr)
	fmt.Println("Listening on port: 4001")

	if err != nil {
		fmt.Println("There was an error:", err)
	}

	// Infinite loop that listens on port 4001, when it receives a message it prints it out.
	for {
		// // Innitiates a new buffered reader on port 4000. The port etc is defined upstream through "conn".
		// // Specifically by "ln".
		// fmt.Println("Waiting for message...")
		// message, _ := bufio.NewReader(conn).ReadString('\n')
		// // Print out the message that was just received, if any, then continues the loop.
		// fmt.Print("Received a message: ", string(message))

		display(conn)
	}

}

func display(conn *net.UDPConn) {

	var buf [2048]byte
	n, err := conn.Read(buf[0:])
	if err != nil {
		fmt.Println("Error Reading")
		return
	} else {
		fmt.Println(hex.EncodeToString(buf[0:n]))
		fmt.Println("Package Done")
	}

}
