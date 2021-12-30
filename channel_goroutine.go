package main

import (
	"fmt"
	"time"
)

func reciever(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func reciever2(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	/* ch1 := make(chan int)
	ch2 := make(chan int)

	// fmt.Println(<-ch1)

	go reciever(ch1)
	go reciever(ch2)

	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	} */

	ch3 := make(chan int, 2)
	// close(ch3)
	// ch3 <- 1
	// fmt.Println(ch3)

	// i, ok := <-ch3
	// fmt.Println(i, ok)

	go reciever2("1.goroutine", ch3)
	go reciever2("2.goroutine", ch3)
	go reciever2("3.goroutine", ch3)

	i := 0
	for i < 100 {
		ch3 <- i
		i++
	}

	close(ch3)
	time.Sleep(3 * time.Second)
}
