package main

import (
	"errors"
	"io/fs"
	"iter"
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

func taskState(path string) (TaskState, error) {
	for part := range pathIter(path) {
		state, err := fileState(part)
		if err != nil {
			return 0, err
		}

		if state != Ok {
			return state, nil
		}
	}

	return Ok, nil
}

func fileState(path string) (TaskState, error) {
	fileInfo, err := fm.Stat(path)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return DoesntExist, nil
	case err != nil:
		return 0, err
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
		return errors.New("task doesn't exist: " + path)
	}

	if state == NotATask {
		return errors.New("it's not a task, this is a file: " + path)
	}

	return nil
}
