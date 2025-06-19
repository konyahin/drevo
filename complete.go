package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func init() {
	helps = append(helps, completeHelp)
}

func completeHelp() {
	fmt.Println("complete file - mark task as completed")
}

func complete(arg string) error {
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

	day := time.Now().Format(time.DateOnly)
	file = fmt.Sprintf("x %s %s", day, file)
	return os.Rename(arg, filepath.Join(dir, file))
}
