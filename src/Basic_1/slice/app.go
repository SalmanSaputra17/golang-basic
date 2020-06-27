package main

import "fmt"

func printLine() {
	fmt.Println("-----------------------------------------")
}

func main() {

	buahPedagang := []string{
		"Melon",
		"Semangka",
		"Mangga",
		"Jeruk",
		"Pisang",
	}

	keranjangPembeli := buahPedagang[0:2]

	printLine()

	fmt.Printf("Dijual %d jenis buah, sebagai berikut:\n", len(buahPedagang))
	for i, row := range buahPedagang {
		fmt.Printf("%d. %s\n", i+1, row)
	}

	printLine()

	fmt.Printf("Seorang pembeli membeli %d jenis buah, yaitu sebagai berikut:\n", len(keranjangPembeli))
	for i, row := range keranjangPembeli {
		fmt.Printf("%d. %s\n", i+1, row)
	}

	printLine()

	tambahBuah := append(buahPedagang, "Pepaya")
	fmt.Println("Pedagang menambah jenis buah dagangannya, Sekarang buah pedagang adalah:")
	for i, row := range tambahBuah {
		fmt.Printf("%d. %s\n", i+1, row)
	}

}
