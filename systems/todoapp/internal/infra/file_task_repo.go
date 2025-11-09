// Package infra provides infrastructure implementations for the todo application.
package infra

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"github.com/tomitahisaki/cs-lab/systems/todoapp/internal/domain"
)

// FileTaskRepo persists tasks into a JSON file.
// It keeps an in-memory index and flushes to disk on each mutation.
// Note: process-local concurrency only (no inter-process locking).
type FileTaskRepo struct {
	mu    sync.RWMutex
	path  string
	seq   int
	items map[int]*domain.Task
}

type fileImage struct {
	Seq   int           `json:"seq"`
	Tasks []domain.Task `json:"tasks"`
}

// NewFileTaskRepo creates/loads a repository persisted at `path`.
// If the file doesn't exist, it starts empty and creates it on first flush.
func NewFileTaskRepo(path string) (*FileTaskRepo, error) {
	r := &FileTaskRepo{
		path:  path,
		items: make(map[int]*domain.Task),
		seq:   0,
	}
	if err := r.load(); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *FileTaskRepo) NextID() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.seq++
	return r.seq
}

func (r *FileTaskRepo) Save(t *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[t.ID] = cloneTask(t)
	return r.flushLocked()
}

func (r *FileTaskRepo) FindAll() ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*domain.Task, 0, len(r.items))
	for _, v := range r.items {
		out = append(out, cloneTask(v))
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

func (r *FileTaskRepo) FindByID(id int) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.items[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return cloneTask(v), nil
}

func (r *FileTaskRepo) Update(t *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[t.ID]; !ok {
		return errors.Join(domain.ErrNotFound, fmt.Errorf("id=%d", t.ID))
	}
	r.items[t.ID] = cloneTask(t)
	return r.flushLocked()
}

// ---------- internal helpers ----------

func (r *FileTaskRepo) load() error {
	// If file doesn't exist yet, start empty.
	if _, err := os.Stat(r.path); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	data, err := os.ReadFile(r.path)
	if err != nil {
		return err
	}
	var img fileImage
	if len(data) == 0 {
		// Treat empty file as empty image.
		img = fileImage{}
	} else if err := json.Unmarshal(data, &img); err != nil {
		return fmt.Errorf("failed to parse %s: %w", r.path, err)
	}

	r.items = make(map[int]*domain.Task, len(img.Tasks))
	maxID := 0
	for i := range img.Tasks {
		t := img.Tasks[i] // value
		if t.ID > maxID {
			maxID = t.ID
		}
		tt := t // copy
		r.items[t.ID] = &tt
	}
	// seq: prefer persisted seq; fallback to max id
	if img.Seq > 0 {
		r.seq = img.Seq
	} else {
		r.seq = maxID
	}
	return nil
}

func (r *FileTaskRepo) flushLocked() error {
	// Must be called with r.mu locked (write).
	dir := filepath.Dir(r.path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	img := fileImage{
		Seq:   r.seq,
		Tasks: make([]domain.Task, 0, len(r.items)),
	}
	for _, v := range r.items {
		img.Tasks = append(img.Tasks, *v)
	}
	sort.Slice(img.Tasks, func(i, j int) bool { return img.Tasks[i].ID < img.Tasks[j].ID })

	b, err := json.MarshalIndent(&img, "", "  ")
	if err != nil {
		return err
	}

	tmp := r.path + ".tmp"
	if err := os.WriteFile(tmp, b, 0o644); err != nil {
		return err
	}
	// Atomic-ish replace on POSIX.
	return os.Rename(tmp, r.path)
}

func cloneTask(t *domain.Task) *domain.Task {
	if t == nil {
		return nil
	}
	c := *t
	return &c
}
