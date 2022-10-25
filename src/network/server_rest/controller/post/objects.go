package post

import (
	"D7024E/node/id"
	"D7024E/node/kademlia/algorithms"
	"D7024E/node/stored"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Create a value, from given json.
func Objects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var value stored.Value
	err := json.NewDecoder(r.Body).Decode(&value)
	if err != nil || (stored.Value{}).Data == value.Data {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		value.ID = *id.NewKademliaID(value.Data)
		fmt.Println("store value with id: ", value.ID.String())
		if value.Ttl.Nanoseconds() == 0 {
			value.Ttl = time.Minute
		}
		go algorithms.NodeStore(value)
		go algorithms.NodeRefresh(value)
		response := "/objects/" + value.ID.String()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
