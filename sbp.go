// Package sbp provides functions for parsing and serializing of the Swift Navigation Binary Protocol (SBP).
package sbp

import (
	"errors"
)

var (
	// ErrInvalidFormat is returned when detect a malformed format.
	ErrInvalidFormat = errors.New("Invalid format")

	// ErrUnsupported is returned if the type or field is unsupported.
	ErrUnsupported = errors.New("Unsupported")
)
