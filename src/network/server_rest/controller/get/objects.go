package get

import (
	"D7024E/node/id"
	"D7024E/node/kademlia/algorithms"
	"D7024E/node/stored"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// private variable for node value lookup to be able to mock the rest request
var nvl func(id.KademliaID) (stored.Value, error) = algorithms.NodeValueLookup

// Get value with given hash.
func Objects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	valueID, err := id.String2KademliaID(params["hash"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	value, err := nvl(*valueID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(value)
	}
}
