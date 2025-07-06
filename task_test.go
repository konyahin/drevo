package main

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	path := "x 2025-07-06 2025-07-05 new task for +project at @home due:yesterday"
	expect := Task{
		Path:       path,
		Text:       "new task for +project at @home due:yesterday",
		Done:       true,
		Completion: "2025-07-06",
		Creation:   "2025-07-05",
		Projects:   []string{"+project"},
		Contexts:   []string{"@home"},
		Tags: map[string]string{
			"due": "yesterday",
		},
	}
	task := *New(path)

	if !reflect.DeepEqual(expect, task) {
		t.Errorf("\nExpect:\n\t%#v\nGot:\n\t%#v", expect, task)
	}
}

func TestNewNoStartDate(t *testing.T) {
	path := "x 2025-07-06 new task"
	expect := Task{
		Path:       path,
		Text:       "new task",
		Done:       true,
		Completion: "2025-07-06",
		Creation:   "",
		Tags:       map[string]string{},
	}
	task := *New(path)

	if !reflect.DeepEqual(expect, task) {
		t.Errorf("\nExpect:\n\t%#v\nGot:\n\t%#v", expect, task)
	}
}

func TestNewUncomplete(t *testing.T) {
	path := "2025-07-06 new task"
	expect := Task{
		Path:       path,
		Text:       "new task",
		Done:       false,
		Completion: "",
		Creation:   "2025-07-06",
		Tags:       map[string]string{},
	}
	task := *New(path)

	if !reflect.DeepEqual(expect, task) {
		t.Errorf("\nExpect:\n\t%#v\nGot:\n\t%#v", expect, task)
	}
}

func TestNewAdditionalSpaces(t *testing.T) {
	path := "2025-07-06    new    task   "
	expect := Task{
		Path:       path,
		Text:       "new task",
		Done:       false,
		Completion: "",
		Creation:   "2025-07-06",
		Tags:       map[string]string{},
	}
	task := *New(path)

	if !reflect.DeepEqual(expect, task) {
		t.Errorf("\nExpect:\n\t%#v\nGot:\n\t%#v", expect, task)
	}
}

func TestNewFewProjects(t *testing.T) {
	path := "2025-07-06 one task +two different +projects"
	expect := Task{
		Path:       path,
		Text:       "one task +two different +projects",
		Done:       false,
		Completion: "",
		Creation:   "2025-07-06",
		Projects:   []string{"+two", "+projects"},
		Tags:       map[string]string{},
	}
	task := *New(path)

	if !reflect.DeepEqual(expect, task) {
		t.Errorf("\nExpect:\n\t%#v\nGot:\n\t%#v", expect, task)
	}
}

func TestNewFewContexts(t *testing.T) {
	path := "2025-07-06 one task @two different @contexts"
	expect := Task{
		Path:       path,
		Text:       "one task @two different @contexts",
		Done:       false,
		Completion: "",
		Creation:   "2025-07-06",
		Contexts:   []string{"@two", "@contexts"},
		Tags:       map[string]string{},
	}
	task := *New(path)

	if !reflect.DeepEqual(expect, task) {
		t.Errorf("\nExpect:\n\t%#v\nGot:\n\t%#v", expect, task)
	}
}

func TestNewFewTags(t *testing.T) {
	path := "2025-07-06 task with due:tomorrow and deadline:2025-07-08"
	expect := Task{
		Path:       path,
		Text:       "task with due:tomorrow and deadline:2025-07-08",
		Done:       false,
		Completion: "",
		Creation:   "2025-07-06",
		Tags: map[string]string{
			"due":      "tomorrow",
			"deadline": "2025-07-08",
		},
	}
	task := *New(path)

	if !reflect.DeepEqual(expect, task) {
		t.Errorf("\nExpect:\n\t%#v\nGot:\n\t%#v", expect, task)
	}
}

func TestNewEmptyTask(t *testing.T) {
	path := ""
	expect := Task{
		Path:       path,
		Text:       "",
		Done:       false,
		Completion: "",
		Creation:   "",
		Tags:       map[string]string{},
	}
	task := *New(path)

	if !reflect.DeepEqual(expect, task) {
		t.Errorf("\nExpect:\n\t%#v\nGot:\n\t%#v", expect, task)
	}
}

func TestNewIncompleteTag(t *testing.T) {
	path := "some:"
	expect := Task{
		Path:       path,
		Text:       "some:",
		Done:       false,
		Completion: "",
		Creation:   "",
		Tags: map[string]string{
			"some": "",
		},
	}
	task := *New(path)

	if !reflect.DeepEqual(expect, task) {
		t.Errorf("\nExpect:\n\t%#v\nGot:\n\t%#v", expect, task)
	}
}

func TestTaskString(t *testing.T) {
	task := Task{
		Path:       "",
		Text:       "new task",
		Done:       true,
		Completion: "2025-07-06",
		Creation:   "2025-07-05",
	}
	path := task.String()

	expect := "x 2025-07-06 2025-07-05 new task"
	if path != expect {
		t.Errorf("\nExpect:\n\t%s\nGot:\n\t%s", expect, path)
	}
}

func TestTaskStringUncomplete(t *testing.T) {
	task := Task{
		Path:       "",
		Text:       "new task",
		Done:       false,
		Completion: "2025-07-06",
		Creation:   "2025-07-05",
	}
	path := task.String()

	expect := "2025-07-05 new task"
	if path != expect {
		t.Errorf("\nExpect:\n\t%s\nGot:\n\t%s", expect, path)
	}
}

func TestTaskStringNoDate(t *testing.T) {
	task := Task{
		Path: "",
		Text: "new task",
		Done: true,
	}
	path := task.String()

	expect := "x new task"
	if path != expect {
		t.Errorf("\nExpect:\n\t%s\nGot:\n\t%s", expect, path)
	}
}
