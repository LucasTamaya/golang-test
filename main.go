package main

import "fmt"

func main() {
	// we access to the sum function from the sum.go file
	// but we have to run those two files: run go main.go sum.go
	theSum := sum(10, 20)
	fmt.Println(theSum)
}
