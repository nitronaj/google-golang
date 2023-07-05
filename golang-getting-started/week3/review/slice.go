package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	//create an empty integer slice of size (length) 3.
	mySlice := make([]int, 3)

	var myInput string
	// loop that runs infinitely until user imputs x or X

	for {

		//Prompts the user to enter an integer to be added to the slice
		fmt.Println("Enter a integer to add to the Slice or X to quit: ")
		fmt.Scan(&myInput)

		// Terminate loop when myIput is x or X
		if myInput == "x" || myInput == "X" {
			break
		}

		//convert input to integer
		myInteger, err := strconv.Atoi(myInput)

		if err != nil {
			fmt.Println("Error while adding integer to Slice")
		} else {

			//append Integer to Slice
			mySlice = append(mySlice, myInteger)

			//Sort the Slice
			sort.Ints(mySlice)

			//Print the sorted Slice
			fmt.Println("Sorted Slice: ", mySlice)

		}
	}
}
