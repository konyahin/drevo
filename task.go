package main

import (
	"strings"
)

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

type wordBuilder struct {
	strings.Builder
}

func New(path string) *Task {
	task := new(Task)
	task.Path = path
	task.Tags = make(map[string]string)

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

	return task
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

func (buf *wordBuilder) writeWord(word string) {
	if buf.Len() != 0 {
		buf.WriteRune(' ')
	}
	buf.WriteString(word)
}
