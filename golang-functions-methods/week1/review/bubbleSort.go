package main

import "fmt"

func main() {

	fmt.Println("Starting the application...")

	mySlc := []int{10, 9, 8, 7, 2, 3, 4, 5, 11}

	fmt.Println("The input slice is: ", mySlc)

	bubbleSort(mySlc)

}

func swap(slc []int, i int, j int) {

	temp := slc[i]
	slc[i] = slc[j]
	slc[j] = temp

}

func bubbleSort(slc []int) {

	for i := 0; i < len(slc); i++ {
		for j := i + 1; j < len(slc); j++ {

			if slc[i] > slc[j] {

				swap(slc, i, j)

			}

		}

	}

	fmt.Println("The sorted slice is: ", slc)

}
