package objectController

import (
	"D7024E/node"
	"encoding/json"
	"net/http"
)

/**
 * Create a object from json given in request body.
 */
func CreateObject(w http.ResponseWriter, r *http.Request) {
	var object node.Object
	_ = json.NewDecoder(r.Body).Decode(&object)
	//node.Objects = append(node.KandemliaNode.Objects, object)
	node.KandemliaNode.Store(object)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(object)
}
