package main

import "fmt"

func init() {
	addCommand(Command{
		"help",
		showHelp,
		func(day string, args []string) error {
			fmt.Println("drevo - hierarchical task manager")
			fmt.Println("Usage:")
			for _, command := range commands {
				fmt.Print("\t")
				command.Help()
			}
			return nil
		},
	})
}

func showHelp() {
	fmt.Println("help - show help for subcommands")
}
