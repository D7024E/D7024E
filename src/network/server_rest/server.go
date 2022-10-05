package server_rest

import (
	"D7024E/log"
	"D7024E/network/server_rest/route"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/**
 * Start a server with REST http structure.
 */
func RestServer(ip string, port int) {
	router := mux.NewRouter()
	route.RegisterRoutes(router)
	log.INFO("Setup for rest over %v:%v", ip, port)
	log.FATAL("Router has stopped working", http.ListenAndServe(ip+":"+strconv.Itoa(port), router))
}
