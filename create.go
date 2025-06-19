package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"
)

func init() {
	helps = append(helps, createHelp)
}

func createHelp() {
	fmt.Println("create file ... - create new task(s)")
}

func create(args []string) error {
	if len(args) < 1 {
		createHelp()
		return nil
	}

	for _, arg := range args {
		if arg == "" {
			continue
		}

		state, err := taskState(arg)
		if err != nil {
			return err
		}

		if state == NotATask {
			return errors.New("taks path contain file (not a folder): " + arg)
		}

		if state == Ok {
			return errors.New("task already exist: " + arg)
		}

		dir, file := filepath.Split(arg)
		day := time.Now().Format(time.DateOnly)
		file = fmt.Sprintf("%s %s", day, file)

		err = fm.CreateTask(filepath.Join(dir, file))
		if err != nil {
			return err
		}
	}

	return nil
}
