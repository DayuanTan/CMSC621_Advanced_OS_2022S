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
	fmt.Println("\nTestCommandLineArgument4MainPanic:")

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }() // os.Args is a "global variable", so keep the state from before the test, and restore it after.

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The main() did not panic for wrong command line arguments.")
		}
	}()

	os.Args = []string{"cmd", "3", fnameTest, "this argument causes panic"}

	fmt.Println("Expected:	Please provide two arguments: M and fname!")
	fmt.Print("Got:		")
	main()

	fmt.Println("Done!")
}

func TestGenerRandomInt(t *testing.T) {
	fmt.Println("\nTestGenerRandomInt: ")

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

	fmt.Println("Done!")
}

func TestReadUntilBlankByte(t *testing.T) {
	fmt.Println("\nTestReadUntilBlankByte:")

	testdata := []byte("1 2 33 444 555\n")
	err := os.WriteFile(fnameReadUntilBlankByteTest, testdata, 0644) //0644: -rw-r--r--
	checkErr(err)
	fmt.Println(fnameReadUntilBlankByteTest, " created for testing readUntilBlankByteTest()!")

	actual := readUntilBlankByte(fnameReadUntilBlankByteTest, 0, 14)
	expected := int64(1)
	if actual != expected {
		t.Errorf("TestReadUntilBlankByte() == %q, want %q", actual, expected)
	}

	actual = readUntilBlankByte(fnameReadUntilBlankByteTest, 1, 14)
	expected = int64(1)
	if actual != expected {
		t.Errorf("TestReadUntilBlankByte() == %q, want %q", actual, expected)
	}

	actual = readUntilBlankByte(fnameReadUntilBlankByteTest, 2, 14)
	expected = int64(3)
	if actual != expected {
		t.Errorf("TestReadUntilBlankByte() == %q, want %q", actual, expected)
	}

	actual = readUntilBlankByte(fnameReadUntilBlankByteTest, 3, 14)
	expected = int64(3)
	if actual != expected {
		t.Errorf("TestReadUntilBlankByte() == %q, want %q", actual, expected)
	}

	actual = readUntilBlankByte(fnameReadUntilBlankByteTest, 4, 14)
	expected = int64(6)
	if actual != expected {
		t.Errorf("TestReadUntilBlankByte() == %q, want %q", actual, expected)
	}

	actual = readUntilBlankByte(fnameReadUntilBlankByteTest, 13, 14)
	expected = int64(13)
	if actual != expected {
		t.Errorf("TestReadUntilBlankByte() == %q, want %q", actual, expected)
	}

	err = os.Remove(fnameReadUntilBlankByteTest) // rm before leaving the test
	checkErr(err)
	if _, err := os.Stat(fnameReadUntilBlankByteTest); !errors.Is(err, os.ErrNotExist) {
		// file exists
		t.Errorf("File %s should not exist after Remove()!", fnameReadUntilBlankByteTest)
	}
	fmt.Println(fnameReadUntilBlankByteTest, " removed successfully after testing!")

	fmt.Println("Done!")
}

func TestSumup(t *testing.T) {
	fmt.Println("\nTestSumup:")

	testdata := []byte("1 2 33 444 555")
	err := os.WriteFile(fnameSumupTest, testdata, 0644) //0644: -rw-r--r--
	checkErr(err)
	fmt.Println(fnameSumupTest, " created for testing sumup()!")

	actual := sumup(fnameSumupTest, 0, 7)
	fmt.Println("Actual: ", actual)

	err = os.Remove(fnameSumupTest) // rm before leaving the test
	checkErr(err)
	if _, err := os.Stat(fnameSumupTest); !errors.Is(err, os.ErrNotExist) {
		// file exists
		t.Errorf("File %s should not exist after Remove()!", fnameSumupTest)
	}
	fmt.Println(fnameSumupTest, " removed successfully after testing!")

	fmt.Println("Done!")
}
