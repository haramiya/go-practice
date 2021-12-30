package main

import "fmt"

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	sl := make([]int, 5)
	fmt.Println(sl)

	sl[0] = 1000
	fmt.Println(sl)

	sl = append(sl, 300)
	fmt.Println(sl)
	fmt.Println(len(sl))
	fmt.Println(cap(sl))

	// sliceの長さ：lenで取得
	// sliceの要領：capで取得
	// 【キャパシティ（容量）について】
	// 要領以上の要素が追加されるとメモリの消費が倍になってしまいます。
	// メモリーを気にするような開発の場合は、容量にも気をつけます。
	// 最初は気にせずやるほうがいいと思います。
	// 過剰にメモリを確保してしまうと実行速度が落ちたりする。
	// 良質なパフォーマンスを実現するには、容量の管理も気にします。
	// sliceのcapacityを増加させるときは、underlying arrayを新しいcapacityで確保し直してポイントを貼り直している。
	// capacityを増やす際の、新しいcapacityの計算方法は下記の通り。

	// 新しいcapacity(仮)を決める。元のcapacityが1024未満なら元のcapacityの2倍、1024以上なら1.25倍する。
	// 1の計算結果から、実際に確保するメモリ容量を計算する。メモリ容量はsliceの要素のtypeやメモリブロックの単位にも依存する。
	// 2で計算したメモリ容量から、最終的なcapacityを計算する。

	sl2 := []int{1, 2, 3, 4, 5}
	sl3 := make([]int, 5, 10) // make([]T, len, cap)

	n := copy(sl3, sl2)

	fmt.Println(sl2, sl3, n)

	sl4 := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(sl4...))
}
