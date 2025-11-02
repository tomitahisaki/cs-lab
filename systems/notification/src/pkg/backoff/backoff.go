// Package backoff provides algorithms for calculating retry wait durations.
// It is typically used in systems that perform retries after transient failures,
// such as API rate limits or temporary network errors.
//
// The goal of a backoff strategy is to avoid overwhelming a remote service
// by spacing out retries over time. The delay grows after each failed attempt,
// usually following an exponential curve with optional random jitter.
//
// Example usage:
//
//	b := backoff.Exponential{
//	    BaseDelay:    500 * time.Millisecond,
//	    Factor:       2.0,
//	    JitterRatio:  0.3,
//	    MaxDelay:     30 * time.Second,
//	    RandomSource: rand.New(rand.NewSource(time.Now().UnixNano())),
//	}
//
//	delay := b.Duration(3) // third retry → e.g., 4 seconds
//	time.Sleep(delay)
package backoff

import (
	"math"
	"math/rand"
	"time"
)

// Exponential defines an exponential backoff strategy.
//
// Exponential backoff increases the delay between retries
// exponentially (BaseDelay * Factor^attempt) and optionally
// adds jitter (random variation) to avoid synchronization
// when many workers retry at the same time.
type Exponential struct {
	BaseDelay    time.Duration // How long to wait for the first retry
	Factor       float64       // How much to multiply the delay each time
	JitterRatio  float64       // Random variation (0.0–1.0)
	MaxDelay     time.Duration // Max delay limit
	RandomSource *rand.Rand    // Random number generator (for jitter)
}

// Duration returns how long to wait before the next retry.
//
// The input parameter `attempt` is the number of attempts already made.
// For example:
//
//	attempt = 0 → first retry (BaseDelay)
//	attempt = 1 → second retry (BaseDelay * Factor)
//	attempt = 2 → third retry (BaseDelay * Factor^2)
//
// and so on.
//
// It applies jitter and caps the delay at MaxDelay
func (e *Exponential) Duration(attempt int) time.Duration {
	if attempt <= 0 {
		attempt = 0
	}

	delay := float64(e.BaseDelay) * math.Pow(e.Factor, float64(attempt))

	// Cap delay at maximum
	if e.MaxDelay > 0 && time.Duration(delay) > e.MaxDelay {
		delay = float64(e.MaxDelay)
	}

	// Apply random jitter
	if e.JitterRatio > 0 && e.RandomSource != nil {
		jitter := (e.RandomSource.Float64()*2 - 1) * e.JitterRatio // -J..+J
		delay = delay * (1 + jitter)
	}

	if delay < 0 {
		delay = 0
	}

	return time.Duration(delay)
}
