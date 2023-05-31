package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
    router := httprouter.New()
    standard := alice.New(app.logRequest)
    router.Handler(http.MethodGet, "/ping", standard.ThenFunc(ping))
    router.Handler(http.MethodPost, "/sign", standard.ThenFunc(app.payloadSign))
    return router
}
