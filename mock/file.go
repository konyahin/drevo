package mock

import (
	"io/fs"
	"log"
	"os"
	"slices"
)

type Stats struct {
	Info os.FileInfo
	Err  error
}

type callType int

const (
	createCall = iota
	statCall
)

type FakeFileManager struct {
	createMocks map[string]error
	statMocks   map[string]*Stats
	calledMocks map[string][]callType
}

func NewFakeFileManager() *FakeFileManager {
	return &FakeFileManager{
		make(map[string]error),
		make(map[string]*Stats),
		make(map[string][]callType),
	}
}

func (ffm *FakeFileManager) addCall(path string, call callType) {
	calls := ffm.calledMocks[path]
	if calls == nil {
		calls = make([]callType, 0)
	}
	calls = append(calls, call)
	ffm.calledMocks[path] = calls
}

func (ffm *FakeFileManager) IsDirCreated(path string) bool {
	return ffm.isCalled(path, createCall)
}

func (ffm *FakeFileManager) isCalled(path string, call callType) bool {
	calls := ffm.calledMocks[path]
	if calls == nil {
		return false
	}
	return slices.Contains(calls, call)
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

func (ffm *FakeFileManager) CreatedError(path string, err error) {
	ffm.createMocks[path] = err
}

func (ffm *FakeFileManager) StatError(path string, err error) {
	ffm.statMocks[path] = &Stats{
		Info: FakeFileInfo{},
		Err:  err,
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
	ffm.addCall(name, createCall)
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
	ffm.addCall(path, statCall)
	return stats.Info, stats.Err
}

func (_ FakeFileManager) Rename(old, new string) error {
	return fs.ErrNotExist
}
