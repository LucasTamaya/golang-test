package main

import "fmt"

func main() {

	myBill := newBill("John Doe's bill")

	fmt.Println(myBill)

	myBill.format()

}
