package network

import (
	"D7024E/log"
	"D7024E/network/route"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	log.INFO("Initiated on %v", GetAddress())
	router := mux.NewRouter()
	route.RegisterRoutes(router)

	log.FATAL("Router has stopped working", http.ListenAndServe(":8000", router))
}
