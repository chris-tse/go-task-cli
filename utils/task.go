package db

import (
	"encoding/json"
	"fmt"
	"os"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func createNewTaskDB(filename string) ([]Task, error) {
	tasks := []Task{}

	// convert to JSON
	data, err := json.Marshal(tasks)

	if err != nil {
		return nil, err
	}

	// write to new file
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func parseTaskDB(data []byte) ([]Task, error) {
	var tasks []Task
	err := json.Unmarshal(data, &tasks)

	// some error with parsing JSON
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func LoadTasks(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)

	// If "file not exists error" occurs, create a new tasks db file
	if os.IsNotExist(err) {
		return createNewTaskDB(filename)
	}

	// If different error occurs, return the error
	if err != nil {
		return nil, err
	}

	// File found, parse JSON into array
	return parseTaskDB(data)
}

func TestLoadOrCreateDB() {
	fmt.Println("testing load or create db")
}

func AddTask(description string) {
	tasks, err := LoadTasks("tasks.json")

	if err != nil {
		fmt.Println("Error loading tasks.json:", err)
		os.Exit(1)
	}

	tasks = append(tasks, Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   "2024-11-10",
		UpdatedAt:   "2024-11-10",
	})

	data, err := json.Marshal(tasks)

	if err != nil {
		fmt.Println("Error marshalling tasks:", err)
		os.Exit(1)
	}

	err = os.WriteFile("tasks.json", data, 0644)

	if err != nil {
		fmt.Println("Error writing tasks.json:", err)
		os.Exit(1)
	}
}
