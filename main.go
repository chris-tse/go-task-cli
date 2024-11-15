package main

import (
	"go-task-cli/io"
	task "go-task-cli/models"
	"os"
)

// task item with
// id: A unique identifier for the task
// description: A short description of the task
// status: The status of the task (todo, in-progress, done)
// createdAt: The date and time when the task was created
// updatedAt: The date and time when the task was last updated

// enum for actions add, update, delete, mark-in-progress, mark-done, list

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		io.PrintHelp()
		os.Exit(0)
	}

	action := args[0]
	options := args[1:]

	task.PerformAction(action, options)

	os.Exit(0)
}
