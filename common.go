package main

import (
	"errors"
	"io/fs"
	"iter"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type TaskState int

const (
	DoesntExist = iota
	Leaf
	Tree
)

func taskState(path string) (TaskState, error) {
	for part := range pathIter(filepath.Dir(path)) {
		state, err := fileState(part)
		if err != nil {
			return 0, err
		}

		if state != Tree {
			return DoesntExist, err
		}
	}

	return fileState(path)
}

func fileState(path string) (TaskState, error) {
	fileInfo, err := os.Stat(path)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return DoesntExist, nil
	case err != nil:
		return 0, err
	case fileInfo.IsDir():
		return Tree, nil
	default:
		return Leaf, nil
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
