package objectController

import (
	"net/http"
)

/**
 * Get a object with given hash as request parameter.
 */
func GetObject(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// object, err := node.KandemliaNode.LookupObject(params["hash"])
	// if err != nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// } else {
	// 	w.WriteHeader(http.StatusOK)
	// 	json.NewEncoder(w).Encode(object)
	// }

}
