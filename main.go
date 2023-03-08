package main

import (
	"flag"
	"fmt"
	"time"

	"cli-todo/task"
)

func main() {
	var (
		add       = flag.String("add", "", "Add a new task")
		delete    = flag.Int("delete", 0, "Delete a task by ID")
		update    = flag.Int("update", 0, "Update a task by ID")
		title     = flag.String("title", "", "Title of the task")
		completed = flag.Bool("completed", false, "Is the task completed?")
		dueDate   = flag.String("due-date", "", "Due date of the task (YYYY-MM-DD)")
	)
	flag.Parse()
	manager := &task.InMemoryTaskManager{}

	if *add != "" {
		due, _ := time.Parse("2006-01-02", *dueDate)
		err := manager.Add(&task.Task{
			Title:     *title,
			Completed: *completed,
			DueDate:   due,
		})
		if err != nil {
			fmt.Println("Error adding task:", err)
			return
		}
		fmt.Println("Task added successfully")
	} else if *delete != 0 {
		err := manager.Delete(*delete)
		if err != nil {
			fmt.Println("Error deleting task:", err)
			return
		}
		fmt.Println("Task deleted successfully")
	} else if *update != 0 {
		due, _ := time.Parse("2006-01-02", *dueDate)
		err := manager.Update(&task.Task{
			ID:        *update,
			Title:     *title,
			Completed: *completed,
			DueDate:   due,
		})
		if err != nil {
			fmt.Println("Error updating task:", err)
			return
		}
		fmt.Println("Task updated successfully")
	} else {
		tasks, err := manager.List()
		if err != nil {
			fmt.Println("Error listing tasks:", err)
			return
		}
		fmt.Println("Tasks:")
		for _, task := range tasks {
			fmt.Printf("%d. %s (Due: %s)\n", task.ID, task.Title, task.DueDate.Format("2006-01-02"))
		}
	}
}
