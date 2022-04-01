package main

import (
	"fmt"
	"proj2_dayuan/lib"
)

func Hello() string {
	return "Hello, world."

}

func main() {
	fmt.Println("This is main.")
	fmt.Println(lib.Hash("hello", uint64(3)))
}
