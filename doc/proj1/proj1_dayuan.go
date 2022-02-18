package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const randomIntAmnt = 100 // how many random int numbers in input_data_file.txt

type Assignment struct {
	Datafile     string
	StartBytePos int64
	EndBytePos   int64
}

func main() {
	var M, fname string
	if len(os.Args) == 3 {
		M = os.Args[1]
		fname = os.Args[2]
		fmt.Println("Two arguments are: ", M, " ", fname)
	} else {
		fmt.Println("Please provide two arguments: M and fname!")
		panic("Main func argument maloperation!")
	}

	geneRandomInt(fname) // generate randome 100 int numbers and store into input_data_file.txt.

	intM, err := strconv.Atoi(M) // string to int
	checkErr(err)

	concurrencySum(intM, fname)
}

func concurrencySum(m int, fname string) {
	fi, err := os.Stat(fname)
	checkErr(err)
	fByteSize := fi.Size()
	partLen := fByteSize / int64(m)
	fmt.Printf("The file %s is %d bytes long. It is partitioned into %d parts.\n", fname, fByteSize, m)

	subsums := make(chan int64)

	startBytePos := int64(0)
	endBytePos := int64(0)
	for i := 0; i < m; i++ {
		// [startBytePos, endBytePos), left included, right excluded
		startBytePos = endBytePos
		if i < m-1 {
			tempEndBytePos := partLen * int64(i+1)
			endBytePos = readUntilBlankByte(fname, tempEndBytePos, fByteSize)
		} else if i == m-1 {
			endBytePos = fByteSize
		}
		assignmenti := Assignment{fname, startBytePos, endBytePos}
		b, err := json.Marshal(assignmenti)
		checkErr(err)

		go worker(b, subsums, fByteSize)
	}

	for i := 0; i < m; i++ {
		fmt.Println(<-subsums)
	}
}

func worker(assignment []byte, subsums chan int64, fByteSize int64) {
	var assignmenti Assignment
	err := json.Unmarshal(assignment, &assignmenti)
	checkErr(err)
	fmt.Println("worker got assignment: ", assignmenti)
	subsums <- sumup(assignmenti.Datafile, assignmenti.StartBytePos, assignmenti.EndBytePos, fByteSize)
}

func readUntilBlankByte(fname string, tempEndBytePos int64, fByteSize int64) int64 {
	f, err := os.Open(fname)
	checkErr(err)
	defer f.Close()

	if tempEndBytePos < fByteSize-1 { // if == fByteSize-1 then it is last pos then no action needed
		for {
			_, err := f.Seek(tempEndBytePos, 0) // locate at endByte
			checkErr(err)
			theByteAfterEndByte := make([]byte, 1)
			_, err = f.Read(theByteAfterEndByte) // read 1 byte to check whether it is blank byte
			checkErr(err)

			if string(theByteAfterEndByte) == " " {
				break
			} else {
				tempEndBytePos++
			}
		}
	}
	return tempEndBytePos
}

func sumup(fname string, startBytePos int64, endBytePos int64, fByteSize int64) int64 { // the input will gurantee that don't cut a number, by concurrencySum
	f, err := os.Open(fname)
	checkErr(err)
	defer f.Close()

	if endBytePos > fByteSize {
		endBytePos = fByteSize
	}
	dataBytes := make([]byte, endBytePos-startBytePos)
	_, err = f.Seek(startBytePos, 0) // go to the position startBytePos
	checkErr(err)
	actualRead, err := f.Read(dataBytes) // read bytes with length=dataBytes from the position startBytePos
	checkErr(err)
	fmt.Printf("Actual read %d bytes: '%s'\n", actualRead, string(dataBytes[:]))

	// convert []byte to strings to []string to []int64
	words := strings.Fields(string(dataBytes))
	sum := int64(0)
	for _, v := range words {
		vInt64, err := strconv.ParseInt(v, 10, 64) // string to int64
		checkErr(err)
		sum += vInt64
	}
	fmt.Println("sum: ", sum)

	return sum
}

func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func geneRandomInt(fname string) {
	sou := rand.NewSource(time.Now().UnixNano())
	ran := rand.New(sou)

	f, err := os.Create(fname)
	checkErr(err)
	defer f.Close()

	for i := 0; i < randomIntAmnt; i++ {
		ranInt := ran.Intn(10000)
		// fmt.Println(ranInt)
		ranInt64Str := strconv.FormatInt(int64(ranInt), 10)
		_, err = f.WriteString(ranInt64Str)
		checkErr(err)
		_, err = f.WriteString(" ")
		checkErr(err)
	}
	_, err = f.WriteString("\n")
	checkErr(err)
	fmt.Println(fname, " has been generated and 100 random int has been stored.")
}
