package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	// Calling go in a separate process to build the application
	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build toold %s: %s", binName, err)
		os.Exit(1)
	}
	fmt.Println("Running tests...")
	result := m.Run() // Runs all of the tests in the package
	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(fileName) // Removes the file on the process's behalf
	os.Exit(result)
}
