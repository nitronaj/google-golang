package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Constants for Slice list and Exit code
const (
	integerSize = 10
	ExitCode    = "X"
)

// Swap function
func Swap(numbers []int, i int) {
	temp := numbers[i]
	numbers[i] = numbers[i+1]
	numbers[i+1] = temp
}

// BubbleSort function
func BubbleSort(numbers []int) {
	Swapped := true
	for Swapped {
		Swapped = false
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				Swap(numbers, i)
				Swapped = true
			}
		}
	}
}

//Regex expressions
func removeSpaces(input string) string {
	regExChar := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	regExNumb := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	result := regExChar.ReplaceAllString(input, "")
	result = regExNumb.ReplaceAllString(result, " ")
	return result
}

func main() {
	intergerScanner := bufio.NewScanner(os.Stdin)
	if err := intergerScanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println("Enter X to quit...")
	fmt.Printf("Enter the list of integers to be sorted: ")

	for intergerScanner.Scan() {
		integerStringValue := intergerScanner.Text()
		if integerStringValue == ExitCode {
			os.Exit(0)
		}

		integerStringValue = removeSpaces(integerStringValue)
		listInteger := strings.Split(integerStringValue, " ")
		sliceInteger := make([]int, 0)
		for _, i := range listInteger {
			j, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println("Invalid input!")
				fmt.Printf("Enter the list of integers to be sorted: ")
				continue
			}
			sliceInteger = append(sliceInteger, j)
		}
		// fmt.Printf("len(sliceInteger): %d; cap(sliceInteger): %d\n", len(sliceInteger), cap(sliceInteger))
		if len(sliceInteger) > integerSize {
			fmt.Printf("Enter a maximum of %d integers\n", integerSize)
			fmt.Printf("Enter the list of integers to be sorted: ")
			continue
		}
		BubbleSort(sliceInteger)
		fmt.Println(sliceInteger)
		fmt.Printf("Enter the list of integers to be sorted: ")
	}
}
