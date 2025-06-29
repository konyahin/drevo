package main

import (
	"os"
	"testing"
)

func TestFind(t *testing.T) {
	dir := t.TempDir()
	task := dir + "/some +active task"

	if err := os.Mkdir(task, 0750); err != nil {
		t.Fatal("Can't create a dir:", err)
	}

	if err := os.Chdir(dir); err != nil {
		t.Fatal("Can't change working dir:", err)
	}

	tasks, err := find([]string{"+active"})
	if err != nil {
		t.Fatal(err)
	}

	if len(tasks) != 1 {
		t.Fatal("Wrong tasks found:", tasks)
	}
}

