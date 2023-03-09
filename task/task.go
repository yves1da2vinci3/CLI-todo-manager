package task

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Task represents a single task
type Task struct {
	gorm.Model
	Title     string
	Completed bool
	DueDate   time.Time `gorm:"type:date"`
}
