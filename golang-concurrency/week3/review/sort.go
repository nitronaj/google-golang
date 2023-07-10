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

func main() {
	fmt.Println("Enter the sequence of integers with space in between")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	seq := strings.Split(input, " ")
	sli := make([]int, 0)
	sliTemp := make([]int, 0)
	for _, v := range seq {
		temp, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		sli = append(sli, temp)
	}

	var a int
	arrA := make([]int, 0)
	arrB := make([]int, 0)
	arrC := make([]int, 0)
	arrD := make([]int, 0)
	a = len(sli)/4
	for i, v := range sli {
		if i < a {
			arrA = append(arrA, v)
			continue
		}
		if i < 2*a {
			arrB = append(arrB, v)
			continue
		}
		if i < 3*a {
			arrC = append(arrC, v)
			continue
		}
		if i < 4*a {
			arrD = append(arrD, v)
			continue
		} else {
			sliTemp = append(sliTemp, v)
		}
	}
	for i, v := range sliTemp {
		if i == 0 {
			arrA = append(arrA, v)
		}
		if i == 1 {
			arrB = append(arrB, v)
		}
		if i == 2 {
			arrC = append(arrC, v)
		}
	}

	var wg sync.WaitGroup
	wg.Add(4)
	go sortSec(arrA, &wg)
	go sortSec(arrB, &wg)
	go sortSec(arrC, &wg)
	go sortSec(arrD, &wg)
	wg.Wait()

	tempArray := [][]int{arrA, arrB, arrC, arrD}
	mergedArray := make([]int, 0)
	for _,v := range tempArray {
		mergedArray = append(mergedArray, v...)
	}
	sort.Ints(mergedArray)
	fmt.Printf("final merged and sorted array: %v", mergedArray)
}

func sortSec(arr []int, wg *sync.WaitGroup) {
	fmt.Printf("array to be sorted: %v\n", arr)
	sort.Ints(arr)
	fmt.Printf("array sorted: %v\n", arr)
	wg.Done()
}