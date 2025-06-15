package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
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

		if state != DoesntExist {
			return errors.New("task already exist: " + arg)
		}

		totree(slices.Collect(pathIter(filepath.Dir(arg))))

		dir, file := filepath.Split(arg)
		day := time.Now().Format(time.DateOnly)
		file = fmt.Sprintf("%s %s", day, file)

		f, err := os.Create(filepath.Join(dir, file))
		if err != nil {
			return err
		}

		_ = f.Close()
	}

	return nil
}
