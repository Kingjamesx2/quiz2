//FILename: cmd/api/routes

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	//Create a new  httprouter ruter instance
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/EntryInfo", app.createFileEntryHandler)
	router.HandlerFunc(http.MethodGet, "/v1/randomizeString/:id", app.showRandomizeStringHandler)

	return router
}
