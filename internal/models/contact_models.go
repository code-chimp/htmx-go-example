package models

import "github.com/code-chimp/htmx-go-example/internal/validator"

// Contact represents a contact persisted to storage.
type Contact struct {
	ID    int    `json:"id"`
	First string `json:"first"`
	Last  string `json:"last"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

// ContactsIndexVM represents a view model containing multiple contacts.
type ContactsIndexVM struct {
	Contacts []*Contact
	Query    string
}

// ContactsViewVM represents a view model containing a single contact.
type ContactsViewVM struct {
	Contact *Contact
}

// ContactForm represents a form for creating or updating a contact.
type ContactForm struct {
	ID                  int    `form:"-"`
	First               string `form:"first"`
	Last                string `form:"last"`
	Phone               string `form:"phone"`
	Email               string `form:"email"`
	validator.Validator `form:"-"`
}
