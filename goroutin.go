package main

import (
	"fmt"
	"time"
)

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// sub()    // goが付いていないと同期処理
	go sub() // goが付いていると非同期処理

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}
