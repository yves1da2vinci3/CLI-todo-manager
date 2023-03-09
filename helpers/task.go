package helpers

import (
	"cli-todo/database"
	"cli-todo/task"
	"errors"
	"fmt"
)

func GetTodos() []task.Task {
	db := database.DBConn
	var todos []task.Task
	db.Find(&todos)
	return todos
}

func GetTodo(id int) task.Task {
	db := database.DBConn
	var todo task.Task
	db.First(&todo, id)
	return todo

}

func NewTodo(todo *task.Task) error {
	db := database.DBConn
	newTask := new(task.Task)
	newTask.Completed = todo.Completed
	newTask.DueDate = todo.DueDate
	newTask.Title = todo.Title
	db.Create(&newTask)

	return nil

}

func DeleteTodo(id int) error {
	db := database.DBConn
	var todo task.Task
	db.First(&todo, id)
	if todo.Title == "" {
		return errors.New("todo not found")
	}

	db.Delete(task.Task{}, "id =?", id)
	fmt.Println("Todo deleted")
	return nil

}
func UpdateTodo(id int, completed bool) error {
	db := database.DBConn
	var todo task.Task
	db.First(&todo, id)
	if todo.Title == "" {
		return errors.New("todo not found")
	}
	todo.Completed = completed
	db.Save(&todo)

	fmt.Println("Todo updated successfully")
	return nil

}
