package main

import (
	"flag"
	"fmt"
	"time"

	"cli-todo/database"
	"cli-todo/helpers"
	"cli-todo/task"

	"github.com/jinzhu/gorm"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "todos.db")
	if err != nil {
		panic(err)
	}
	// only use for the first run
	fmt.Println("Database initialized ")
	database.DBConn.AutoMigrate(&task.Task{})
	fmt.Println("Database migrated")
}

func main() {
	initDatabase()

	defer database.DBConn.Close()
	var (
		add            = flag.String("add", "", "Add a new task")
		delete         = flag.Int("delete", 0, "Delete a task by ID")
		see            = flag.Int("see", 0, "get a task by ID")
		update         = flag.Int("update", 0, "Update a task by ID")
		completed      = flag.Bool("completed", false, "Is the task completed?")
		completionDate = flag.String("completionDate", "", "Due date of the task (YYYY-MM-DD)")
	)
	flag.Parse()

	if *add != "" {
		fmt.Println("date", *completionDate)
		dateString := *completionDate
		date, _ := time.Parse("2006-01-02", dateString)
		fmt.Println(date)
		err := helpers.NewTodo(&task.Task{Title: *add, Completed: *completed, DueDate: date})

		if err != nil {
			fmt.Println("Error adding task:", err)
			return
		}
		fmt.Println("Task added successfully")
	} else if *delete != 0 {
		err := helpers.DeleteTodo(*delete)
		if err != nil {
			fmt.Println("Error deleting task:", err)
			return
		}
		fmt.Println("Task deleted successfully")
	} else if *update != 0 {
		err := helpers.UpdateTodo(*update, *completed)
		if err != nil {
			fmt.Println("Error updating task:", err)
			return
		}
		fmt.Println("Task updated successfully")
	} else if *see != 0 {

		task := helpers.GetTodo(*see)
		fmt.Printf("%d. %s (Due date: %s) status : %v \n", task.ID, task.Title, task.DueDate, task.Completed)
	} else {
		tasks := helpers.GetTodos()

		fmt.Println("Tasks:")
		for _, task := range tasks {
			fmt.Printf("%d. %s (Due date: %s) status : %v \n", task.ID, task.Title, task.DueDate, task.Completed)
		}
	}
}
