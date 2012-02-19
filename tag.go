package main

// This file implements fuzzy tagging

import (
	"strings"
)


func Fuzzy(file string) string {
	file = strings.ToLower(file)
	return file
}
