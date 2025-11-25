package main

import (
	"fmt"
)

func init() {
	addCommand(Command{
		"uncomplete",
		uncompleteHelp,
		func(day string, args []string) error {
			var root string
			if len(args) > 2 {
				root = args[2]
			}
			return uncomplete(root)
		},
	})
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
