package server_rest

import (
	"D7024E/log"
	"D7024E/network/server_rest/controller/get"
	"D7024E/network/server_rest/controller/post"
	"D7024E/network/server_rest/middleware/logging"
	"net/http"

	"github.com/gorilla/mux"
)

// Start rest server.
func RestServer(ip string, port int) {
	router := mux.NewRouter()
	registerRoutes(router)
	log.INFO("Setup for rest over %v:%v", ip, port)
	log.FATAL("Rest router has stopped working", http.ListenAndServe(":4000", router)) //ip+":"+strconv.Itoa(port)
}

// Register routes and middleware.
func registerRoutes(router *mux.Router) {
	router.HandleFunc("/objects", post.Objects).Methods("POST")
	router.HandleFunc("/objects/{hash}", get.Objects).Methods("GET")
	router.Use(logging.Start())
}
