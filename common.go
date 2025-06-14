package main

import (
	"os"
)

type TaskState int

const (
	DoesntExist = iota
	Leaf
	Tree
)

func taskState(path string) (TaskState, error) {
	fileInfo, err := os.Stat(path)
	switch {
	case os.IsNotExist(err):
		return DoesntExist, nil
	case err != nil:
		return 0, err
	case fileInfo.IsDir():
		return Tree, nil
	default:
		return Leaf, nil
	}
}
