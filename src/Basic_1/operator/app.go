package main

import "fmt"

func main() {
	// operator aritmatika
	var value = (((2+6)%3)*4 - 2) / 3
	fmt.Printf("Result: %d\n", value)

	// operator perbandingan
	var result = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", result, isEqual)

	// operator logika
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}
