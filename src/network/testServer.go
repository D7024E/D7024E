package network

import (
	"bufio"
	"fmt"
	"net"
)

func TestServer() {
	// Define a local network struct.
	/*
		serverNet := Network{
			Port: "4000",
		}
	*/

	// Activate a listener on port 4000.
	ln, err := net.Listen("tcp", ":4001")
	fmt.Println("Listening on port: 4001")

	if err != nil {
		fmt.Println("There was an error when listening to port 4000", err)
	}
	// Accepts incoming connection.
	conn, _ := ln.Accept()

	// Infinite loop that listens on port 4000, when it receives a message it prints it out.
	for {
		// Innitiates a new buffered reader on port 4000. The port etc is defined upstream through "conn".
		// Specifically by "ln".
		fmt.Println("Waiting for message...")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// Print out the message that was just received, if any, then continues the loop.
		fmt.Print("Received a message:", string(message))
	}

}
