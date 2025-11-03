package domain

import "fmt"

// ErrInvalid represents a validation error for a domain entity.
// It includes a short cause message for clarity.
type ErrInvalid struct {
	Cause string
}

func (e ErrInvalid) Error() string {
	if e.Cause == "" {
		return "invalid message"
	}
	return fmt.Sprintf("invalid message: %s", e.Cause)
}

// With returns a new ErrInvalid with the given cause text.
// Example:
//
//	return ErrInvalid{}.With("id is empty")
func (e ErrInvalid) With(cause string) ErrInvalid {
	return ErrInvalid{Cause: cause}
}
