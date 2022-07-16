package domain

import (
	"time"
)

type Task struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
}

type Tasks []Task
