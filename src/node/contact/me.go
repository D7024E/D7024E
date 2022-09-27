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
			me.distance = me.ID.CalcDistance(me.ID)
		}
	}
	return me
}

func getAddress() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
