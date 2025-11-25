package main

import (
	"fmt"
)

func init() {
	addCommand(Command{
		"complete",
		completeHelp,
		func(day string, args []string) error {
			var root string
			if len(args) > 2 {
				root = args[2]
			}
			return complete(day, root)
		},
	})
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
