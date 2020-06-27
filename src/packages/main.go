package main

import (
	"fmt"
	"packages/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	max := math.Max(xs)
	min := math.Min(xs)

	fmt.Printf("Average: %.2f \n", avg);
	fmt.Printf("Max: %.2f \n", max);
	fmt.Printf("Min: %.2f \n", min);
}