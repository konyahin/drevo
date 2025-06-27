package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var day string = "2025-06-20"

func isDirCreated(path string) bool {
	if stats, err := os.Stat(path); err == nil {
		return stats.IsDir()
	}

	return false
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
	if err := create(day, []string{""}); err != nil {
		t.Error(err)
	}
}

func TestCreateTask(t *testing.T) {
	fileName, task := generateTask(t)

	if err := create(day, []string{task}); err != nil {
		t.Error(err)
	}

	if !isDirCreated(fileName) {
		t.Error("Dir is not created:", fileName)
	}
}

func TestCreateTaskFail(t *testing.T) {
	dir := t.TempDir()
	_ = os.Chmod(dir, 0600)
	fileName, task := generateTaskForDir(t, dir)

	if err := create(day, []string{task}); err == nil {
		t.Error("Should be an error")
	}

	if isDirCreated(fileName) {
		t.Error("Dir is created:", fileName)
	}
}

func TestCreateTaskAlreadyExist(t *testing.T) {
	dir := t.TempDir()
	fileName, task := generateTaskForDir(t, dir)
	if err := os.Mkdir(fileName, 750); err != nil {
		t.Error("Can't create a dir:", err)
	}

	if err := create(day, []string{task}); err != nil {
		t.Error(err)
	}
}

func TestCreateTaskInFolder(t *testing.T) {
	dir := t.TempDir() + "/another_folder"
	fileName, task := generateTaskForDir(t, dir)

	if err := create(day, []string{task}); err != nil {
		t.Error(err)
	}

	if !isDirCreated(fileName) {
		t.Error("Dir is not created:", fileName)
	}
}

func TestCreateTaskInFile(t *testing.T) {
	baseFile := t.TempDir() + "/some_file"
	file, err := os.Create(baseFile)
	if err != nil {
		t.Error("Can't create a file:", err)
	}
	_ = file.Close()

	fileName, task := generateTaskForDir(t, baseFile)

	err = create(day, []string{task})
	if !strings.HasPrefix(err.Error(), "taks path contain file (not a folder)") {
		t.Error(err)
	}

	if isDirCreated(fileName) {
		t.Error("Dir is created:", fileName)
	}
}
