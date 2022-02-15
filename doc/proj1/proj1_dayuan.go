package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	var M, fname string
	if len(os.Args) == 3 {
		M = os.Args[1]
		fname = os.Args[2]
		fmt.Println("Two arguments are: ", M, " ", fname)
	} else {
		fmt.Println("Please provide two arguments: M and fname!")
	}

	geneRandomInt(fname) // generate randome 100 int numbers and store into input_data_file.txt.
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

	for i := 0; i < 100; i++ {
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
}
