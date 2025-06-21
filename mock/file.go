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
	createMocks map[string]error
	statMocks   map[string]*Stats
}

func NewFakeFileManager() *FakeFileManager {
	return &FakeFileManager{
		make(map[string]error),
		make(map[string]*Stats),
	}
}

func (ffm *FakeFileManager) DirExist(path string) {
	ffm.statMocks[path] = &Stats{
		Info: FakeFileInfo{Dir: true},
		Err:  nil,
	}
}

func (ffm *FakeFileManager) FileExist(path string) {
	ffm.statMocks[path] = &Stats{
		Info: FakeFileInfo{Dir: false},
		Err:  nil,
	}
}

func (ffm *FakeFileManager) DoesntExist(path string) {
	ffm.statMocks[path] = &Stats{
		Info: FakeFileInfo{},
		Err:  fs.ErrNotExist,
	}
}

func (ffm *FakeFileManager) DirCreated(path string) {
	ffm.createMocks[path] = nil
}

func (ffm FakeFileManager) CreateTask(name string) error {
	err, ok := ffm.createMocks[name]
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
	stats, ok := ffm.statMocks[path]
	if !ok {
		log.Fatal("Missing mock for Stat: ", path)
	}
	return stats.Info, stats.Err
}

func (_ FakeFileManager) Rename(old, new string) error {
	return fs.ErrNotExist
}
