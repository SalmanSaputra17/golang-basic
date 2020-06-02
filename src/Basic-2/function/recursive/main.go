package main

import "fmt"

func main() {
	// defer func() {
	// 	str := recover()
	// 	fmt.Println(str)
	// }()

	// panic("PANIC TEST")

	fmt.Println(factorial(2))

	// fmt.Println(evenOdd(1))
	// fmt.Println(evenOdd(2))
	// fmt.Println(evenOdd(3))
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func evenOdd(number int) (x int, status bool) {
	if number%2 == 0 {
		x, status = 1, true
	} else {
		x, status = 0, false
	}

	return
}
