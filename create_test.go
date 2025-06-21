package main

import (
	"konyahin.xyz/deltatree/mock"
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
}

func TestCreateTaskAlreadyExist(t *testing.T) {
	task := t.Name()
	fileName := "2025-06-20 " + task

	ffm.DirExist(fileName)

	err := create([]string{task})
	if err.Error() != "task already exist: TestCreateTaskAlreadyExist" {
		t.Error(err)
	}
}
