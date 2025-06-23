package main

import (
	"fmt"
	"konyahin.xyz/deltatree/mock"
	"os"
	"strings"
	"testing"
)

var day string = "2025-06-20"
var ffm *mock.FakeFileManager = mock.NewFakeFileManager()

func init() {
	fm = ffm
}

func isDirCreated(path string) bool {
	stats, err := os.Stat(path)
	if err != nil {
		return false
	}

	return stats.IsDir()
}

func generateTaskForDir(t *testing.T, dir string) (fullPath, task string) {
	task = fmt.Sprintf("%s/%s", dir, t.Name())
	fullPath = fmt.Sprintf("%s/%s %s", dir, day, t.Name())
	return
}

func generateTask(t *testing.T) (fullPath, task string) {
	return generateTaskForDir(t, t.TempDir())
}

func TestCreateEmptyArgs(t *testing.T) {
	err := create(day, []string{""})
	if err != nil {
		t.Error(err)
	}
}

func TestCreateTask(t *testing.T) {
	fm = realFileManager{}

	fileName, task := generateTask(t)

	err := create(day, []string{task})
	if err != nil {
		t.Error(err)
	}

	if !isDirCreated(fileName) {
		t.Error("Dir is not created:", fileName)
	}
}

func TestCreateTaskFail(t *testing.T) {
	fm = realFileManager{}

	dir := t.TempDir()
	_ = os.Chmod(dir, 0600)
	fileName, task := generateTaskForDir(t, dir)

	err := create(day, []string{task})
	if err == nil {
		t.Error("Should be an error")
	}

	if isDirCreated(fileName) {
		t.Error("Dir is created:", fileName)
	}
}

func TestCreateTaskAlreadyExist(t *testing.T) {
	fm = ffm

	task := t.Name()
	fileName := "2025-06-20 " + task

	ffm.DirExist(fileName)

	err := create(day, []string{task})
	if err != nil {
		t.Error(err)
	}
}

func TestCreateTaskInFolder(t *testing.T) {
	fm = ffm

	folder := "folder" + t.Name()
	task := folder + "/" + t.Name()
	fileName := "2025-06-20 " + t.Name()
	fullPath := folder + "/" + fileName

	ffm.DirExist(folder)
	ffm.DoesntExist(fullPath)
	ffm.DirCreated(fullPath)

	err := create(day, []string{task})
	if err != nil {
		t.Error(err)
	}

	if !ffm.IsDirCreated(fullPath) {
		t.Error("Dir is not created:", fullPath)
	}
}

func TestCreateTaskInFile(t *testing.T) {
	fm = ffm

	folder := "folder" + t.Name()
	task := folder + "/" + t.Name()
	fileName := "2025-06-20 " + t.Name()
	fullPath := folder + "/" + fileName

	ffm.FileExist(folder)
	ffm.DoesntExist(fullPath)

	err := create(day, []string{task})
	if !strings.HasPrefix(err.Error(), "taks path contain file (not a folder)") {
		t.Error(err)
	}

	if ffm.IsDirCreated(fullPath) {
		t.Error("Dir is created:", fullPath)
	}
}
