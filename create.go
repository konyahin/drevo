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

func create(args []string) error {
	if len(args) < 1 {
		createHelp()
		return nil
	}

	for _, arg := range args {
		state, err := taskState(arg)
		if err != nil {
			return err
		}

		if state != DoesntExist {
			return errors.New(arg + " already exist")
		}

		dirPath := filepath.Dir(arg)
		if dirPath != "." {
			err = os.MkdirAll(dirPath, 0750)
			if err != nil {
				return err
			}
		}

		f, err := os.Create(arg)
		if err != nil {
			return err
		}

		_ = f.Close()
	}

	return nil
}
