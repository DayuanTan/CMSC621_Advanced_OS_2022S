package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCommandLineArgument(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }() // os.Args is a "global variable", so keep the state from before the test, and restore it after.

	fmt.Println("Expected: \nTwo arguments are:  3   fnameTestname")
	fmt.Println("Got:")
	os.Args = []string{"cmd", "3", "fnameTestname"}
	main()
}
