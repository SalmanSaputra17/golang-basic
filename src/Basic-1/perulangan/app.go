package main

import "fmt"

func main() {

	// 1
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("Angka: %d\n", i)
	// }

	// fmt.Println("--------------------------------")

	// 2
	// j := 10
	// for j > 0 {
	// 	fmt.Printf("Angka: %d\n", j)
	// 	j--
	// }

	// fmt.Println("--------------------------------")

	// 3
	// for x := 1; x <= 10; x++ {
	// 	if x%2 == 1 {
	// 		continue
	// 	}

	// 	if x > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Angka: ", x)
	// }

	// fmt.Println("--------------------------------")

	// 4
	// for y := 0; y < 5; y++ {
	// 	for z := y; z < 5; z++ {
	// 		fmt.Print(z, " ")
	// 	}

	// 	fmt.Println()
	// }

	// 5
	for a := 5; a > 0; a-- {
		for b := a; b < 5; b++ {
			fmt.Print("*", " ")
		}

		fmt.Println()
	}

	for i := 5; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print("*", " ")
		}

		fmt.Println()
	}

	fmt.Println("--------------------------------")

	// 6
outerLoop:
	for c := 0; c < 5; c++ {
		for d := 0; d < 5; d++ {
			if c == 3 {
				break outerLoop
			}

			fmt.Print("matriks [", c, "][", d, "]", "\n")
		}
	}

}
