package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Pass the servermux as the 'next' parameter to the secureHeaders middleware.
	// Because secureHeaders is just a function, and the function return a
	// http.Handler we don't need to do anything else.

	// return app.logRequest(secureHeaders(mux))
	return standardMiddleware.Then(mux)
}
