package main

import (
	"errors"
	"os"
	"testing"
)

func TestCompleteTask(t *testing.T) {
	dir := t.TempDir()
	task := dir + "/test task"
	completedTask := dir + "/x 2025-06-20 test task"

	if err := os.Mkdir(task, 0750); err != nil {
		t.Fatal("Can't create a dir:", err)
	}

	if err := complete("2025-06-20", task); err != nil {
		t.Fatal(err)
	}

	if isDirCreated(task) {
		t.Fatal("The dir should be deleted:", task)
	}

	if !isDirCreated(completedTask) {
		t.Fatal("The dir should be created:", completedTask)
	}
}

func TestCompleteTaskWithDate(t *testing.T) {
	dir := t.TempDir()
	task := dir + "/2025-06-15 test task"
	completedTask := dir + "/x 2025-06-27 2025-06-15 test task"

	if err := os.Mkdir(task, 0750); err != nil {
		t.Fatal("Can't create a dir:", err)
	}

	if err := complete("2025-06-27", task); err != nil {
		t.Fatal(err)
	}

	if isDirCreated(task) {
		t.Fatal("The dir should be deleted:", task)
	}

	if !isDirCreated(completedTask) {
		t.Fatal("The dir should be created:", completedTask)
	}
}

func TestCompleteTaskWithoutTask(t *testing.T) {
	dir := t.TempDir()
	task := dir + "/2025-06-27 test task"

	err := complete("", task)
	if err == nil {
		t.Fatal("Should be an error")
	}

	if !errors.Is(err, ErrTaskDoesntExist) {
		t.Fatal("Wrong error:", err)
	}
}

func TestCompleteTaskAlreadyComplete(t *testing.T) {
	dir := t.TempDir()
	task := dir + "/x test task"

	if err := os.Mkdir(task, 0750); err != nil {
		t.Fatal("Can't create a dir:", err)
	}

	if err := complete("2025-06-20", task); err != nil {
		t.Fatal(err)
	}

	if !isDirCreated(task) {
		t.Fatal("The dir should stay:", task)
	}
}
