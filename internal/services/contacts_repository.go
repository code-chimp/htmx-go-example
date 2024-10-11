package services

import (
	"encoding/json"
	"errors"
	"github.com/code-chimp/htmx-go-example/internal/models"
	"os"
	"strings"
)

// ContactRepository manages a collection of contacts.
type ContactRepository struct {
	contacts []*models.Contact
}

// NewRepository creates a new ContactRepository from the data in the contacts.json file.
// It reads the JSON file, unmarshal the data into a slice of Contact structs, and returns a new ContactRepository.
// Returns an error if the file cannot be read or the JSON cannot be unmarshalled.
func NewRepository() (*ContactRepository, error) {
	file, err := os.Open("./data/contacts.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var contacts []*models.Contact
	if err := json.NewDecoder(file).Decode(&contacts); err != nil {
		return nil, err
	}

	return &ContactRepository{contacts: contacts}, nil
}

// saveToFile writes the current state of the contacts slice to the contacts.json file.
// Returns an error if the file cannot be created or the JSON cannot be marshaled.
func (r *ContactRepository) saveToFile() error {
	file, err := os.Create("./data/contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(r.contacts)
}

// getNextID returns the next available ID for a new contact.
func (r *ContactRepository) getNextID() int {
	maxID := 0
	for _, contact := range r.contacts {
		if contact.ID > maxID {
			maxID = contact.ID
		}
	}
	return maxID + 1
}

// Get returns a contact by ID if found, or an error if not found.
func (r *ContactRepository) Get(id int) (*models.Contact, error) {
	for _, c := range r.contacts {
		if c.ID == id {
			return c, nil
		}
	}
	return nil, errors.New("contact not found")
}

// GetAll returns all contacts in the repository. If a query string is provided, it filters the contacts
// whose Email, First, or Last includes the query string (case insensitive).
func (r *ContactRepository) GetAll(query ...string) ([]*models.Contact, error) {
	if len(query) == 0 || query[0] == "" {
		return r.contacts, nil
	}

	q := strings.ToLower(query[0])
	var filteredContacts []*models.Contact

	for _, c := range r.contacts {
		if strings.Contains(strings.ToLower(c.Email), q) ||
			strings.Contains(strings.ToLower(c.First), q) ||
			strings.Contains(strings.ToLower(c.Last), q) ||
			strings.Contains(strings.ToLower(c.Phone), q) {
			filteredContacts = append(filteredContacts, c)
		}
	}

	return filteredContacts, nil
}

// Insert adds a new contact to the repository and persists the change to the contacts.json file.
// Returns the inserted contact and an error if the file cannot be saved.
func (r *ContactRepository) Insert(contact *models.Contact) error {
	contact.ID = r.getNextID()
	r.contacts = append(r.contacts, contact)

	err := r.saveToFile()
	if err != nil {
		return err
	}

	return nil
}

// Update modifies an existing contact in the repository and persists the change to the contacts.json file.
// Returns an error if the contact is not found or the file cannot be saved.
func (r *ContactRepository) Update(contact *models.Contact) error {
	for i, c := range r.contacts {
		if c.ID == contact.ID {
			r.contacts[i] = contact
			return r.saveToFile()
		}
	}
	return errors.New("contact not found")
}

// Delete removes a contact from the repository by ID and persists the change to the contacts.json file.
// Returns an error if the contact is not found or the file cannot be saved.
func (r *ContactRepository) Delete(id int) error {
	for i, c := range r.contacts {
		if c.ID == id {
			r.contacts = append(r.contacts[:i], r.contacts[i+1:]...)
			return r.saveToFile()
		}
	}
	return errors.New("contact not found")
}

// EmailUnique checks if a contact with the same email address already exists in the repository.
func (r *ContactRepository) EmailUnique(email string, id int) bool {
	for _, c := range r.contacts {
		if c.Email == email && c.ID != id {
			return false
		}
	}
	return true
}
