package main

import "fmt"

func main() {

	// maps are like dict in Python or object in JS
	// we specify the type of the key and the type of values
	menu := map[string]int{
		"burger": 40,
		"pizzas": 20,
		"sodas":  10,
	}

	fmt.Println(menu)
	fmt.Println(menu["sodas"])

	// loops throught maps

	for k, v := range menu {
		fmt.Println(k, ":", v)
	}
}
