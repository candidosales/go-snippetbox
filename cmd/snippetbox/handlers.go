package main

import (
	"candidosg/snippetbox/pkg/forms"
	"candidosg/snippetbox/pkg/models"
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	s, err := app.snippets.Latest()
	if err != nil {
		app.ServerError(w, err)
		return
	}
	app.Render(w, r, "home.page.tmpl", &templateData{Snippets: s})
}

func (app *application) ShowSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}
	// Use the SnippetModel object's Get method to retrieve the data for a
	// specific record based on its ID. If no matching record is found,
	// return a 404 Not Found response.
	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.NotFound(w)
		return
	} else if err != nil {
		app.ServerError(w, err)
		return
	}

	app.Render(w, r, "show.page.tmpl", &templateData{Snippet: s})
}

func (app *application) CreateSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.Render(w, r, "create.page.tmpl", &templateData{Form: forms.New(nil)})
}

func (app *application) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	// First we call r.ParseForm() which adds any data in POST request bodies
	// to the r.PostForm map. This also works in the same way for PUT and PATCH
	// requests. If there are any errors, we use our app.ClientError helper to send
	// a 400 Bad Request response to the user.

	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	// Create a new forms.Form struct containing the POSTed data from the
	// form, then use the validation methods to check the content
	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		app.Render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}

	// Create a new snippet record in the database using the form data.
	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.ServerError(w, err)
		return
	}

	// Use the Put() method to add a string value ("Your snippet was saved
	// successfully!") and the corresponding key ("flash") to the session
	// data. Note that if there's no existing session for the current user
	// (or their session has expired) then a new, empty, session for them
	// will automatically be created by the session middleware.
	app.session.Put(r, "flash", "Snippet successfully created!")

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (app *application) SignupUserForm(w http.ResponseWriter, r *http.Request) {
	app.Render(w, r, "signup.page.tmpl", &templateData{Form: forms.New(nil)})
}

func (app *application) SignupUser(w http.ResponseWriter, r *http.Request) {
	// Parse the form data.
	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	// Validate the form contents using the form helper we made earlier.
	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 10)
	// If there are any errors, redisplay the signup form.
	if !form.Valid() {
		app.Render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	}

	// Try to create a new user record in the database. If the email already exists
	// add an error message to the form and re-display it.
	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
	if err == models.ErrDuplicateEmail {
		form.Errors.Add("email", "Address is already in use")
		app.Render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	} else if err != nil {
		app.ServerError(w, err)
		return
	}

	// Otherwise add a confirmation flash message to the session confirming that
	// their signup worked and asking them to log in.
	app.session.Put(r, "flash", "Your signup was successful. Please log in.")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) LoginUserForm(w http.ResponseWriter, r *http.Request) {
	app.Render(w, r, "login.page.tmpl", &templateData{Form: forms.New(nil)})
}

func (app *application) LoginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	// Check whether the credentials are valid. If they're not, add a generic error
	// message to the form failures map and re-display the login page.
	form := forms.New(r.PostForm)
	id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))

	if err == models.ErrInvalidCredentials {
		form.Errors.Add("generic", "Email or Password is incorrect")
		app.Render(w, r, "login.page.tmpl", &templateData{Form: form})
		return
	} else if err != nil {
		app.ServerError(w, err)
		return
	}

	// Add the ID of the current user to the session, so that they are now 'logged
	// in'.
	app.session.Put(r, "userID", id)

	http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
}

func (app *application) LogoutUser(w http.ResponseWriter, r *http.Request) {
	// Remove the userID from the session data so that the user is 'logged out'.
	app.session.Remove(r, "userID")

	// Add a flash message to the session to confirm to the user that they've been logged out.
	app.session.Put(r, "flash", "You've been logged out successfully!")

	http.Redirect(w, r, "/", 303)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
