package main

import "os"

func isDirCreated(path string) bool {
	if stats, err := os.Stat(path); err == nil {
		return stats.IsDir()
	}

	return false
}

