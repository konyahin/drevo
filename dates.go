package main

import "time"

type Dates interface {
	CurrentDate() string
}

type realDates struct{}

func (_ realDates) CurrentDate() string {
	return time.Now().Format(time.DateOnly)
}
