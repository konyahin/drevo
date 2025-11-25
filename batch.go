package main

import (
	"fmt"
	"os"
	"strings"
)

func init() {
	addCommand(Command{
		"batch",
		batchHelp,
		func(day string, args []string) error {
			fileName, err := editTempFile()
			if err == nil {
				err = batch(day, fileName)
			}
			return err
		},
	})
}

func batchHelp() {
	fmt.Println("batch - create tasks from temp file, one task per line")
}

func batch(day string, fileName string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	tasks := strings.Split(string(content), "\n")
	return create(day, tasks)
}
