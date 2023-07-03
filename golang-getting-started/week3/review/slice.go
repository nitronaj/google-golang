package main

import (
	"fmt"
	"strconv"
)

func main() {
	integers := make([]int, 3)

	for {
		var input string

		_, err := fmt.Scan(&input)

		// do checks
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		} else if input == "X" {
			fmt.Println("Exiting...")
			break
		}

		// parse the number to integer
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Could not convert to integer, please try again.")
			continue
		}

		integers = appendSorted(integers, number)

		fmt.Println(integers)

	}
}

func appendSorted(integers []int, number int) []int {
	for i, v := range integers {
		if number < v {
			new_integers := make([]int, 0, len(integers)+1)
			new_integers = append(new_integers, integers[:i]...)
			new_integers = append(new_integers, number)
			new_integers = append(new_integers, integers[i:]...)
			return new_integers
		}
	}
	return append(integers, number)
}
