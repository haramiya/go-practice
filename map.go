package main

import "fmt"

func main() {
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 300}
	fmt.Println(m2)

	m3 := map[int]string{100: "A", 300: "B"}
	fmt.Println(m3)

	s, ok := m3[400]
	if !ok {
		fmt.Println("error")
	}
	fmt.Println(s, ok)

	delete(m3, 300)
	fmt.Println(m3)

	for k := range m2 {
		fmt.Println(k)
	}
}
