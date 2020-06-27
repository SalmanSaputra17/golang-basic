package main

import (
	"fmt"
	"math"
)

func main() {
	// lingkaran
	var diameter float64 = 7
	var area, circumFerence = calculateRing(diameter)

	fmt.Printf("Luas: %.2f\n", area)
	fmt.Printf("Keliling: %.2f\n", circumFerence)

	// persegi panjang
	panjang := 12
	lebar := 8
	var luas, keliling = calculateSquare(panjang, lebar)

	fmt.Printf("Luas: %d\n", luas)
	fmt.Printf("Keliling: %d\n", keliling)
}

func calculateRing(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	var circumference = math.Pi * d

	// return value
	return area, circumference
}

func calculateSquare(p, l int) (luas int, keliling int) {
	// hitung luas
	luas = p * l
	// hitung keliling
	keliling = 2 * (p + l)

	return
}
