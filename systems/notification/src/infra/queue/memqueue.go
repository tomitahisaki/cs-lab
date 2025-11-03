// Package queue is a simple in-memory FIFO queue with delayed re-enqueue support.
// it makes use of a min-heap to manage delayed messages efficiently. it is asynchronous
package queue

import (
	"container/heap"
	"context"
	"sync"
	"time"

	// <-- replace 'yourmodule' with your module path
	"github.com/USERNAME/cs-lab/systems/notification/src/domain"
)

// MemQueue is a simple in-memory FIFO queue with delayed re-enqueue support.
// It is safe for concurrent use.
type MemQueue struct {
	bufCap  int                 // How many messages can be buffered in the ready channel
	mu      sync.Mutex          // Guards access to delayed heap
	buf     chan domain.Message // Holds messages ready to be consumed immediately
	delayed delayedMinHeap      // Holds **features** messages scheduled for future delivery
	wakeup  chan struct{}       // Wakes up the scheduler goroutine when new delayed messages are added
	stopCh  chan struct{}       // Used to gracefully stop the background scheduler
	once    sync.Once           // Ensures `Close()` only runs once
}

// New creates a new MemQueue with the provided buffer capacity for ready messages.
// A capacity between 1-1024 is typically fine for local dev / tests.
func New(capacity int) *MemQueue {
	if capacity <= 0 {
		capacity = 1
	}
	q := &MemQueue{
		bufCap:  capacity,
		buf:     make(chan domain.Message, capacity),
		wakeup:  make(chan struct{}, 1),
		stopCh:  make(chan struct{}),
		delayed: delayedMinHeap{},
	}
	go q.scheduler()
	return q
}

// Enqueue inserts a message. If msg.NextAttemptAt is zero or due, it is made available
// to consumers immediately. Otherwise it's scheduled and delivered after NextAttemptAt.
func (q *MemQueue) Enqueue(msg domain.Message) {
	now := time.Now()
	// Treat zero time or past-due as immediate.
	if msg.NextAttemptAt.IsZero() || !msg.NextAttemptAt.After(now) {
		q.buf <- msg
		return
	}

	q.mu.Lock()
	heap.Push(&q.delayed, msg)
	q.mu.Unlock()
	q.notify()
}

// Dequeue blocks until a message is available or ctx is canceled.
// Returns (zero, false) on context cancelation; otherwise (msg, true).
func (q *MemQueue) Dequeue(ctx context.Context) (domain.Message, bool) {
	select {
	case <-ctx.Done():
		return domain.Message{}, false
	case m := <-q.buf:
		return m, true
	}
}

// Close stops the background scheduler goroutine. Safe to call multiple times.
// Primarily intended for tests / controlled shutdown.
func (q *MemQueue) Close() {
	q.once.Do(func() {
		close(q.stopCh)
		// Drain wakeup to allow scheduler to exit promptly if blocked on wakeup.
		q.notify()
	})
}

// Internal: wake the scheduler if it's waiting.
func (q *MemQueue) notify() {
	select {
	case q.wakeup <- struct{}{}:
	default:
		// already notified
	}
}

// Background loop: moves due items from delayed heap into the ready buffer.
func (q *MemQueue) scheduler() {
	timer := time.NewTimer(time.Hour) // reused timer; will be reset
	defer timer.Stop()

	for {
		// Determine next wake time (if any).
		q.mu.Lock()
		var waitDur time.Duration
		now := time.Now()

		if q.delayed.Len() == 0 {
			waitDur = -1 // indicates "no timer; wait on wakeup/stop"
		} else {
			next := q.delayed[0].NextAttemptAt
			if next.After(now) {
				waitDur = next.Sub(now)
			} else {
				waitDur = 0
			}
		}
		q.mu.Unlock()

		if waitDur < 0 {
			// No delayed items; wait for either a new item or stop.
			select {
			case <-q.wakeup:
				continue
			case <-q.stopCh:
				return
			}
		}

		if waitDur == 0 {
			// Move all due items to ready buffer.
			for {
				q.mu.Lock()
				if q.delayed.Len() == 0 || q.delayed[0].NextAttemptAt.After(time.Now()) {
					q.mu.Unlock()
					break
				}
				msg := heap.Pop(&q.delayed).(domain.Message)
				q.mu.Unlock()

				// Block if buf is full; this is acceptable for local dev/testing needs.
				select {
				case q.buf <- msg:
				case <-q.stopCh:
					return
				}
			}
			// Loop again to compute next wait or drain immediately due items.
			continue
		}

		// Positive wait: arm timer until the next due message.
		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
		timer.Reset(waitDur)

		select {
		case <-timer.C:
			// time to deliver due items (handled at next loop iteration)
		case <-q.wakeup:
			// new delayed item pushed; recompute wait
		case <-q.stopCh:
			return
		}
	}
}

// --- delayed heap ---

type delayedMinHeap []domain.Message

func (h delayedMinHeap) Len() int { return len(h) }
func (h delayedMinHeap) Less(i, j int) bool {
	// Earlier NextAttemptAt has higher priority.
	return h[i].NextAttemptAt.Before(h[j].NextAttemptAt)
}
func (h delayedMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *delayedMinHeap) Push(x any) {
	*h = append(*h, x.(domain.Message))
}

func (h *delayedMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
