package main

import (
	"fmt"
	"os"
	"strings"
)

func init() {
	helps = append(helps, batchHelp)
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
