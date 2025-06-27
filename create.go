package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func init() {
	helps = append(helps, createHelp)
}

func createHelp() {
	fmt.Println("create file ... - create new task(s)")
}

func create(day string, args []string) error {
	if len(args) < 1 {
		createHelp()
		return nil
	}

	for _, arg := range args {
		if arg == "" {
			continue
		}

		dir, file := filepath.Split(arg)
		file = fmt.Sprintf("%s %s", day, file)
		fullPath := filepath.Join(dir, file)

		state, err := taskState(fullPath)
		if err != nil {
			return err
		}

		if state == NotATask {
			return errors.New("taks path contain file (not a folder): " + fullPath)
		}

		// already exist - do nothing
		if state == Ok {
			return nil
		}

		err = os.MkdirAll(fullPath, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
