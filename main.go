package main

import (
	"fmt"
)

func updateNameByMemoryAddress(v *string) {
	*v = "wedge"
}

func main() {
	name := "John"

	// $var allow us to access to the memory address of variables
	fmt.Println("The memory address of name is", &name)

	nameMemoryValue := &name
	// *&name allow us to acces to the value of a memory address

	updateNameByMemoryAddress(nameMemoryValue)

	fmt.Println(name)
}
