package main

import (
	"fmt"
)

// initが先に実行されて、mainが次に実行される
func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("main")
}
