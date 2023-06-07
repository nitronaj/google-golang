package main

import (
	"fmt"
	"strings"
)

func findian(s string) string {
	var lowerString string = strings.ToLower(s)
	if strings.HasPrefix(lowerString, "i") && strings.Contains(lowerString, "a") && strings.HasSuffix(lowerString, "n") {
		return "Found!"
	}

	return "Not Found!"
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(findian(s))
}
