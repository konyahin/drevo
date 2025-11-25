package main

type Command struct {
	Name string
	Help func()
	Call func(day string, args []string) error
}
