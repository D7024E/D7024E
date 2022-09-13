package objectController

import (
	"encoding/json"
	"net/http"
)

func GetObjects(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Objects)
}
