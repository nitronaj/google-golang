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

	fmt.Println(nums)
	sort.Ints(nums)

}

func main() {
	var wg sync.WaitGroup
	var intNumbers []int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numbers := strings.Fields(scanner.Text())

	for _, number := range numbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println(err)
		}

		intNumbers = append(intNumbers, num)
	}

	fmt.Println(intNumbers)
	numberOfPartitions := 4
	lengthOfEachPartition := len(intNumbers) / 4

	subArray := make([][]int, numberOfPartitions)
	fmt.Println(subArray)

	for i := 0; i < len(intNumbers); i = i + lengthOfEachPartition {
		wg.Add(1)
		go sortArray(intNumbers[i:i+lengthOfEachPartition], &wg)
	}

	// // Create four goroutines to sort the array.
	// for i := 0; i < 4; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()

	// 		// Sort the subarray.
	// 		subarray := intNumbers[i*2 : (i+1)*2]
	// 		sort.Ints(subarray)
	// 	}(i)
	// }

	wg.Wait()
	fmt.Println(intNumbers)
}
