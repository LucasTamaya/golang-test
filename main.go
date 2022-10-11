package main

import "fmt"

func main() {

	// arrays
	// in GO, an array has a defined length that can't be change. It has also the type of values
	// that are going to be inside the array
	ages := [3]int{1, 2, 3}

	names := [3]string{"John", "Koko", "Lolo"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices
	// slices are like arrays in JS or Python. We just have to specify the type value but the length
	// can change

	slices := []int{1, 2, 3, 4, 5}

	// to add element to array
	slices = append(slices, 10)

	// to get the last element of an array
	lastSlices := slices[len(slices)-1]

	fmt.Println(slices)
	fmt.Println(lastSlices)

	// slices ranges

	ranges := slices[0:2]
	fmt.Println(ranges)

}
