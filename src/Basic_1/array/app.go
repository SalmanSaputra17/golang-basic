package main

import "fmt"

func main() {
	// deklarasi array
	var fruits = [4]string{
		"Melon",
		"Jeruk",
		"Apel",
		"Semangka",
	}

	fmt.Printf("Jumlah elemen: %d\n", len(fruits))
	for i, fruit := range fruits {
		fmt.Printf("Elemen ke-%d: %s\n", i, fruit)
	}

	// perulangan array dengan menggunakan underscore variable
	var names = [...]string{
		"Salman",
		"Saputra",
	}

	fmt.Println("Daftar Nama: ")
	for _, name := range names {
		fmt.Printf(name + "\n")
	}

	// menggunakan keyword make
	var arrInt = make([]int, 3)
	arrInt[0] = 12
	arrInt[1] = 126
	arrInt[2] = 9017

	for _, row := range arrInt {
		fmt.Println(row)
	}

	// array multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// looping terhadap array multidimensi
	for _, row := range numbers1 {
		for _, row2 := range row {
			fmt.Printf("%d ", row2)
		}
	}
}
