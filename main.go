package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var commands map[string]Command = make(map[string]Command)

func addCommand(command Command) {
	commands[command.Name] = command
}

func main() {
	if len(os.Args) < 2 {
		commands["find"].Call("", nil)
		return
	}

	day := time.Now().Format(time.DateOnly)

	var err error
	command, ok := commands[os.Args[1]]
	if ok {
		err = command.Call(day, os.Args)
	} else {
		fmt.Println("unknown subcommand -", os.Args[1])
	}

	if err != nil {
		log.Println(err)
	}
}
