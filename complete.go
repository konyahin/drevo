package main

import (
	"fmt"
)

func init() {
	helps = append(helps, completeHelp)
}

func completeHelp() {
	fmt.Println("complete file - mark task as completed")
}

func complete(day, arg string) error {
	if arg == "" {
		completeHelp()
		return nil
	}

	err := shouldBeATask(arg)
	if err != nil {
		return err
	}

	task := New(arg)
	return task.Complete(day)
}
