package domain

import "context"

// Notifier defines the interface for any outbound notification channel.
// Implementations should convert transport-level or API errors into
// domain-level errors: ErrRetryable (temporary) or ErrPermanent (non-retryable).
type Notifier interface {
	// Send delivers the given Message through the channel.
	// It must be safe for concurrent use.
	//
	// Returns:
	//   - nil on success
	//   - ErrRetryable on transient failure (e.g. network, rate-limit)
	//   - ErrPermanent on unrecoverable failure (e.g. invalid payload)
	Send(ctx context.Context, msg Message) error
}
