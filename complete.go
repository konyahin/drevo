package main

import (
	"errors"
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

	state, err := taskState(arg)
	if err != nil {
		return err
	}

	if state == DoesntExist {
		return errors.New("task doesn't exist: " + arg)
	}

	dir, file := filepath.Split(arg)
	if strings.HasPrefix(file, "x ") {
		return errors.New("task already completed: " + arg)
	}

	day := time.Now().Format(time.DateOnly)
	file = fmt.Sprintf("x %s %s", day, file)
	return os.Rename(arg, filepath.Join(dir, file))
}
