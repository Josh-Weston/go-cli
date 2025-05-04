package main

import (
	"fmt"
	"os"
	"strings"

	"josh-weston.com/ch2/todo"
)

const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1) // Exit with a code so other scripts knowe that the tool did not execute correctly
	}

	switch {
	// User did not provide any arguments, list the todo items
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		// All other arguments are expected to be the names of new tasks
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

}
