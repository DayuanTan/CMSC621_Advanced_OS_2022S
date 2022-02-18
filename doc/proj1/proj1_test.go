package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"
)

const fnameTest = "input_data_file_test.txt"
const fnamereadUntilWhitespaceByteTest = "test_data_for_readUntilWhitespaceByte.test"
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

func TestreadUntilWhitespaceByte(t *testing.T) {
	fmt.Println("\nTestreadUntilWhitespaceByte:")

	testdata := []byte("1 2 33 444 555\n")
	err := os.WriteFile(fnamereadUntilWhitespaceByteTest, testdata, 0644) //0644: -rw-r--r--
	checkErr(err)
	fmt.Println(fnamereadUntilWhitespaceByteTest, " created for testing readUntilWhitespaceByteTest()!")

	actual := readUntilWhitespaceByte(fnamereadUntilWhitespaceByteTest, 0, 14)
	expected := int64(1)
	if actual != expected {
		t.Errorf("TestreadUntilWhitespaceByte() == %q, want %q", actual, expected)
	}

	actual = readUntilWhitespaceByte(fnamereadUntilWhitespaceByteTest, 1, 14)
	expected = int64(1)
	if actual != expected {
		t.Errorf("TestreadUntilWhitespaceByte() == %q, want %q", actual, expected)
	}

	actual = readUntilWhitespaceByte(fnamereadUntilWhitespaceByteTest, 2, 14)
	expected = int64(3)
	if actual != expected {
		t.Errorf("TestreadUntilWhitespaceByte() == %q, want %q", actual, expected)
	}

	actual = readUntilWhitespaceByte(fnamereadUntilWhitespaceByteTest, 3, 14)
	expected = int64(3)
	if actual != expected {
		t.Errorf("TestreadUntilWhitespaceByte() == %q, want %q", actual, expected)
	}

	actual = readUntilWhitespaceByte(fnamereadUntilWhitespaceByteTest, 4, 14)
	expected = int64(6)
	if actual != expected {
		t.Errorf("TestreadUntilWhitespaceByte() == %q, want %q", actual, expected)
	}

	actual = readUntilWhitespaceByte(fnamereadUntilWhitespaceByteTest, 13, 14)
	expected = int64(13)
	if actual != expected {
		t.Errorf("TestreadUntilWhitespaceByte() == %q, want %q", actual, expected)
	}

	err = os.Remove(fnamereadUntilWhitespaceByteTest) // rm before leaving the test
	checkErr(err)
	if _, err := os.Stat(fnamereadUntilWhitespaceByteTest); !errors.Is(err, os.ErrNotExist) {
		// file exists
		t.Errorf("File %s should not exist after Remove()!", fnamereadUntilWhitespaceByteTest)
	}
	fmt.Println(fnamereadUntilWhitespaceByteTest, " removed successfully after testing!")

	fmt.Println("Done!")
}

func TestSumup(t *testing.T) {
	fmt.Println("\nTestSumup:")

	testdata := []byte("1 2 33 444 555")
	err := os.WriteFile(fnameSumupTest, testdata, 0644) //0644: -rw-r--r--
	checkErr(err)
	fmt.Println(fnameSumupTest, " created for testing sumup()!")

	// the input of sumup cannot cut a number, e.g. sumup(fnameSumupTest, 1, 6) cuts the 33
	// this is guranteed in concurrencySum already
	actual := sumup(fnameSumupTest, 0, 7, 14) // [0,7) = "1 2 33 "
	expected, err := json.Marshal(SubSumResult{int64(2), int64(1), int64(1), int64(33), int64(0), int64(7)})
	checkErr(err)
	if !bytes.Equal(actual, expected) {
		t.Errorf("TestSumup() == %q, want %q", actual, expected)
	}

	actual = sumup(fnameSumupTest, 1, 7, 14) // [1,7) = " 2 33 "
	expected, err = json.Marshal(SubSumResult{int64(0), int64(0), int64(2), int64(33), int64(1), int64(7)})
	checkErr(err)
	if !bytes.Equal(actual, expected) {
		t.Errorf("TestSumup() == %q, want %q", actual, expected)
	}

	actual = sumup(fnameSumupTest, 1, 10, 14) // [1,10) = " 2 33 444"
	expected, err = json.Marshal(SubSumResult{int64(33), int64(1), int64(2), int64(444), int64(1), int64(10)})
	checkErr(err)
	if !bytes.Equal(actual, expected) {
		t.Errorf("TestSumup() == %q, want %q", actual, expected)
	}

	actual = sumup(fnameSumupTest, 1, 11, 14) // [1,11) = " 2 33 444 ""
	expected, err = json.Marshal(SubSumResult{int64(33), int64(1), int64(2), int64(444), int64(1), int64(11)})
	checkErr(err)
	if !bytes.Equal(actual, expected) {
		t.Errorf("TestSumup() == %q, want %q", actual, expected)
	}

	actual = sumup(fnameSumupTest, 2, 11, 14) // [2,11) = "2 33 444 "
	expected, err = json.Marshal(SubSumResult{int64(33), int64(1), int64(2), int64(444), int64(2), int64(11)})
	checkErr(err)
	if !bytes.Equal(actual, expected) {
		t.Errorf("TestSumup() == %q, want %q", actual, expected)
	}

	actual = sumup(fnameSumupTest, 3, 11, 14) // [3,11) = " 33 444 "
	expected, err = json.Marshal(SubSumResult{int64(0), int64(0), int64(33), int64(444), int64(3), int64(11)})
	checkErr(err)
	if !bytes.Equal(actual, expected) {
		t.Errorf("TestSumup() == %q, want %q", actual, expected)
	}

	actual = sumup(fnameSumupTest, 7, 14, 14) // [7,14) = "444 555"
	expected, err = json.Marshal(SubSumResult{int64(0), int64(0), int64(444), int64(555), int64(7), int64(14)})
	checkErr(err)
	if !bytes.Equal(actual, expected) {
		t.Errorf("TestSumup() == %q, want %q", actual, expected)
	}

	actual = sumup(fnameSumupTest, 7, 15, 14) // [7,15) is out of range so it is truncated to [7,14)="444 555"
	expected, err = json.Marshal(SubSumResult{int64(0), int64(0), int64(444), int64(555), int64(7), int64(14)})
	checkErr(err)
	if !bytes.Equal(actual, expected) {
		t.Errorf("TestSumup() == %q, want %q", actual, expected)
	}

	err = os.Remove(fnameSumupTest) // rm before leaving the test
	checkErr(err)
	if _, err := os.Stat(fnameSumupTest); !errors.Is(err, os.ErrNotExist) {
		// file exists
		t.Errorf("File %s should not exist after Remove()!", fnameSumupTest)
	}
	fmt.Println(fnameSumupTest, " removed successfully after testing!")

	fmt.Println("Done!")
}
