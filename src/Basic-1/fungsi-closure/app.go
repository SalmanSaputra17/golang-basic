package main

import (
	"fmt"
)

func main() {

	// closure
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}

		return min, max
	}

	var numbers = []int{2, 4, 1, 7, 3, 62, 2, 5, 7}
	var min, max = getMinMax(numbers)

	fmt.Printf("data: %v\nmin: %v\nmax: %v\n", numbers, min, max)

	// closure IIFE
	var newNumber = func(min int) (arr []int) {
		for _, i := range numbers {
			if i < min {
				continue
			} else {
				arr = append(arr, i)
			}
		}
		return
	}(2)

	fmt.Printf("Angka Awal: %v\n", numbers)
	fmt.Printf("Angka Baru: %v\n", newNumber)

	var maxNumber = 3
	var numberList = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers = findLower(numberList, maxNumber)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numberList)
	fmt.Printf("find \t: %d\n\n", maxNumber)

	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]
}

// closure sebagai nilai kembalian
func findLower(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		return res
	}
}
