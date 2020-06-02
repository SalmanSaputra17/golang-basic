package main

import(
	"fmt"
	"time"
)

// func Sleep(x int) {
// 	<- time.After(time.Second * time.Duration(x))
// }

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from channel 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from channel 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <- c1:
				fmt.Println("message 1:", msg1)
			case msg2 := <- c2:
				fmt.Println("message 2:", msg2)
			case <- time.After(time.Second):
				fmt.Println("timeout")
				time.Sleep(time.Second * 1)
			default:
				fmt.Println("nothing ready in channel")
				time.Sleep(time.Second * 1)
			}
		}
	}()

	var input string
  	fmt.Scanln(&input)
}