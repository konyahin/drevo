package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func init() {
	helps = append(helps, uncompleteHelp)
}

func uncompleteHelp() {
	fmt.Println("uncomplete file - mark task as uncompleted")
}

func uncomplete(arg string) error {
	if arg == "" {
		uncompleteHelp()
		return nil
	}

	err := shouldBeATask(arg)
	if err != nil {
		return err
	}

	dir, file := filepath.Split(arg)
	if !strings.HasPrefix(file, "x ") {
		return nil
	}

	file = file[2:]
	parts := strings.SplitN(file, " ", 2)

	if isDate(parts[0]) {
		file = parts[1]
	}
	return fm.Rename(arg, filepath.Join(dir, file))
}
