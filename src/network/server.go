package network

import (
	"D7024E/log"
	"encoding/json"
	"net"
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
	json.NewEncoder(w).Encode(object)
}

func getObjects(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
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
	json.NewEncoder(w).Encode(&Object{})
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.INFO(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func getAddress() net.Addr {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.FATAL("Failed to retrieve own IP")
	}

	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr)
}

func StartRouter() {
	log.INFO("Initiated on %v", getAddress())
	router := mux.NewRouter()

	objects = append(objects, Object{Name: "Test_Object", Hash: "TEST_HASH"})
	// objects = append(objects, Object{Name: "THIS_IS_NAME", Hash: "THIS_IS_HASH"})

	router.HandleFunc("/objects", createObject).Methods("POST")
	router.HandleFunc("/all/objects", getObjects).Methods("GET")
	router.HandleFunc("/objects/{hash}", getObject).Methods("GET")

	router.Use(headerMiddleware)
	router.Use(loggingMiddleware)

	log.FATAL("Router has stopped working", http.ListenAndServe(":8000", router))
}
