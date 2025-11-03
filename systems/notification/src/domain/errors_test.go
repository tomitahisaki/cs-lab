package domain

import (
	"errors"
	"testing"
)

func TestErrInvalid(t *testing.T) {
	err := ErrInvalid{}.With("missing text")
	if err.Error() != "invalid message: missing text" {
		t.Fatalf("unexpected message: %v", err.Error())
	}

	var inv ErrInvalid
	if !errors.As(err, &inv) {
		t.Fatal("should unwrap to ErrInvalid")
	}
}
