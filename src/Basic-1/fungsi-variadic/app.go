package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers = []int{2, 432, 5, 7, 2, 4342, 532, 9}
	// avg := calculate(2, 432, 5, 7, 2, 4342, 532, 9)
	avg := calculate(numbers...)
	msg := fmt.Sprintf("Rata-rata: %.2f", avg)
	fmt.Println(msg)

	myHobbies("Salman Saputra", "Membaca", "Futsal", "Badminton", "etc")
}

/*
* fungsi variadic adalah fungsi yang bisa memiliki parameter tak terbatas
* dengan tipe data yang sama
 */
func calculate(numbers ...int) float64 {
	var total int
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

func myHobbies(name string, hobbies ...string) {
	var strHobbies = strings.Join(hobbies, ", ")
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("My Hobbies are %s\n", strHobbies)
}

// ----------------------------------------------------

// func printAvg(numbers []int) float64 {
// 	var total int = 0
// 	for _, number := range numbers {
// 		total += number
// 	}

// 	var avg = float64(total) / float64(len(numbers))
// 	return avg
// }
