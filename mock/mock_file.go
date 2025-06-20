package mock

import (
	"io/fs"
	"log"
	"os"
)

type Stats struct {
	Info os.FileInfo
	Err  error
}

type FakeFileManager struct {
	CreateMocks map[string]error
	StatMocks   map[string]*Stats
}

func NewFakeFileManager() *FakeFileManager {
	return &FakeFileManager{
		make(map[string]error),
		make(map[string]*Stats),
	}
}

func (ffm FakeFileManager) CreateTask(name string) error {
	err, ok := ffm.CreateMocks[name]
	if !ok {
		log.Fatal("Missing mock for CreateTask: ", name)
	}
	return err
}

func (_ FakeFileManager) CreateTempFile() (*os.File, error) {
	return nil, fs.ErrNotExist
}

func (_ FakeFileManager) RemoveFile(path string) {
}

func (ffm FakeFileManager) Stat(path string) (os.FileInfo, error) {
	stats, ok := ffm.StatMocks[path]
	if !ok {
		log.Fatal("Missing mock for Stat: ", path)
	}
	return stats.Info, stats.Err
}

func (_ FakeFileManager) Rename(old, new string) error {
	return fs.ErrNotExist
}
