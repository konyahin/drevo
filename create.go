package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
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

		if state != DoesntExist {
			return errors.New(arg + " already exist")
		}

		totree(slices.Collect(pathIter(filepath.Dir(arg))))

		f, err := os.Create(arg)
		if err != nil {
			return err
		}

		_ = f.Close()
	}

	return nil
}
