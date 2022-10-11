package main

import (
	"fmt"
	"strings"
)

func sayHello(name string) {
	fmt.Printf("Hello %v", name)
}

// if we want to return, we must specify the return type value, like TS
func sum(a int, b int) int {
	return a + b
}

// return mutliples elements in a function
func getInitials(n string) (string, string) {

	// we use the strings lib to split the names
	splitNames := strings.Split(n, " ")

	var initials []string

	for _, v := range splitNames {
		initials = append(initials, v[:1])
	}

	return initials[0], initials[1]

}

func main() {
	sayHello("John Doe")

	first, second := getInitials("John Doe")

	fmt.Println(first, second)
}
