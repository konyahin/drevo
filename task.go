package main

import (
	"os"
	"path/filepath"
	"strings"
)

type wordBuilder struct {
	strings.Builder
}

func (buf *wordBuilder) writeWord(word string) {
	if buf.Len() != 0 {
		buf.WriteRune(' ')
	}
	buf.WriteString(word)
}

type Task struct {
	Path       string
	Text       string
	Done       bool
	Completion string
	Creation   string
	Projects   []string
	Contexts   []string
	Tags       map[string]string
}

func New(path string) *Task {
	task := new(Task)
	task.Path = path
	task.Tags = make(map[string]string)

	for folder := range strings.SplitSeq(path, "/") {
		task.enrichFromPath(folder)
	}

	return task
}

func (task *Task) enrichFromPath(path string) {
	words := strings.Split(path, " ")

	if words[0] == "x" {
		task.Done = true
		words = words[1:]
	}

	if isDate(words[0]) {
		if task.Done {
			task.Completion = words[0]
		} else {
			task.Creation = words[0]
		}
		words = words[1:]
	}

	if task.Done && isDate(words[0]) {
		task.Creation = words[0]
		words = words[1:]
	}

	var buf wordBuilder
	for _, word := range words {
		if word == "" {
			continue
		}

		buf.writeWord(word)

		switch {
		case word[0] == '+':
			task.Projects = append(task.Projects, word)
		case word[0] == '@':
			task.Contexts = append(task.Contexts, word)
		case strings.Contains(word, ":"):
			key, value, _ := strings.Cut(word, ":")
			task.Tags[key] = value
		}
	}
	task.Text = buf.String()
}

func (task Task) String() string {
	var buf wordBuilder

	if task.Done {
		buf.writeWord("x")
	}

	if task.Done && task.Completion != "" {
		buf.writeWord(task.Completion)
	}

	if task.Creation != "" {
		buf.writeWord(task.Creation)
	}

	buf.writeWord(task.Text)

	return buf.String()
}

func (task *Task) Complete(day string) error {
	if task.Done {
		return nil
	}

	task.Done = true
	task.Completion = day
	return task.update()
}

func (task *Task) update() error {
	oldpath := task.Path
	parent, _ := filepath.Split(task.Path)
	name := task.String()

	task.Path = filepath.Join(parent, name)
	if oldpath == task.Path {
		return nil
	}

	return os.Rename(oldpath, task.Path)
}
