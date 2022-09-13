package network

import (
	"D7024E/log"
	objectController "D7024E/network/controller/object"
	headerMiddleware "D7024E/network/middleware/header"
	loggingMiddleware "D7024E/network/middleware/logging"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	log.INFO("Initiated on %v", GetAddress())
	router := mux.NewRouter()

	objectController.Objects = append(objectController.Objects, objectController.Object{Name: "Test_Object", Hash: "TEST_HASH"})

	router.HandleFunc("/objects", objectController.CreateObject).Methods("POST")
	router.HandleFunc("/all/objects", objectController.GetObjects).Methods("GET")
	router.HandleFunc("/objects/{hash}", objectController.GetObject).Methods("GET")

	router.Use(headerMiddleware.HeaderMiddleware)

	logMiddleware := loggingMiddleware.NewLogMiddleware(log.InfoLogger)
	router.Use(logMiddleware.Start())

	log.FATAL("Router has stopped working", http.ListenAndServe(":8000", router))
}
