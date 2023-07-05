// Write a Bubble Sort program in Go. The program
// should prompt the user to type in a sequence of up to 10 integers. The program
// should print the integers out on one line, in sorted order, from least to
// greatest. Use your favorite search tool to find a description of how the bubble
// sort algorithm works.

// As part of this program, you should write a
// function called BubbleSort() which
// takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
// order.

// A recurring operation in the bubble sort algorithm is
// the Swap operation which swaps the position of two adjacent elements in the
// slice. You should write a Swap() function which performs this operation. Your Swap()
// function should take two arguments, a slice of integers and an index value i which
// indicates a position in the slice. The Swap() function should return nothing, but it should swap
// the contents of the slice in position i with the contents in position i+1.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(numbers []int, index int) {
	numbers[index+1], numbers[index] = numbers[index], numbers[index+1]
}

func BubbleSort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}
		}
	}
}

func main() {
	var intNumbers []int
	var numberOfIntegers = 9

	// Create a scanner to read from stdin.
	scanner := bufio.NewScanner(os.Stdin)

	// Read the string of space separated integers from stdin.
	scanner.Scan()
	numbers := strings.Fields(scanner.Text())
	if len(numbers) < numberOfIntegers {
		numberOfIntegers = len(numbers)
	}

	numbers = numbers[:numberOfIntegers]

	for _, number := range numbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println(err)
		}

		intNumbers = append(intNumbers, num)
	}

	fmt.Println(intNumbers)
	BubbleSort(intNumbers)
	fmt.Println(intNumbers)

}
