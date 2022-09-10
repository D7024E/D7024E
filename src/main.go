package main

import "D7024E/network"

// "sync"
// "time"

// const (
// 	SupernodeIP string = "172.21.0.2"
// )

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go network.Receiver()
	// time.Sleep(1 * time.Second)
	// network.Sender("127.0.0.1", "Hello world!")
	// wg.Wait()
	network.StartRouter()
}
