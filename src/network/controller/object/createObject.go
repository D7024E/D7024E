package objectController

import (
	"encoding/json"
	"net/http"
)

/**
 * Create a object from json given in request body.
 */
func CreateObject(w http.ResponseWriter, r *http.Request) {
	var object Object
	_ = json.NewDecoder(r.Body).Decode(&object)
	Objects = append(Objects, object)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(object)
}
