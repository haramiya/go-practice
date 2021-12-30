package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i == 100 {
			break
		}
		fmt.Println(i)
	}

	p := 0
	for p < 10 {
		fmt.Println(p)
		p++
	}

	for i2 := 0; i2 < 10; i2++ {
		if i2 == 3 {
			continue
		}
		fmt.Println(i2)
	}

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	sl := []string{"Python", "PHP", "Go"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
