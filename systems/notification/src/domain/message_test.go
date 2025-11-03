package domain

import (
	"errors"
	"testing"
)

func TestMessage_Validate_OK(t *testing.T) {
	msg := Message{
		ID:      "abc123",
		Channel: ChannelSlack,
		Text:    "Hello!",
	}

	if err := msg.Validate(); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestMessage_Validate_Errors(t *testing.T) {
	cases := []struct {
		name string
		msg  Message
	}{
		{"empty id", Message{Channel: ChannelSlack, Text: "hi"}},
		{"empty text", Message{ID: "1", Channel: ChannelSlack}},
		{"empty channel", Message{ID: "1", Text: "hi"}},
	}

	for _, errorCase := range cases {
		t.Run(errorCase.name, func(t *testing.T) {
			err := errorCase.msg.Validate()
			if err == nil {
				t.Fatalf("expected error for %s, got nil", errorCase.name)
			}

			var inv ErrInvalid
			if !errors.As(err, &inv) {
				t.Fatalf("expected ErrInvalid, got %T (%v)", err, err)
			}
		})
	}
}
