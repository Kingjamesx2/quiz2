//filename cmd/api/routes.go

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	//Create a new  httprouter ruter instance
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/info", app.createInfoHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/randomizeString", app.showRandomStringHandler)

	return router
}
