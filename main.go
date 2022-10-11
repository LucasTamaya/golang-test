package main

import (
	"fmt"
)

func sayHello(name string) {
	fmt.Printf("Hello %v", name)
}

// if we want to return, we must specify the return type value, like TS
func sum(a int, b int) int {
	return a + b
}

func main() {
	sayHello("John Doe")
}
