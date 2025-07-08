package main

import (
	"fmt"
)

func init() {
	helps = append(helps, uncompleteHelp)
}

func uncompleteHelp() {
	fmt.Println("uncomplete file - mark task as uncompleted")
}

func uncomplete(arg string) error {
	if arg == "" {
		uncompleteHelp()
		return nil
	}

	err := shouldBeATask(arg)
	if err != nil {
		return err
	}

	task := New(arg)
	return task.Uncomplete()
}
