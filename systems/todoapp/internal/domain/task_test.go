// internal/domain/task_test.go
package domain_test

import (
	"encoding/json"
	"testing"

	"github.com/USERNAME/cs-lab/systems/todoapp/internal/domain"
)

func TestTask_JSONRoundTrip(t *testing.T) {
	in := domain.Task{ID: 1, Title: "Test Task", Done: true}

	b, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var out domain.Task
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if out != in {
		t.Fatalf("mismatch: %#v vs %#v", out, in)
	}
}

func TestTask_ZeroValueDoneIsFalse(t *testing.T) {
	var task domain.Task
	if task.Done {
		t.Fatal("zero-value Done must be false")
	}
}
