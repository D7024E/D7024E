package post

import (
	"D7024E/node/id"
	"D7024E/node/kademlia/algorithms"
	"D7024E/node/stored"
	"encoding/json"
	"net/http"
)

func Objects(w http.ResponseWriter, r *http.Request) {
	var value stored.Value
	err := json.NewDecoder(r.Body).Decode(&value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		value.ID = *id.NewKademliaID(value.Data)
		algorithms.NodeStore(value)
		response := "/objects/" + value.ID.String()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
