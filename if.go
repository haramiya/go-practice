package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 0
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I do not know")
	}

	// エラーハンドリング
	var s string = "A"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)
}
