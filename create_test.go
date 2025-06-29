package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

var day string = "2025-06-20"

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
		t.Fatal(err)
	}
}

func TestCreateTask(t *testing.T) {
	fileName, task := generateTask(t)

	if err := create(day, []string{task}); err != nil {
		t.Fatal(err)
	}

	if !isDirCreated(fileName) {
		t.Fatal("Dir is not created:", fileName)
	}
}

func TestCreateTaskFail(t *testing.T) {
	dir := t.TempDir()
	_ = os.Chmod(dir, 0600)
	fileName, task := generateTaskForDir(t, dir)

	if err := create(day, []string{task}); err == nil {
		t.Fatal("Should be an error")
	}

	if isDirCreated(fileName) {
		t.Fatal("The dir is created:", fileName)
	}
}

func TestCreateTaskAlreadyExist(t *testing.T) {
	dir := t.TempDir()
	fileName, task := generateTaskForDir(t, dir)
	if err := os.Mkdir(fileName, 0750); err != nil {
		t.Fatal("Can't create a dir:", err)
	}

	if err := create(day, []string{task}); err != nil {
		t.Fatal(err)
	}
}

func TestCreateTaskInFolder(t *testing.T) {
	dir := t.TempDir() + "/another_folder"
	fileName, task := generateTaskForDir(t, dir)

	if err := create(day, []string{task}); err != nil {
		t.Fatal(err)
	}

	if !isDirCreated(fileName) {
		t.Fatal("The dir is not created:", fileName)
	}
}

func TestCreateTaskInFile(t *testing.T) {
	baseFile := t.TempDir() + "/some_file"
	file, err := os.Create(baseFile)
	if err != nil {
		t.Fatal("Can't create a file:", err)
	}
	_ = file.Close()

	fileName, task := generateTaskForDir(t, baseFile)

	err = create(day, []string{task})
	if !errors.Is(err, ErrFileInPath) {
		t.Fatal(err)
	}

	if isDirCreated(fileName) {
		t.Fatal("The dir is created:", fileName)
	}
}

func TestCreateFewTask(t *testing.T) {
	fileName, task := generateTask(t)
	fileName2, task2 := generateTask(t)

	if err := create(day, []string{task, task2}); err != nil {
		t.Fatal(err)
	}

	if !isDirCreated(fileName) {
		t.Fatal("Dir is not created:", fileName)
	}

	if !isDirCreated(fileName2) {
		t.Fatal("Dir is not created:", fileName2)
	}
}

func TestCreateFewTaskAlreadyExist(t *testing.T) {
	dir := t.TempDir()
	fileName, task := generateTaskForDir(t, dir)
	fileName2, task2 := generateTask(t)
	if err := os.Mkdir(fileName, 0750); err != nil {
		t.Fatal("Can't create a dir:", err)
	}

	if err := create(day, []string{task, task2}); err != nil {
		t.Fatal(err)
	}

	if !isDirCreated(fileName2) {
		t.Fatal("Dir is not created:", fileName2)
	}
}

func TestCreateTaskWithoutDate(t *testing.T) {
	dir := t.TempDir()
	fileName := dir + "/" + t.Name()

	if err := create("", []string{fileName}); err != nil {
		t.Fatal(err)
	}

	if !isDirCreated(fileName) {
		t.Fatal("Dir is not created:", fileName)
	}
}
