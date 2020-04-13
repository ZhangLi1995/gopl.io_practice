package main

import "fmt"

// slice 类似于下面的聚合类型，而非纯引用类型
type IntSlice struct {
	ptr      *int
	len, cap int
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// slice 仍有增长空间，扩展 slice 内容
		z = x[:zlen]
	} else {
		// slice 已无空间，为它分配一个新的底层数组
		// 为了达到分摊线性复杂度，容量扩展一倍
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // 内置 copy 函数
	}
	copy(z[len(x):], y)
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	var a []int
	a = append(a, 1)
	a = append(a, 2, 3)
	a = append(a, 4, 5, 6)
	a = append(a, a...)
	fmt.Println(a)
}
