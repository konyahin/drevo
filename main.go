package main

import (
	"fmt"
	"os"
)

type helpFunc func()

var helps []helpFunc

func showHelp() {
	fmt.Println("deltatree - hierarchical task manager")
	fmt.Println("Usage:")
	for _, f := range helps {
		fmt.Print("\t")
		f()
	}
}

func printTasks(tasks []string) {
	for _, task := range tasks {
		fmt.Println(task)
	}
}

func main() {
	if len(os.Args) < 2 {
		tasks, err := find(nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		printTasks(tasks)
		return
	}

	var err error
	switch os.Args[1] {
	case "help":
		showHelp()
	case "totree":
		err = totree(os.Args[2:])
	case "create":
		err = create(os.Args[2:])
	case "find":
		var tasks []string
		tasks, err = find(os.Args[2:])
		if err == nil {
			printTasks(tasks)
		}
	case "batch":
		var root string
		if len(os.Args) > 2 {
			root = os.Args[2]
		}
		err = batch(root)
	case "complete":
		var root string
		if len(os.Args) > 2 {
			root = os.Args[2]
		}
		err = complete(root)
	default:
		fmt.Println("unknown subcommand -", os.Args[1])
	}

	if err != nil {
		fmt.Println(err)
	}
}
