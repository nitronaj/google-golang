// Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
// The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3.
// During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
// The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
// The slice must grow in size to accommodate any number of integers which the user decides to enter.
// The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func getNums() []int {
	var input string
	var nums []int = make([]int, 3)
	var count int
	for {
		fmt.Scan(&input)
		if input == "X" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("The input is not number")
			continue
		}

		if count < len(nums) {
			nums[count] = num
		} else {
			nums = append(nums, num)
		}
		count++
	}
	return nums
}

func main() {

	nums := getNums()

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Printf("The list of the numbers: %v \n", nums)
}
