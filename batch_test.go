package main

import (
	"os"
	"strings"
	"testing"
)

func TestBatch(t *testing.T) {
	dir := t.TempDir()
	tasks := []string{"task1", "task2"}
	tasksBytes := []byte(strings.Join(tasks, "\n"))

	if err := os.Chdir(dir); err != nil {
		t.Fatal("Can't change dir:", err)
	}

	inputPath := dir + "/input.txt"
	if err := os.WriteFile(inputPath, tasksBytes, 0666); err != nil {
		t.Fatal("Can't write to file:", err)
	}

	if err := batch("", inputPath); err != nil {
		t.Fatal(err)
	}

	if !isDirCreated(dir + "/task1") {
		t.Fatal("The dir should be created:", dir+"/task1")
	}

	if !isDirCreated(dir + "/task2") {
		t.Fatal("The dir should be created:", dir+"/task2")
	}
}
