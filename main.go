package main

import (
	"fmt"
	task "go-task-cli/utils"
	"log"
	"os"
	"text/tabwriter"
)

// task item with
// id: A unique identifier for the task
// description: A short description of the task
// status: The status of the task (todo, in-progress, done)
// createdAt: The date and time when the task was created
// updatedAt: The date and time when the task was last updated

// enum for actions add, update, delete, mark-in-progress, mark-done, list

const (
	AddAction            = "add"
	UpdateAction         = "update"
	DeleteAction         = "delete"
	MarkInProgressAction = "mark-in-progress"
	MarkDoneAction       = "mark-done"
	ListAction           = "list"
)

func printHelp() {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.TabIndent)

	fmt.Println("Usage: task-cli <command> [options]")
	fmt.Fprintln(w, "\nCommands:")
	fmt.Fprintln(w, "\tadd <task>\t\t\tAdd a new task")
	fmt.Fprintln(w, "\tupdate <id> <new task>\t\t\tUpdate an existing task")
	fmt.Fprintln(w, "\tdelete <id>\t\t\tDelete a task")
	fmt.Fprintln(w, "\tmark-in-progress <id>\t\t\tMark a task as in-progress")
	fmt.Fprintln(w, "\tmark-done <id>\t\t\tMark a task as done")
	fmt.Fprintln(w, "\tlist\t\t\tList all tasks")
	fmt.Fprintln(w, "\tlist <status>\t\t\tList tasks with the given status")
	w.Flush()
}

func performAction(action string, db []task.Task, options []string) {
	switch action {
	case ListAction:
		if len(db) == 0 {
			fmt.Println("No tasks found")
			fmt.Println("Add a task with \n\ttask-cli add <task>")
			os.Exit(0)
		}

		for _, item := range db {
			fmt.Println(item.ID, "\t", item.Description, "\t", item.Status)
		}

		os.Exit(0)
	case AddAction:
		if len(options) == 0 {
			fmt.Println("No task provided")
			fmt.Println("Usage: \n\ttask-cli add <task>")
		}

		task.AddTask(options[0])
	case UpdateAction:
		fmt.Println("update")
	case DeleteAction:
		fmt.Println("delete")
	case MarkInProgressAction:
		fmt.Println("mark-in-progress")
	case MarkDoneAction:
		fmt.Println("mark-done")
	default:
		fmt.Println("Unknown action")
		printHelp()
	}

	os.Exit(0)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		printHelp()
		os.Exit(0)
	}

	action := args[0]
	options := args[1:]

	db, err := task.LoadTasks("tasks.json")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	performAction(action, db, options)

	os.Exit(0)
}
