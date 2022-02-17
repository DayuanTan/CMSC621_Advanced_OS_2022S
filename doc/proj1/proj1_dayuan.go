package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const randomIntAmnt = 100

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

	intM, err := strconv.Atoi(M)
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

	for i := 0; i < m; i++ {
		startBytePos := partLen * int64(i)
		var endBytePos int64
		if i < m-1 {
			endBytePos = partLen * int64(i+1)
		} else if i == m-1 {
			endBytePos = fByteSize
		}
		assignmenti := Assignment{fname, startBytePos, endBytePos}
		b, err := json.Marshal(assignmenti)
		checkErr(err)

		go worker(b, subsums)
	}

	for i := 0; i < m; i++ {
		fmt.Println(<-subsums)
	}
}

func worker(assignment []byte, subsums chan int64) {
	var assignmenti Assignment
	err := json.Unmarshal(assignment, &assignmenti)
	checkErr(err)
	fmt.Println("worker got assignment: ", assignmenti)
	subsums <- 101
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
