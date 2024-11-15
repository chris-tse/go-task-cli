package task

import (
	"encoding/json"
	"fmt"
	io "go-task-cli/io"
	"os"
)

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

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

func Read(filename string) ([]Task, error) {
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
	return io.ParseJSONFile[Task](data)
}

func WriteTasks(filename string, tasks []Task) error {
	data, err := json.Marshal(tasks)

	if err != nil {
		return err
	}

	return io.WriteToFile(filename, data)
}

func TestLoadOrCreateDB() {
	fmt.Println("testing load or create db")
}

func Add(description string, taskList []Task) {

	taskList = append(taskList, Task{
		ID:          len(taskList) + 1,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   "2024-11-10",
		UpdatedAt:   "2024-11-10",
	})

	data, err := json.Marshal(taskList)

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
