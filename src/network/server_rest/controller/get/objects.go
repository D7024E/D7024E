package get

import (
	"D7024E/node/id"
	"D7024E/node/stored"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Get value with given hash.
func Objects(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	valueID, err := id.String2KademliaID(params["hash"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	value, err := stored.GetInstance().FindValue(*valueID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(value)
	}
}
