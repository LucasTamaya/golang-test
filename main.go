package main

import "fmt"

func main() {

	compteur := 0

	for compteur <= 10 {
		fmt.Println(compteur)

		if compteur == 10 {
			fmt.Println("BOOM")
			compteur++
			break
		}

		compteur++
	}

	names := []string{"Jojo", "Lolo", "Koko"}
	ages := []int{10, 15, 20}

	// loop throught an array with index and value
	for index, value := range names {
		fmt.Printf("The index is %v, and the value is %v \n", index, value)
	}

	// loop just throught value
	for _, value := range ages {
		fmt.Printf("The age is %v \n", value)
	}

}
