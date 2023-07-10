package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// Write a program to sort an array of integers. The program should partition the array into 4 parts,
// each of which is sorted by a different goroutine. Each partition should be of approximately equal size.
// Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

// The program should prompt the user to input a series of integers.
// Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
// When sorting is complete, the main goroutine should print the entire sorted list

func sortArray(nums []int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("---", nums)
	sort.Ints(nums)
	fmt.Println("Sorted slice:", nums)
}

func readNumbers() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.Fields(scanner.Text())
}

func covertToSliceOfIns(numbers []string) []int {
	var intNumbers []int
	for _, number := range numbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println(err)
		}

		intNumbers = append(intNumbers, num)
	}

	return intNumbers
}

func main() {
	var wg sync.WaitGroup
	var intNumbers []int // []int{1, 5, 9, 3, 7, 4, 8, 6, 2, 0, 25, -1}

	numbers := readNumbers()
	intNumbers = covertToSliceOfIns(numbers)

	numberOfPartitions := 4
	lengthOfEachPartition := len(intNumbers) / numberOfPartitions

	for i := 0; i < len(intNumbers); i = i + lengthOfEachPartition {
		wg.Add(1)
		go sortArray(intNumbers[i:i+lengthOfEachPartition], &wg)
	}

	wg.Wait()
	sort.Ints(intNumbers)
	fmt.Println("Sorted Final Array:", intNumbers)
}
