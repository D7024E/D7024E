package route

import (
	"D7024E/network/server_rest/controller/post"
	headerMiddleware "D7024E/network/server_rest/middleware/header"
	loggingMiddleware "D7024E/network/server_rest/middleware/logging"

	"github.com/gorilla/mux"
)

/**
 * Register the routes and middlewares to the router.
 */
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/objects", post.Objects).Methods("POST")

	router.Use(headerMiddleware.HeaderMiddleware)
	router.Use(loggingMiddleware.Start())
}
