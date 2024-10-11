package validator

import (
	"slices"
	"strings"
	"unicode/utf8"
)

// Validator represents a validation object that holds validation errors.
type Validator struct {
	Errors map[string]string // A map to store validation errors.
}

// Valid checks if there are no validation errors.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError adds an error message for a specific field.
func (v *Validator) AddError(field, message string) {
	if v.Errors == nil {
		v.Errors = make(map[string]string)
	}

	if _, ok := v.Errors[field]; ok {
		v.Errors[field] = v.Errors[field] + ", " + message
		return
	}

	v.Errors[field] = message
}

// CheckField adds an error message if the validation check fails.
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// NotBlank checks if a string is not blank (i.e., it contains non-whitespace characters).
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MaxChars checks if a string contains no more than a specified number of characters.
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// PermittedValue checks if a value is within a list of permitted values.
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}
