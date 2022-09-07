package network

import (
	"fmt"
	"net"

	_ "bufio"
)

func TestClient() {
	// Network configurations.'
	const (
		SERVER_HOST = "localhost"
		SERVER_PORT = "4001"
	)

	// Define the network.
	conn, _ := net.Dial("tcp", "127.0.0.1:4001")

	// Send a message to the client.
	fmt.Println("Sending message...")
	sentWords, err := fmt.Fprintf(conn, "Hello world!\n")
	if err != nil {
		fmt.Println("Something went wrong in the sender...")
	}
	fmt.Println("Message was sent, it was ", sentWords, "chars long...")

}
