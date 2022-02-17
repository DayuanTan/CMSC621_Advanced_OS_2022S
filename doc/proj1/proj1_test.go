package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

const fnameTest = "input_data_file_test.txt"
const fnameReadUntilBlankByteTest = "test_data_for_readUntilBlankByte.test"
const fnameSumupTest = "test_data_for_sumuptest.test"

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
	if _, err := os.Stat(fnameTest); !errors.Is(err, os.ErrNotExist) {
		// file exists
		t.Errorf("File %s should not exist at this moment!", fnameTest)
	}

	geneRandomInt(fnameTest)
	fmt.Println(fnameTest, " created for testing geneRandomInt()!")
	if _, err := os.Stat(fnameTest); errors.Is(err, os.ErrNotExist) {
		// file does not exist
		t.Errorf("File %s should exist after calling geneRandomInt()!", fnameTest)
	}

	err := os.Remove(fnameTest) // rm before leaving the test
	checkErr(err)
	if _, err := os.Stat(fnameTest); !errors.Is(err, os.ErrNotExist) {
		// file exists
		t.Errorf("File %s should not exist after Remove()!", fnameTest)
	}
	fmt.Println(fnameTest, " removed successfully after testing!")
}

func TestReadUntilBlankByte(t *testing.T) {
	testdata := []byte("1 2 33 444 555\n")
	err := os.WriteFile(fnameReadUntilBlankByteTest, testdata, 0644) //0644: -rw-r--r--
	checkErr(err)
	fmt.Println(fnameReadUntilBlankByteTest, " created for testing readUntilBlankByteTest()!")

	actual := readUntilBlankByte(fnameReadUntilBlankByteTest, 0, 14)
	expected := int64(1)
	if actual != expected {
		t.Errorf("TestReadUntilBlankByte() == %q, want %q", actual, expected)
	}
	fmt.Println("0")

	actual = readUntilBlankByte(fnameReadUntilBlankByteTest, 1, 14)
	expected = int64(1)
	if actual != expected {
		t.Errorf("TestReadUntilBlankByte() == %q, want %q", actual, expected)
	}
	fmt.Println("1")

	actual = readUntilBlankByte(fnameReadUntilBlankByteTest, 2, 14)
	expected = int64(3)
	if actual != expected {
		t.Errorf("TestReadUntilBlankByte() == %q, want %q", actual, expected)
	}
	fmt.Println("2")

	actual = readUntilBlankByte(fnameReadUntilBlankByteTest, 3, 14)
	expected = int64(3)
	if actual != expected {
		t.Errorf("TestReadUntilBlankByte() == %q, want %q", actual, expected)
	}
	fmt.Println("3")

	actual = readUntilBlankByte(fnameReadUntilBlankByteTest, 4, 14)
	expected = int64(6)
	if actual != expected {
		t.Errorf("TestReadUntilBlankByte() == %q, want %q", actual, expected)
	}
	fmt.Println("4")

	actual = readUntilBlankByte(fnameReadUntilBlankByteTest, 13, 14)
	expected = int64(13)
	if actual != expected {
		t.Errorf("TestReadUntilBlankByte() == %q, want %q", actual, expected)
	}
	fmt.Println("14")

}

func TestSumup(t *testing.T) {
	testdata := []byte("1 2 33 444 555")
	err := os.WriteFile(fnameSumupTest, testdata, 0644) //0644: -rw-r--r--
	checkErr(err)
	fmt.Println(fnameSumupTest, " created for testing sumup()!")

	sumup(fnameSumupTest, 0, 7)

	err = os.Remove(fnameSumupTest) // rm before leaving the test
	checkErr(err)
	if _, err := os.Stat(fnameSumupTest); !errors.Is(err, os.ErrNotExist) {
		// file exists
		t.Errorf("File %s should not exist after Remove()!", fnameSumupTest)
	}
	fmt.Println(fnameSumupTest, " removed successfully after testing!")
}
