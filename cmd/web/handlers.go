package main

import (
	"errors"
	"github.com/code-chimp/htmx-go-example/internal/models"
	"net/http"
	"strconv"
)

// getHome is a temporary handler to redirect users to the /contacts page.
func (app *application) getHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

// getContacts displays the contacts page.
func (app *application) getContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := app.contacts.GetAll()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Contacts = contacts

	app.render(w, r, http.StatusOK, "contacts.gohtml", data)
}

// getContact displays a specific contact based on its ID.
func (app *application) getContact(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	contact, err := app.contacts.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Contact = contact

	app.render(w, r, http.StatusOK, "contact.gohtml", data)
}
