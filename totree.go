package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func init() {
	helps = append(helps, totreeHelp)
}

func totreeHelp() {
	fmt.Println("totree file ... - convert a leafe(s) into a task folder(s)")
}

func totree(args []string) error {
	if len(args) < 1 {
		totreeHelp()
		return nil
	}

	for _, arg := range args {
		state, err := taskState(arg)
		if err != nil {
			return err
		}

		switch state {
		case DoesntExist:
			err = os.MkdirAll(arg, 0755)
		case Leaf:
			err = leafToTree(arg)
		case Tree:
			// already tree, do nothing
		default:
			panic("unknown TaskState in totree!")
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func leafToTree(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	tempDir := path + ".tmp"
	defer os.RemoveAll(tempDir)

	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return err
	}

	taskFile := filepath.Join(tempDir, ".task")
	if err := os.WriteFile(taskFile, content, 0644); err != nil {
		return err
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	return os.Rename(tempDir, path)
}
