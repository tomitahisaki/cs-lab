package backoff

import (
	"math/rand"
	"testing"
	"time"
)

func TestDuration_GrowsExponentially(t *testing.T) {
	exponential := Exponential{
		BaseDelay:    500 * time.Millisecond,
		Factor:       2.0,
		JitterRatio:  0.0,
		MaxDelay:     0,
		RandomSource: rand.New(rand.NewSource(1)),
	}

	tests := []struct {
		attempt int
		want    time.Duration
	}{
		{0, 500 * time.Millisecond},
		{1, 1 * time.Second},
		{2, 2 * time.Second},
		{3, 4 * time.Second},
	}

	for _, test := range tests {
		got := exponential.Duration(test.attempt)
		if got != test.want {
			t.Errorf("Duration(%d) = %v; want %v", test.attempt, got, test.want)
		}
	}
}

func TestDuration_CappedByMaxDelay(t *testing.T) {
	exponential := Exponential{
		BaseDelay:    1 * time.Second,
		Factor:       3.0,
		JitterRatio:  0.0,
		MaxDelay:     5 * time.Second,
		RandomSource: rand.New(rand.NewSource(1)),
	}

	got := exponential.Duration(3) // 1s * 3^3 = 27s → should be capped at 5s
	want := exponential.MaxDelay   // expected cap value

	if got != want {
		t.Errorf("expected delay to be capped at %v, got %v", want, got)
	}
}

func TestDuration_JitterWithinRange(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	exponential := Exponential{
		BaseDelay:    1 * time.Second,
		Factor:       1.0,
		JitterRatio:  0.3,
		MaxDelay:     0,
		RandomSource: r,
	}

	for range 50 {
		delay := exponential.Duration(0)
		if delay < 700*time.Millisecond || delay > 1300*time.Millisecond {
			t.Fatalf("jitter out of range: got %v", delay)
		}
	}
}

func TestDuration_NegativeAttempt(t *testing.T) {
	exponential := Exponential{
		BaseDelay:    1 * time.Second,
		Factor:       2.0,
		JitterRatio:  0.0,
		RandomSource: rand.New(rand.NewSource(1)),
	}

	got := exponential.Duration(-3)
	want := 1 * time.Second // should default to attempt=0
	if got != want {
		t.Errorf("negative attempt fallback: want %v, got %v", want, got)
	}
}
