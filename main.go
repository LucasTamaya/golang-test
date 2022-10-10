package main

import "fmt"

func main() {
	// the Println print by going to the line at the end
	fmt.Println("Another print with fmt.Println()")
	fmt.Println("Another print with fmt.Println()")

	// the Print do the same but without going to the line at the end
	fmt.Print("Print with fmt.Print() ")
	fmt.Print("Print with fmt.Print()\n")

	name := "John Doe"

	// the PrintF allow us to print with variables thanks to this
	// format specifier %v will take the name in parameter
	fmt.Printf("Hello %v \n", name)

	// %T: to print the type of variable in parameter
	fmt.Printf("The name variable is of type %T", name)

	// %f: to print float numbers: we can specify decimals if we put
	// 0.1 or 0.2 etc.. => %0.2f = rounded to 2 decimals

	// We can pass multiple variable and multiple format specifier
}
