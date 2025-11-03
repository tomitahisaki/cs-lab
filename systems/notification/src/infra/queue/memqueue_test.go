package queue

import (
	"context"
	"testing"
	"time"

	// <-- replace 'yourmodule' with your module path
	"github.com/USERNAME/cs-lab/systems/notification/src/domain"
)

func TestFIFOImmediate(t *testing.T) {
	q := New(8)
	defer q.Close()

	msgs := []domain.Message{
		{ID: "a"},
		{ID: "b"},
		{ID: "c"},
	}
	for _, m := range msgs {
		q.Enqueue(m)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	for _, expect := range []string{"a", "b", "c"} {
		got, ok := q.Dequeue(ctx)
		if !ok {
			t.Fatalf("expected message %q, got ctx cancel", expect)
		}
		if got.ID != expect {
			t.Fatalf("expected %q, got %q", expect, got.ID)
		}
	}
}

func TestDelayedDelivery(t *testing.T) {
	q := New(1)
	defer q.Close()

	delay := 120 * time.Millisecond
	q.Enqueue(domain.Message{
		ID:            "delayed",
		NextAttemptAt: time.Now().Add(delay),
	})

	// Try to dequeue too early — should timeout.
	earlyCtx, earlyCancel := context.WithTimeout(context.Background(), 40*time.Millisecond)
	defer earlyCancel()
	if _, ok := q.Dequeue(earlyCtx); ok {
		t.Fatalf("unexpectedly dequeued a message before it was due")
	}

	// Now wait past due time and dequeue.
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	got, ok := q.Dequeue(ctx)
	if !ok {
		t.Fatalf("expected to dequeue after delay")
	}
	if got.ID != "delayed" {
		t.Fatalf("expected 'delayed', got %q", got.ID)
	}
}

func TestMixedImmediateThenDelayed(t *testing.T) {
	q := New(2)
	defer q.Close()

	// One immediate, one delayed.
	q.Enqueue(domain.Message{ID: "immediate"})
	q.Enqueue(domain.Message{
		ID:            "later",
		NextAttemptAt: time.Now().Add(100 * time.Millisecond),
	})

	ctx1, cancel1 := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel1()

	first, ok := q.Dequeue(ctx1)
	if !ok {
		t.Fatalf("expected first message")
	}
	if first.ID != "immediate" {
		t.Fatalf("expected 'immediate' first, got %q", first.ID)
	}

	// Second should arrive after delay.
	ctx2, cancel2 := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel2()
	second, ok := q.Dequeue(ctx2)
	if !ok {
		t.Fatalf("expected second message")
	}
	if second.ID != "later" {
		t.Fatalf("expected 'later', got %q", second.ID)
	}
}

func TestDequeueContextCancel(t *testing.T) {
	q := New(1)
	defer q.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()

	if _, ok := q.Dequeue(ctx); ok {
		t.Fatalf("expected no message and ctx cancel")
	}
}
