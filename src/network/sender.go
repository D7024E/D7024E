package network

import (
	"fmt"
	"net"
)

func Sender(ip string, message string) {
	// Define the network.
	conn, _ := net.Dial("udp", "127.0.0.1:4001")

	// Send a message to the client.
	fmt.Println("Sending message...")
	fmt.Println("Trying to send -", message, "-")

	// This is the actual functionality, the rest is just logging.
	// Prints "message" to the listener defined by "conn".
	sentWords, err := fmt.Fprintf(conn, message)
	if err != nil {
		fmt.Println("Something went wrong in the sender...")
	}

	// Log how long the sent message was and theen sleep before looping over again.
	fmt.Println("Message was sent, it was", sentWords, "chars long...")

}
