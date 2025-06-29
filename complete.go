package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

	dir, file := filepath.Split(arg)
	if strings.HasPrefix(file, "x ") {
		return nil
	}

	file = fmt.Sprintf("x %s %s", day, file)
	return os.Rename(arg, filepath.Join(dir, file))
}
