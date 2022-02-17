package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const randomIntAmnt = 100

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
	subsums := make(chan int64)

	for i := 0; i < m; i++ {
		go worker("hi", subsums)
	}

	for i := 0; i < m; i++ {
		fmt.Println(<-subsums)
	}
}

func worker(fname string, subsums chan int64) {
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
