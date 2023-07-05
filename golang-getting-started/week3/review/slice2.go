package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	numbers := make([]int, 3)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a number, and I'll add it to a sorted slice (default [0, 0, 0]): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\r\n")
		if input == "x" {
			fmt.Print("Exiting...")
			break
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Print(err)
			continue
		}
		numbers = append(numbers, number)
		sort.Slice(numbers, func(i, j int) bool {
			return numbers[i] < numbers[j]
		})
		fmt.Println(numbers)
	}
}
