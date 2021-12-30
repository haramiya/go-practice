package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World")

	// byte型（uint8）型
	byteA := []byte{72, 73}
	fmt.Println(byteA)
	fmt.Println(string(byteA))

	// 配列型
	arr1 := [4]int{1, 2, 3}
	arr2 := []int{1, 2, 3}
	arr3 := [...]string{"test", "test", "test"} // [...]で要素数を自動計算
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr1[0])
	// 他のプログラミング言語の配列型になれていると、若干奇妙かもしれませんが、Goの配列型は要素数を変更できません。

	// 要素の追加などを行う場合は、後のレクチャーで登場するスライス型を使います。

	// 配列型　＝　要素数を変更できない。
	// スライス型　＝　要素数を変更可能。

	// interface型（すべての型を汎用的に表す）
	var x interface{}
	fmt.Println(x) // 初期値はnil
	x = 1
	fmt.Println(x)

	x = [3]int{1, 2, 3}
	fmt.Println(x)

	s := "100"
	i, _ := strconv.Atoi(s) // _を使うことで2つ目の値を破棄できる
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

}
