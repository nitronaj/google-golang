package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Scan(&num)

	fmt.Printf("%.0f\n", math.Trunc(num))
}
