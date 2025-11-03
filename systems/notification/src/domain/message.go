// Package domain holds core entities and interfaces for the notification system.
// It is intentionally pure (standard library only) and contains no infrastructure logic.
package domain

import "time"

// Channel identifies a delivery channel (e.g., "slack", "email").
type Channel string

const (
	ChannelSlack Channel = "slack"
	// Extend here later: ChannelEmail, ChannelLINE, ...
)

// Message is the unit of work for outbound notifications.
type Message struct {
	ID            string            // Unique identifier (uuid string recommended)
	Channel       Channel           // Target channel
	Text          string            // Message content (MVP: plain text)
	Metadata      map[string]string // Optional extra info (trace_id, tags, etc.)
	AttemptCount  int               // Number of attempts already made
	NextAttemptAt time.Time         // When to try next (set by dispatcher)
}

// Validate returns an error if the message is not suitable for dispatch.
func (m Message) Validate() error {
	if m.ID == "" {
		return ErrInvalid{}.With("id is empty")
	}
	if m.Text == "" {
		return ErrInvalid{}.With("text is empty")
	}
	if m.Channel == "" {
		return ErrInvalid{}.With("channel is empty")
	}
	return nil
}
