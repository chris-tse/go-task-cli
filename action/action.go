package action

import (
	"fmt"
	"go-task-cli/io"
	task "go-task-cli/models"
	"os"
)

type Action string

const (
	Add            Action = "add"
	Update         Action = "update"
	Delete         Action = "delete"
	MarkInProgress Action = "mark-in-progress"
	MarkDone       Action = "mark-done"
	List           Action = "list"
)

func (a Action) isValid() bool {
	switch a {
	case Add, Update, Delete,
		MarkInProgress, MarkDone, List:
		return true
	}
	return false
}

func PerformAction(action string, options []string) {
	parsedAction := Action(action)

	if !parsedAction.isValid() {
		fmt.Println("Unknown action")
		io.PrintHelp()
		os.Exit(0)
		return
	}

	db, err := task.Read("tasks.json")

	if err != nil {
		fmt.Println("Error loading tasks.json:", err)
		os.Exit(1)
	}

	switch parsedAction {

	case List:
		if len(db) == 0 {
			fmt.Println("No tasks found")
			fmt.Println("Add a task with \n\ttask-cli add <task>")
			os.Exit(0)
		}

		for _, item := range db {
			fmt.Println(item.ID, "\t", item.Description, "\t", item.Status)
		}

	case Add:
		if len(options) == 0 {
			fmt.Println("No task provided")
			fmt.Println("Usage: \n\ttask-cli add <task>")
		}

		task.Add(options[0], db)
	case Update:
		fmt.Println("update")
	case Delete:
		fmt.Println("delete")
	case MarkInProgress:
		fmt.Println("mark-in-progress")
	case MarkDone:
		fmt.Println("mark-done")
	default:
		// default already handled in isValid()
	}

	WriteTasks("tasks.json", db)
	os.Exit(0)
}
