package main

import (
	"errors"
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

func TestGenerRandomInt(t *testing.T) {
	filename := "input_test.txt"
	if _, err := os.Stat(filename); !errors.Is(err, os.ErrNotExist) {
		// file exists
		t.Errorf("File %s should not exist at this moment!", filename)
	}

	geneRandomInt(filename)
	fmt.Println(filename, " created for testing!")
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		// file does not exist
		t.Errorf("File %s should exist after calling geneRandomInt()!", filename)
	}

	err := os.Remove(filename) // rm before leave the test
	checkErr(err)
	if _, err := os.Stat(filename); !errors.Is(err, os.ErrNotExist) {
		// file exists
		t.Errorf("File %s should not exist after Remove()!", filename)
	}
	fmt.Println(filename, " removed successfully after testing!")

}
