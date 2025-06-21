package main

import (
	"errors"
	"konyahin.xyz/deltatree/mock"
	"strings"
	"testing"
)

var ffm *mock.FakeFileManager = mock.NewFakeFileManager()
var ft mock.FakeTime = mock.FakeTime{}

func init() {
	fm = ffm
	dates = ft
}

func TestCreateEmptyArgs(t *testing.T) {
	err := create([]string{""})
	if err != nil {
		t.Error(err)
	}
}

func TestCreateTask(t *testing.T) {
	task := t.Name()
	fileName := "2025-06-20 " + task

	ffm.DoesntExist(fileName)
	ffm.DirCreated(fileName)

	err := create([]string{task})
	if err != nil {
		t.Error(err)
	}

	if !ffm.IsDirCreated(fileName) {
		t.Error("Dir is not created:", fileName)
	}
}

func TestCreateTaskFail(t *testing.T) {
	task := t.Name()
	fileName := "2025-06-20 " + task

	fsErr := errors.New("creation failed")
	ffm.DoesntExist(fileName)
	ffm.CreatedError(fileName, fsErr)

	err := create([]string{task})
	if err != fsErr {
		t.Error(err)
	}

	if !ffm.IsDirCreated(fileName) {
		t.Error("Dir is not created:", fileName)
	}
}

func TestCreateTaskAlreadyExist(t *testing.T) {
	task := t.Name()
	fileName := "2025-06-20 " + task

	ffm.DirExist(fileName)

	err := create([]string{task})
	if err != nil {
		t.Error(err)
	}
}

func TestCreateTaskStatError(t *testing.T) {
	task := t.Name()
	fileName := "2025-06-20 " + task

	fsErr := errors.New("some fs error")
	ffm.StatError(fileName, fsErr)

	err := create([]string{task})
	if err != fsErr {
		t.Error("Wrong error. Expect", fsErr, "got", err)
	}
}

func TestCreateTaskInFolder(t *testing.T) {
	folder := "folder" + t.Name()
	task := folder + "/" + t.Name()
	fileName := "2025-06-20 " + t.Name()
	fullPath := folder + "/" + fileName

	ffm.DirExist(folder)
	ffm.DoesntExist(fullPath)
	ffm.DirCreated(fullPath)

	err := create([]string{task})
	if err != nil {
		t.Error(err)
	}

	if !ffm.IsDirCreated(fullPath) {
		t.Error("Dir is not created:", fullPath)
	}
}

func TestCreateTaskInFile(t *testing.T) {
	folder := "folder" + t.Name()
	task := folder + "/" + t.Name()
	fileName := "2025-06-20 " + t.Name()
	fullPath := folder + "/" + fileName

	ffm.FileExist(folder)
	ffm.DoesntExist(fullPath)

	err := create([]string{task})
	if !strings.HasPrefix(err.Error(), "taks path contain file (not a folder)") {
		t.Error(err)
	}

	if ffm.IsDirCreated(fullPath) {
		t.Error("Dir is created:", fullPath)
	}
}
