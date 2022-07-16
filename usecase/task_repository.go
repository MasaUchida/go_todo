package usecase

import "example/todo/domain"

type TaskRepository interface {
	Store(task domain.Task) error
	FindAll() (domain.Tasks, error)
	FindById(int) (domain.Task, error)
}
