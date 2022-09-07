package main

import (
	"D7024E/network"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go network.TestServer()
	time.Sleep(1 * time.Second)
	network.TestClient()
	wg.Wait()
}
