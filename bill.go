package main

import (
	"fmt"
)

// we create the bill type
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

var itemsData = map[string]float64{"burger": 10.90, "pizza": 9.90}

func newBill(name string) bill {

	return bill{
		name:  name,
		items: itemsData,
		tip:   0,
	}
}

// Receiver Function
// this creates a function that is only accessible on bill objects
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64

	for k, v := range b.items {
		total += v
		fs += fmt.Sprintln(k, ":", v, "$")
	}

	fs += fmt.Sprintln("Total: ", total, "$")

	fmt.Print(fs)

	return fs
}
