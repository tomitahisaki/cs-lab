// Package domain contains the core business entities of the todo application.
package domain

import "errors"

var ErrNotFound = errors.New("task not found")

// Task is a struct that represents a to-do task. it is core entity in the todo application.
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TaskRepository interface {
	NextID() int
	Save(task *Task) error
	FindAll() ([]*Task, error)
	FindByID(id int) (*Task, error)
	Update(task *Task) error
}
