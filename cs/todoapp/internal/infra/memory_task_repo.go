// Package infra contains infrastructure implementations for the todo application.
package infra

import (
	"sort"
	"sync"

	"github.com/tomitahisaki/cs-lab/cs/todoapp/internal/domain"
)

type MemoryTaskRepo struct {
	mu    sync.RWMutex
	seq   int
	items map[int]*domain.Task // ★ ポインタを保持
}

func NewMemoryTaskRepo() *MemoryTaskRepo {
	return &MemoryTaskRepo{
		items: make(map[int]*domain.Task),
	}
}

func (r *MemoryTaskRepo) NextID() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.seq++
	return r.seq
}

// ★ インターフェイスに合わせてポインタ受け取り
func (r *MemoryTaskRepo) Save(t *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// そのまま保持（必要になったらディープコピーに差し替え可）
	r.items[t.ID] = t
	return nil
}

// ★ 戻り値も []*domain.Task に
func (r *MemoryTaskRepo) FindAll() ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]*domain.Task, 0, len(r.items))
	for _, v := range r.items {
		out = append(out, v)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

func (r *MemoryTaskRepo) FindByID(id int) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.items[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return v, nil
}

// ★ 引数もポインタ
func (r *MemoryTaskRepo) Update(t *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[t.ID]; !ok {
		return domain.ErrNotFound
	}
	r.items[t.ID] = t
	return nil
}
