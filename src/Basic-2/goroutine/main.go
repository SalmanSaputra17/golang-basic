package main

import(
	"fmt"
	"time"
	"math/rand"
)

func loopNum(n int) {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Loop ke - %d.%d\n", n, i)
		interval := time.Duration(rand.Intn(100))
		time.Sleep(time.Millisecond * interval)
	}
}

func main() {
	for i := 1; i <= 3; i++ {
		go loopNum(i)
	}

	var yourInput string
	fmt.Scanln(&yourInput)
}