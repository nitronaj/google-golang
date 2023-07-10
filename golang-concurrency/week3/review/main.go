package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a sequence of integers separated by a space:")
	userInput := readInput()
	numbers := toNumbersSlice(userInput)

	sliceSize := len(numbers) / 4
	firstPosition := 0
	channel := make(chan []int)
	go sortToChannel("Routine1: ", numbers[firstPosition:sliceSize], channel)
	firstPosition += sliceSize
	go sortToChannel("Routine2: ", numbers[firstPosition:firstPosition+sliceSize], channel)
	firstPosition += sliceSize
	go sortToChannel("Routine3: ", numbers[firstPosition:firstPosition+sliceSize], channel)
	firstPosition += sliceSize
	go sortToChannel("Routine4: ", numbers[firstPosition:], channel)

	sorted1 := <-channel
	sorted2 := <-channel
	sorted3 := <-channel
	sorted4 := <-channel

	result := make([]int, 0)
	result = append(result, sorted1...)
	result = append(result, sorted2...)
	result = append(result, sorted3...)
	result = append(result, sorted4...)

	sort("Main: ", result)
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading the string: ", err)
		os.Exit(1)
	}
	return strings.TrimSuffix(value, "\n")
}

func toNumbersSlice(input string) []int {
	splitIntput := strings.Fields(input)
	numbers := make([]int, len(splitIntput))
	var err error
	for n := range splitIntput {
		numbers[n], err = strconv.Atoi(splitIntput[n])
		if err != nil {
			fmt.Println("Error converting to int: ", err)
			os.Exit(1)
		}
	}
	return numbers
}

func sortToChannel(routineLabel string, numbers []int, channel chan []int) {
	sort(routineLabel, numbers)
	channel <- numbers
}

func sort(routineLabel string, numbers []int) {
	fmt.Println(routineLabel, numbers)
	for i := range numbers {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j+1] < numbers[j] {
				Swap(numbers, j)
			}
		}
	}
	fmt.Println(routineLabel, numbers)
}

func Swap(numbers []int, pos int) {
	numbers[pos], numbers[pos+1] = numbers[pos+1], numbers[pos]
}
