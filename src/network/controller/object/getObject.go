package objectController

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

/**
 * Get a object with given hash as request parameter.
 */
func GetObject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for i, item := range Objects {
		if item.Hash == params["hash"] {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return
		} else if i == len(Objects)-1 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	json.NewEncoder(w).Encode(&Object{})
}
