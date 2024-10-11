package main

import (
	"errors"
	"fmt"
	"github.com/code-chimp/htmx-go-example/internal/models"
	"github.com/code-chimp/htmx-go-example/internal/services"
	"github.com/code-chimp/htmx-go-example/internal/validator"
	"net/http"
	"strconv"
)

// getHome is a temporary handler to redirect users to the /contacts page.
func (app *application) getHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

// getContacts displays the contacts page.
func (app *application) getContacts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	contacts, err := app.contacts.GetAll(query)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.render(w, r, http.StatusOK, "contacts.index.go.tmpl", models.ContactsIndexVM{Contacts: contacts, Query: query})
}

// getContact displays a specific contact based on its ID.
func (app *application) getContact(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusNotFound)
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

	app.render(w, r, http.StatusOK, "contacts.view.go.tmpl", models.ContactsViewVM{Contact: contact})
}

// getNewContact displays the form for creating a new contact.
func (app *application) getNewContact(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "contacts.new.go.tmpl", models.ContactForm{})
}

// postNewContact processes the form for creating a new contact.
func (app *application) postNewContact(w http.ResponseWriter, r *http.Request) {
	form := models.ContactForm{}

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	validateContactForm(&form, app.contacts, 0)

	if !form.Valid() {
		app.render(w, r, http.StatusUnprocessableEntity, "contacts.new.go.tmpl", form)
		return
	}

	contact := models.Contact{
		First: form.First,
		Last:  form.Last,
		Phone: form.Phone,
		Email: form.Email,
	}

	err = app.contacts.Insert(&contact)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/contacts/%d", contact.ID), http.StatusSeeOther)
}

// getEditContact displays the form for editing a specific contact based on its ID.
func (app *application) getEditContact(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusNotFound)
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

	form := models.ContactForm{
		ID:    contact.ID,
		First: contact.First,
		Last:  contact.Last,
		Phone: contact.Phone,
		Email: contact.Email,
	}

	app.render(w, r, http.StatusOK, "contacts.edit.go.tmpl", form)
}

// postEditContact processes the form for editing a specific contact based on its ID.
func (app *application) postEditContact(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusNotFound)
		return
	}

	form := models.ContactForm{
		ID: id,
	}

	err = app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	validateContactForm(&form, app.contacts, id)

	if !form.Valid() {
		app.render(w, r, http.StatusUnprocessableEntity, "contacts.edit.go.tmpl", form)
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

	contact.First = form.First
	contact.Last = form.Last
	contact.Phone = form.Phone
	contact.Email = form.Email

	err = app.contacts.Update(contact)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/contacts/%d", contact.ID), http.StatusSeeOther)
}

// deleteContact deletes a specific contact based on its ID.
func (app *application) deleteContact(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusNotFound)
		return
	}

	err = app.contacts.Delete(id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

// validateContactForm validates the contact form fields.
func validateContactForm(form *models.ContactForm, repo *services.ContactRepository, id int) {
	form.CheckField(validator.NotBlank(form.Email), "Email", "Email is required.")
	form.CheckField(repo.EmailUnique(form.Email, id), "Email", "Email is already in use.")
	form.CheckField(validator.NotBlank(form.First), "First", "First name is required.")
	form.CheckField(validator.NotBlank(form.Last), "Last", "Last name is required.")
	form.CheckField(validator.NotBlank(form.Phone), "Phone", "Phone is required.")
}
