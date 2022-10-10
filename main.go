// this code should be compile to an executable code at the end
// so if we write some utils functions, it will something else than main
package main

import "fmt"

// this function called main is important because it is the entry point of our app
// we should have juste one function with this name in our project
func main() {
	var name string = "Pedro"
	fmt.Println("Hello", name)
}
