package infra

import (
	"errors"
	"sync"
	"testing"

	"github.com/tomitahisaki/cs-lab/cs/todoapp/internal/domain"
)

// コンパイル時に TaskRepository を実装していることを保証
var _ domain.TaskRepository = (*MemoryTaskRepo)(nil)

func TestMemoryTaskRepo_BasicCRUD(t *testing.T) {
	r := NewMemoryTaskRepo()

	// NextID は連番
	id1 := r.NextID()
	id2 := r.NextID()
	if id2 != id1+1 {
		t.Fatalf("NextID must be monotonic: got %d then %d", id1, id2)
	}

	// Save → FindByID
	a := &domain.Task{ID: id1, Title: "A"}
	b := &domain.Task{ID: id2, Title: "B"}
	if err := r.Save(a); err != nil {
		t.Fatal(err)
	}
	if err := r.Save(b); err != nil {
		t.Fatal(err)
	}

	gotA, err := r.FindByID(id1)
	if err != nil {
		t.Fatal(err)
	}
	if gotA.Title != "A" {
		t.Fatalf("FindByID: want A, got %s", gotA.Title)
	}

	// FindAll は ID 昇順
	all, err := r.FindAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(all) != 2 {
		t.Fatalf("FindAll: want 2, got %d", len(all))
	}
	if all[0].ID != id1 || all[1].ID != id2 {
		t.Fatalf("FindAll order: want [%d,%d], got [%d,%d]", id1, id2, all[0].ID, all[1].ID)
	}

	// Update 成功
	gotA.Title = "A2"
	if err := r.Update(gotA); err != nil {
		t.Fatal(err)
	}
	gotA2, _ := r.FindByID(id1)
	if gotA2.Title != "A2" {
		t.Fatalf("Update not applied: got %s", gotA2.Title)
	}
}

func TestMemoryTaskRepo_NotFound(t *testing.T) {
	r := NewMemoryTaskRepo()

	if _, err := r.FindByID(999); !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("FindByID: want ErrNotFound, got %v", err)
	}
	if err := r.Update(&domain.Task{ID: 999, Title: "X"}); !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("Update: want ErrNotFound, got %v", err)
	}
}

func TestMemoryTaskRepo_ConcurrentAccess(t *testing.T) {
	r := NewMemoryTaskRepo()

	var wg sync.WaitGroup
	const n = 200

	// 複数 goroutine から同時に NextID + Save
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := r.NextID()
			_ = r.Save(&domain.Task{ID: id, Title: "t"})
		}()
	}
	wg.Wait()

	all, err := r.FindAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(all) != n {
		t.Fatalf("concurrent save count mismatch: want %d, got %d", n, len(all))
	}

	// 同時に FindByID/Update をざっくり叩く
	wg = sync.WaitGroup{}
	for _, tk := range all {
		wg.Add(1)
		go func(tk *domain.Task) {
			defer wg.Done()
			_, _ = r.FindByID(tk.ID)
			tk.Title = "updated"
			_ = r.Update(tk)
		}(tk)
	}
	wg.Wait()
}
