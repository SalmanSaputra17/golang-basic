package main

import "fmt"

func main() {
	/*
	* deklarasi variable menggunakan metode manifest typing
	* yaitu, menuliskan variable sekaligus dengan tipe datanya
	 */
	var firstName string = "Salman"
	var lastName string

	lastName = "Saputra"

	// cetak isi variable
	fmt.Printf("Hello %s %s!\n", firstName, lastName)

	/*
	* deklarasi variable dengan menggunakan metode type inference
	* yaitu, menuliskan variable tidak menggunakan tipe datanya
	* variable seperti ini hanya bisa ditulis didalam fungsi saja
	* penggunaan := hanya saat deklarasi saja
	 */
	strOne := "John"
	strTwo := "Doe"

	fmt.Printf("Halo %s %s\n", strOne, strTwo)

	// deklarasi multi-variable
	one, isTrue, decimal, say := 1, true, 2.2, "Hello World !"

	fmt.Println(one, isTrue, decimal, say)

	/*
	* variable underscore, digunakan untuk menampung value yang tidak akan digunakan
	* isi dari variable underscore tidak dapat ditampilkan
	 */
	_ = "Hai, i learn golang"
}
