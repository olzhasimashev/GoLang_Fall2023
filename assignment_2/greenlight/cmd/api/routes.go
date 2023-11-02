package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/blenders", app.createBlenderHandler)
	router.HandlerFunc(http.MethodGet, "/v1/blenders/:id", app.showBlenderHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/blenders/:id", app.updateBlenderHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/blenders/:id", app.deleteBlenderHandler)

	return router
}
