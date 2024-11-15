package io

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func PrintHelp() {
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
