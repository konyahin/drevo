package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func init() {
	addCommand((Command{
		"find",
		findHelp,
		func(day string, args []string) error {
			if len(args) > 2 {
				args = args[2:]
			} else {
				args = nil
			}

			tasks, err := find(args)
			if err == nil {
				for _, task := range tasks {
					fmt.Println(task)
				}
			}
			return err
		},
	}))
}

func findHelp() {
	fmt.Println("find text ... - finds tasks matching ALL specified patterns (logical AND)")
}

func find(args []string) ([]string, error) {
	var tasks []string
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		switch {
		case path == ".":
			return nil
		case d.IsDir() && (strings.Contains(path, "/.") || strings.HasPrefix(path, ".")):
			return fs.SkipDir
		case !d.IsDir():
			return nil
		}

		for _, pattern := range args {
			if !strings.Contains(path, pattern) {
				return nil
			}
		}

		tasks = append(tasks, path)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
