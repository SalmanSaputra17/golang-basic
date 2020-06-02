package main

import "fmt"

func main() {

	// tipe data numerik non-desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -123123123

	fmt.Printf("Positif number: %d\n", positiveNumber)
	fmt.Printf("Negatif number: %d\n", negativeNumber)

	// tipe data numerik desimal
	var decimalNumber = 80.216
	fmt.Printf("- %f\n", decimalNumber)
	fmt.Printf("- %.2f\n", decimalNumber)

	// tipe data boolean
	exist := true
	fmt.Printf("Exist? %t\n", exist)

	// tipe data string
	var message = `Nama saya "John Wick". Salam kenal. Mari belajar "Golang".`
	fmt.Println(message)

}
