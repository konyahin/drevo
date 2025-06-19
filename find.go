package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func init() {
	helps = append(helps, findHelp)
}

func findHelp() {
	fmt.Println("find text ... - finds files matching ALL specified patterns (logical AND)")
}

func find(args []string) ([]string, error) {
	var tasks []string
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		switch {
		case path == ".":
			return nil
		case d.IsDir() && (strings.Contains(path, "/.") || strings.HasPrefix(path, ".")):
			return fs.SkipDir
		case !d.IsDir() || strings.Contains(path, "/.") || strings.HasPrefix(path, "."):
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
