package main

import (
	"fmt"
)

func anything(a interface{}) {
	switch v := a.(type) {
	case int:
		fmt.Println("is int")
	case string:
		fmt.Println(v + " is string")
	default:
		fmt.Println("I do not know")
	}
}

func main() {
	n := 1
	switch n {
	case 1, 3:
		fmt.Println("1 or 3")
	case 2, 4:
		fmt.Println("2 or 4")
	default:
		fmt.Println("I do not know")
	}

	switch n := 4; n {
	case 1, 3:
		fmt.Println("1 or 3")
	case 2, 4:
		fmt.Println("2 or 4")
	default:
		fmt.Println("I do not know")
	}

	anything("aaa")
	anything(123)

}
