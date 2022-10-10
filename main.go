package main

import "fmt"

func main() {
	// here we specify explicitly the type of the variable
	var name string = "Pedro"
	// but GO can do it by its own
	var nameTwo = "John"

	// the fast way to init a var is like so, but we can use this method only in function
	nameThree := "Guigui"

	fmt.Println(name, nameTwo, nameThree)

	// like Python, GO has int and float numbers

	// int number
	intNumber := 15

	// float number
	floatNumber := 10.2

	fmt.Println(intNumber, floatNumber)
}
