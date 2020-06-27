package main

import "fmt"

func main() {
	x := 1.5
	square(&x)
	// fmt.Println(x)

	a, b := 1, 2
	// swap(&a, &b)
	swapWithoutTemp(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

func square(x *float64) {
	*x = *x * *x
}

// swap variable value
func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

// swap variable value without temporary variable
func swapWithoutTemp(x *int, y *int) {
	*x = *x + *y
	*y = *x - *y
	*x = *x - *y
}
