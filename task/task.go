package task

import "time"

// Task represents a single task
type Task struct {
	ID        string
	Title     string
	Completed bool
	DueDate   time.Time
}
