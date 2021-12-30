package main

import "fmt"

func Plus(x int, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 関数を返す関数
func ReturnFunc() func() {
	return func() {
		fmt.Println("I am a function")
	}
}

// 関数を引数にとる関数
func CallFunction(f func()) {
	f()
}

func main() {
	i := Plus(1, 10)
	fmt.Println(i)

	i2, i3 := Div(9, 4)
	i4, _ := Div(10, 4)
	fmt.Println(i2, i3, i4)

	// 無名関数
	f := func(x, y int) int {
		return x * y
	}
	i5 := f(3, 4)
	fmt.Println(i5)

	i6 := func(x, y int) int {
		return x * y
	}(4, 5)
	fmt.Println(i6)

	// 関数を返す関数
	f2 := ReturnFunc()
	f2()

	// 関数を引数にとる関数
	CallFunction((func() {
		fmt.Println("Call Function")
	}))

	// クロージャー

	// Goの無名関数はクロージャーで、

	// クロージャーとは日本語では関数閉包と呼ばれ、関数と関数の処理に関する関数外の環境をセットして閉じ込めた物です。

	// ジェネレーター

	// 何らかのルールに従って連続した値を返し続ける仕組みの事。
}
