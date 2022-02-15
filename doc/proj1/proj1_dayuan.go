package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 3 {
		M := os.Args[1]
		fname := os.Args[2]
		fmt.Println("Two arguments are: ", M, " ", fname)
	} else {
		fmt.Println("Please provide two arguments: M and fname!")
	}

	runsth()
}

func runsth() {
	fmt.Println("hello")
}
