package main

import (
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	dynamic := alice.New()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.getHome))
	mux.Handle("GET /contacts", dynamic.ThenFunc(app.getContacts))
	mux.Handle("GET /contacts/{id}", dynamic.ThenFunc(app.getContact))
	mux.Handle("GET /contacts/new", dynamic.ThenFunc(app.getNewContact))
	mux.Handle("POST /contacts/new", dynamic.ThenFunc(app.postNewContact))
	mux.Handle("GET /contacts/{id}/edit", dynamic.ThenFunc(app.getEditContact))
	mux.Handle("POST /contacts/{id}/edit", dynamic.ThenFunc(app.postEditContact))
	mux.Handle("POST /contacts/{id}/delete", dynamic.ThenFunc(app.deleteContact))

	baseMiddlewares := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return baseMiddlewares.Then(mux)
}
