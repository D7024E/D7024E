package network

import (
	"D7024E/log"
	headerMiddleware "D7024E/network/middleware/header"
	loggingMiddleware "D7024E/network/middleware/logging"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Object struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}

var objects []Object

func createObject(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	var object Object
	_ = json.NewDecoder(r.Body).Decode(&object)
	objects = append(objects, object)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(object)
}

func getObjects(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(objects)
}

func getObject(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range objects {
		if item.Hash == params["hash"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Object{})
}

func StartRouter() {
	log.INFO("Initiated on %v", GetAddress())
	router := mux.NewRouter()

	objects = append(objects, Object{Name: "Test_Object", Hash: "TEST_HASH"})
	// objects = append(objects, Object{Name: "THIS_IS_NAME", Hash: "THIS_IS_HASH"})

	router.HandleFunc("/objects", createObject).Methods("POST")
	router.HandleFunc("/all/objects", getObjects).Methods("GET")
	router.HandleFunc("/objects/{hash}", getObject).Methods("GET")

	router.Use(headerMiddleware.HeaderMiddleware)

	logMiddleware := loggingMiddleware.NewLogMiddleware(log.InfoLogger)
	router.Use(logMiddleware.Start())

	log.FATAL("Router has stopped working", http.ListenAndServe(":8000", router))
}
