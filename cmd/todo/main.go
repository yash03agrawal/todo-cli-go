package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/yash03agrawal/todoCliGo"
)

func main() {

	task := flag.String("task", "", "Task to be included in the todo list")
	complete := flag.Int("complete", 0, "Task index to be marked as complete")
	list := flag.Bool("list", false, "List all tasks present in todo list")

	flag.Parse()

	l := &todo.TaskList{}
	var todoFileName string = "todo.json"

	if err := l.RetrieveFromFile(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *task != "":
		l.AddTask(*task)
		if err := l.SaveToFile(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *complete > 0:
		if err := l.CompleteTask(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.SaveToFile(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *list:
		for _, item := range *l {
			if !item.IsDone {
				fmt.Println("Task: " + item.Task + " CreatedAt: " + item.CreatedAt.Format("2006-January-02"))
			}
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid options")
		os.Exit(1)
	}
}
