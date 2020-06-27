package main

import "fmt"

func main() {

	nilai := 26

	switch {
	case nilai >= 100:
		fmt.Println("grade: A")
	case (nilai >= 60) && (nilai < 90):
		fmt.Println("grade: B")
	case (nilai >= 30) && (nilai < 50):
		fmt.Println("grade: C")
	default:
		{
			fmt.Println("grade: D")
			fmt.Println("Belajar lagi !!!")
		}
	}

}
