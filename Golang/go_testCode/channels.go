package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func channels() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

}

// Buffered Channels - Here we can specify how many values will a channel hold if it exceeds that then a deadlock occurs
func buffChan() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

//Range and Close

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func rangeMain() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	//Here we don't need waitGroup because range in the following loop waits for the channel to close before terminating
	for i := range c {
		fmt.Println(i)
	}
}
