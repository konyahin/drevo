package main

import "os"

type FileManager interface {
	CreateTask(name string) error
}

type realFileManager struct{}

func (_ realFileManager) CreateTask(name string) error {
	return os.MkdirAll(name, 0755)
}
