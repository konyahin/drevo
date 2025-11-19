package main

import (
	"errors"
	"fmt"
	"io/fs"
	"iter"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type TaskState int

const (
	DoesntExist = iota
	NotATask
	Ok
)

var (
	ErrTaskDoesntExist error = errors.New("task doesn't exist")
	ErrNotATask        error = errors.New("it's not a task, it is a file")
	ErrFileInPath      error = errors.New("path contain file, not a folder")
)

func taskState(path string) (TaskState, error) {
	for part := range pathIter(path) {
		state, err := fileState(part)
		if err != nil {
			return -1, err
		}

		if state != Ok {
			return state, nil
		}
	}

	return Ok, nil
}

func fileState(path string) (TaskState, error) {
	fileInfo, err := os.Stat(path)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return DoesntExist, nil
	case err != nil:
		return -1, err
	case fileInfo.IsDir():
		return Ok, nil
	default:
		return NotATask, nil
	}
}

func pathIter(path string) iter.Seq[string] {
	return func(yield func(string) bool) {
		parts := strings.Split(path, string(filepath.Separator))
		partPath := ""
		if strings.HasPrefix(path, "/") {
			partPath = "/"
		}
		for _, part := range parts {
			partPath = filepath.Join(partPath, part)
			if !yield(partPath) {
				return
			}
		}
	}
}

func isDate(s string) bool {
	_, err := time.Parse(time.DateOnly, s)
	return err == nil
}

func shouldBeATask(path string) error {
	state, err := taskState(path)
	if err != nil {
		return err
	}

	if state == DoesntExist {
		return fmt.Errorf("%w: %s", ErrTaskDoesntExist, path)
	}

	if state == NotATask {
		return fmt.Errorf("%w: %s", ErrNotATask, path)
	}

	return nil
}

func editTempFile() (string, error) {
	tmpfile, err := os.CreateTemp("", "drevo")
	if err != nil {
		return "", err
	}
	_ = tmpfile.Close()

	fileName := tmpfile.Name()
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	cmd := exec.Command(editor, fileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		return "", err
	}

	return fileName, nil
}
