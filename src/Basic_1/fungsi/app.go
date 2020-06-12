package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var me = []string{"Salman", "Saputra"}
	sayHello("Good Morning", me)

	rand.Seed(time.Now().Unix())

	var randomValue int
	randomValue = randomWithRange(719, 112)
	fmt.Println("Random Number:", randomValue)

	divideNumber(10, 5)
}

func sayHello(message string, name []string) {
	var nameString = strings.Join(name, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

func divideNumber(x, y int) {
	if y == 0 {
		fmt.Printf("Invalid Divider, %d cannot divide by %d\n", x, y)
		return
	}

	result := x / y
	fmt.Printf("%d / %d = %d\n", x, y, result)
}
