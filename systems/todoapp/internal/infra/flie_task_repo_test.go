// systems/todoapp/internal/infra/file_task_repo_test.go
package infra

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tomitahisaki/cs-lab/systems/todoapp/internal/domain"
)

func TestFileTaskRepo_PersistRoundtrip(t *testing.T) {
	tmp := t.TempDir()
	p := filepath.Join(tmp, "tasks.json")

	// 1st process: create & save
	r1, err := NewFileTaskRepo(p)
	if err != nil {
		t.Fatal(err)
	}
	id := r1.NextID()
	if err := r1.Save(&domain.Task{ID: id, Title: "A"}); err != nil {
		t.Fatal(err)
	}

	// 2nd process: reopen & read
	r2, err := NewFileTaskRepo(p)
	if err != nil {
		t.Fatal(err)
	}
	all, err := r2.FindAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(all) != 1 || all[0].Title != "A" {
		t.Fatalf("roundtrip mismatch: %+v", all)
	}

	// Update persists
	all[0].Done = true
	if err := r2.Update(all[0]); err != nil {
		t.Fatal(err)
	}

	r3, err := NewFileTaskRepo(p)
	if err != nil {
		t.Fatal(err)
	}
	got, _ := r3.FindByID(id)
	if !got.Done {
		t.Fatalf("update not persisted")
	}

	_ = os.RemoveAll(tmp)
}
