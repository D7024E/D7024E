package main

import (
	"D7024E/cli"
	"D7024E/network/server"
	"D7024E/network/server_rest"
	"D7024E/node"
	"D7024E/node/contact"
	"net"
)

func main() {
	go server_rest.RestServer(contact.GetInstance().Address, 4000)
	go server.UDPListener(net.ParseIP(contact.GetInstance().Address), 4001)
	go node.StartKademliaNode()
	cli.CliListener()
}
