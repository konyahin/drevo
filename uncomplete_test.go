package main

import (
	"os"
	"strings"
	"testing"
)

func TestUncompleteTask(t *testing.T) {
	dir := t.TempDir()
	task := dir + "/x test task"
	uncompletedTask := dir + "/test task"

	if err := os.Mkdir(task, 0750); err != nil {
		t.Fatal("Can't create a dir:", err)
	}

	if err := uncomplete(task); err != nil {
		t.Fatal(err)
	}

	if isDirCreated(task) {
		t.Fatal("The dir should be deleted:", task)
	}

	if !isDirCreated(uncompletedTask) {
		t.Fatal("The dir should be created:", uncompletedTask)
	}
}

func TestUncompleteTaskWithDate(t *testing.T) {
	dir := t.TempDir()
	task := dir + "/x 2025-06-27 test task"
	uncompletedTask := dir + "/test task"

	if err := os.Mkdir(task, 0750); err != nil {
		t.Fatal("Can't create a dir:", err)
	}

	if err := uncomplete(task); err != nil {
		t.Fatal(err)
	}

	if isDirCreated(task) {
		t.Fatal("The dir should be deleted:", task)
	}

	if !isDirCreated(uncompletedTask) {
		t.Fatal("The dir should be created:", uncompletedTask)
	}
}

func TestUncompleteTaskWithoutTask(t* testing.T) {
	dir := t.TempDir()
	task := dir + "/x test task"

	err := uncomplete(task)
	if err == nil {
		t.Fatal("Should be an error")
	}

	if !strings.HasPrefix(err.Error(), "task doesn't exist") {
		t.Fatal("Wrong error:", err)
	}
}

func TestUncompleteTaskAlreadyUncomplete(t *testing.T) {
	dir := t.TempDir()
	task := dir + "/test task"

	if err := os.Mkdir(task, 0750); err != nil {
		t.Fatal("Can't create a dir:", err)
	}

	if err := uncomplete(task); err != nil {
		t.Fatal(err)
	}

	if !isDirCreated(task) {
		t.Fatal("The dir should stay:", task)
	}
}
