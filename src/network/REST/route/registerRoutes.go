package route

import (
	headerMiddleware "D7024E/network/middleware/header"
	loggingMiddleware "D7024E/network/middleware/logging"

	"github.com/gorilla/mux"
)

/**
 * Register the routes and middlewares to the router.
 */
func RegisterRoutes(router *mux.Router) {
	// router.HandleFunc("/objects", objectController.CreateObject).Methods("POST")
	// router.HandleFunc("/all/objects", objectController.GetObjects).Methods("GET")
	// router.HandleFunc("/objects/{hash}", objectController.GetObject).Methods("GET")

	router.Use(headerMiddleware.HeaderMiddleware)
	router.Use(loggingMiddleware.Start())
}
