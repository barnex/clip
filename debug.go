package main

// This file implements log and debug functions

import (
	"os"
	"fmt"
)

// Print debug message to stderr.
func Debug(msg ...interface{}) {
	fmt.Fprintln(os.Stderr, msg...)
}

// Print error message and exit.
func Err(msg ...interface{}) {
	fmt.Fprintln(os.Stderr, msg...)
	os.Exit(3)
}

// Check for error.
// If err != nil, print it and exit.
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
