package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.RecoverPanic, app.LogRequest, SecureHeaders)

	// Create a new middleware chain containing the middleware specific to
	// our dynamic application routes. For now, this chain will only contain
	// the session middleware but we'll add more to it later.
	dynamicMiddleware := alice.New(app.session.Enable, NoSurf, app.Authenticate)
	dynamicWithAuthenticationMiddleware := dynamicMiddleware.Append(app.RequireAuthenticatedUser)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.Home))
	mux.Get("/snippet/create", dynamicWithAuthenticationMiddleware.ThenFunc(app.CreateSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.ThenFunc(app.CreateSnippet))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.ShowSnippet))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.SignupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.SignupUser))

	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.LoginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.LoginUser))
	mux.Post("/user/logout", dynamicWithAuthenticationMiddleware.ThenFunc(app.LogoutUser))

	fileServer := http.FileServer(http.Dir("./web/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// Pass the servermux as the 'next' parameter to the secureHeaders middleware.
	// Because secureHeaders is just a function, and the function return a
	// http.Handler we don't need to do anything else.
	return standardMiddleware.Then(mux)
}
