package usecase

import (
	"github.com/tomitahisaki/cs-lab/systems/todoapp/internal/domain"
)

type TaskUsecase struct {
	repo domain.TaskRepository
}

func NewTaskUsecase(r domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{repo: r}
}

// Add: タスクを追加（現段階はバリデーションなし）
func (u *TaskUsecase) Add(title string) (*domain.Task, error) {
	id := u.repo.NextID()
	t := &domain.Task{ID: id, Title: title, Done: false}
	if err := u.repo.Save(t); err != nil {
		return nil, err
	}
	return t, nil
}

// List: すべてのタスクを取得（ID昇順はRepoが担保）
func (u *TaskUsecase) List() ([]*domain.Task, error) {
	return u.repo.FindAll()
}

// Done: 指定IDを完了にする
func (u *TaskUsecase) Done(id int) error {
	t, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}
	if t.Done {
		return nil // 既に完了なら何もしない（冪等）
	}
	t.Done = true
	return u.repo.Update(t)
}
