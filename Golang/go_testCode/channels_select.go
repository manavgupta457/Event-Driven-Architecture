package main

import "fmt"

func fibonacci_select(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default: //Will run while channel c and quit are not receiving any values
			fmt.Println("*")
		}
	}

}

func selectMain() {
	c := make(chan int)

	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci_select(c, quit)
}
