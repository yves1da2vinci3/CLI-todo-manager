package task

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"
)

// TaskManager is an interface for managing tasks
type TaskManager interface {
	Add(task *Task) error
	Delete(id int) error
	Update(task *Task) error
	Get(id int) (*Task, error)
	List() ([]*Task, error)
}

// InMemoryTaskManager is a TaskManager implementation that uses an in-memory data store
type InMemoryTaskManager struct {
	tasks []string
}

// Add adds a new task
func (m *InMemoryTaskManager) Add(task *Task) error {
	// generate a unique id for each task
	const idLength = 4
	idBytes := make([]byte, idLength)
	taskIdBytes := base64.URLEncoding.EncodeToString(idBytes)
	taskId := taskIdBytes[:idLength]
	// create a new file
	file, err := os.Create(taskId + ".txt")
	if err != nil {
		return err
	}
	defer file.Close()

	// write task to file
	writer := bufio.NewWriter(file)
	title := task.Title
	var completed string

	if task.Completed {
		completed = "true"
		strings.Builder()
	} else {
		completed = "false"
	}

	date := task.DueDate

	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}
	// save task to todos Directory

	return nil
}

// Delete deletes a task by ID
func (m *InMemoryTaskManager) Delete(id int) error {
	for i, task := range m.tasks {
		if task.ID == id {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return nil
		}
	}
	// delete the file

	return errors.New("task not found")
}

// Update updates an existing task
func (m *InMemoryTaskManager) Update(task *Task) error {
	for _, t := range m.tasks {
		if t.ID == task.ID {
			t.Title = task.Title
			t.Completed = task.Completed
			t.DueDate = task.DueDate
			return nil
		}
	}
	return errors.New("task not found")
}

// Get retrieves a task by ID
func (m *InMemoryTaskManager) Get(id string) (*Task, error) {
	file, err := os.Open(id + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read each line of the file
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// Check if there were any errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	for _, task := range m.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return nil, errors.New("task not found")
}

// List returns a list of all tasks
func (m *InMemoryTaskManager) List() ([]*Task, error) {
	return m.tasks, nil
}
