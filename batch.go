package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func init() {
	helps = append(helps, batchHelp)
}

func batchHelp() {
	fmt.Println("batch [root] - create tasks from temp file in root dir, one task per line")
}

func batch(day string, root string) error {
	tmpfile, err := os.CreateTemp("", "deltatask")
	if err != nil {
		return err
	}
	defer func() { _ = os.Remove(tmpfile.Name()) }()

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	cmd := exec.Command(editor, tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		return err
	}

	content, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		return err
	}

	tasks := strings.Split(string(content), "\n")
	for i, task := range tasks {
		if task != "" {
			tasks[i] = filepath.Join(root, task)
		}
	}

	return create(day, tasks)
}
