package usecase

import (
	"errors"
	"testing"

	"github.com/tomitahisaki/cs-lab/cs/todoapp/internal/domain"
	"github.com/tomitahisaki/cs-lab/cs/todoapp/internal/infra"
)

func TestTaskUsecase_AddListDone(t *testing.T) {
	repo := infra.NewMemoryTaskRepo()
	uc := NewTaskUsecase(repo)

	// Add 2件
	a, err := uc.Add("Write tests")
	if err != nil {
		t.Fatal(err)
	}
	b, err := uc.Add("Ship code")
	if err != nil {
		t.Fatal(err)
	}

	// List
	all, err := uc.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(all) != 2 {
		t.Fatalf("List: want 2, got %d", len(all))
	}
	if all[0].ID != a.ID || all[1].ID != b.ID {
		t.Fatalf("List order mismatch: want [%d,%d], got [%d,%d]", a.ID, b.ID, all[0].ID, all[1].ID)
	}

	// Done
	if err := uc.Done(a.ID); err != nil {
		t.Fatal(err)
	}
	got, err := repo.FindByID(a.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !got.Done {
		t.Fatalf("Done not applied: task #%d Done=false", a.ID)
	}

	// Done は冪等
	if err := uc.Done(a.ID); err != nil {
		t.Fatal(err)
	}
}

func TestTaskUsecase_Done_NotFound(t *testing.T) {
	repo := infra.NewMemoryTaskRepo()
	uc := NewTaskUsecase(repo)

	if err := uc.Done(999); !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("Done: want ErrNotFound, got %v", err)
	}
}
