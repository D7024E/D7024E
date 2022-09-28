package contact

import (
	"D7024E/log"
	"D7024E/node/id"
	"net"
	"strings"
	"sync"
)

var me *Contact
var lock = &sync.Mutex{}

// Returns a pointer to the "me" instance if it exists,
// otherwise it creates a new instance and returns a pointer to that.
func GetInstance() *Contact {
	if me != nil {
		return me
	} else {
		lock.Lock()
		defer lock.Unlock()
		if me == nil {
			log.INFO("New \"me\" instance created")
			me = &Contact{
				ID:      id.NewRandomKademliaID(),
				Address: getAddress(),
			}
			// Makes sure that the distance to itself is zero.
			me.distance = me.ID.CalcDistance(me.ID)
		}
	}
	return me
}

// Sends a message to a non-existant address so that the nodes own address can be retrieved.
// Then returns the address.
func getAddress() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
