package rest

import (
	"D7024E/log"
	"D7024E/network/serverREST/route"
	"net/http"

	"github.com/gorilla/mux"
)

/**
 * Start a server with REST http structure.
 */
func Start() {
	log.INFO("Initiated on %v", GetAddress())
	router := mux.NewRouter()
	route.RegisterRoutes(router)

	log.FATAL("Router has stopped working", http.ListenAndServe(":8000", router))
}
