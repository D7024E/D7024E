package main

import (
	"fmt"
	"net"

	"D7024E/node"
)

type Network struct {
	Ip   string
	Port string
}

func Listen(ip string, port int) {
	// TODO

}

func (network *Network) SendPingMessage(contact *node.Contact) {
	// Define a connection using a http protocol and which ip-port  this is stored in "conn".
	conn, _ := net.Dial("udp", "172.0.0.1:4000")
	// Sends the message "Hello world!" to the I/O writer "conn", effectively we print "through"
	// the socket.
	fmt.Fprintf(conn, "Hello world!"+"\n")
}

func (network *Network) SendFindContactMessage(contact *node.Contact) {
	// TODO
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
