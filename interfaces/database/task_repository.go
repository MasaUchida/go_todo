package database

import "example/todo/domain"

type TaskRepository struct {
	SqlHandler
}

func (repo *TaskRepository) Store(t domain.Task) (err error) {
	_, err = repo.Execute(
		"INSERT INTO tasks (title, content) VALUES(?,?)", t.Title, t.Content,
	)
	if err != nil {
		return
	}
	return
}

func (repo *TaskRepository) FindById(Identifier int) (task domain.Task, err error) {
	row, err := repo.Query("SELECT id, title, content FROM tasks WHERE id = ?", Identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var (
		id      int
		title   string
		content string
	)
	row.Next()
	if err = row.Scan(&id, &title, &content); err != nil {
		return
	}
	task.ID = id
	task.Title = title
	task.Content = content
	return
}

func (repo *TaskRepository) FindAll() (tasks domain.Tasks, err error) {
	var (
		id      int
		title   string
		content string
	)
	rows, err := repo.Query("SELECT id, title, content FROM tasks")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&id, &title, &content); err != nil {
			return
		}
		task := domain.Task{
			ID:      id,
			Title:   title,
			Content: content,
		}
		tasks = append(tasks, task)
	}
	return
}
