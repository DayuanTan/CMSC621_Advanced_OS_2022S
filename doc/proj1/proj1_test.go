package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

const fnameTest = "input_data_file_test.txt"

func TestCommandLineArgument4MainPanic(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }() // os.Args is a "global variable", so keep the state from before the test, and restore it after.

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The main() did not panic for wrong command line arguments.")
		}
	}()

	os.Args = []string{"cmd", "3", fnameTest, "this argument causes panic"}
	main()
}

func TestGenerRandomInt(t *testing.T) {
	filename := "input_data_file_test.txt"
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
