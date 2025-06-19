package main

import "os"

type FileManager interface {
	CreateTask(name string) error
	CreateTempFile() (*os.File, error)
	RemoveFile(path string)
	Stat(path string) (os.FileInfo, error)
	Rename(old, new string) error
}

type realFileManager struct{}

func (_ realFileManager) CreateTask(name string) error {
	return os.MkdirAll(name, 0755)
}

func (_ realFileManager) CreateTempFile() (*os.File, error) {
	return os.CreateTemp("", "deltatask")
}

func (_ realFileManager) RemoveFile(path string) {
	_ = os.Remove(path)
}

func (_ realFileManager) Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

func (_ realFileManager) Rename(old, new string) error {
	return os.Rename(old, new)
}
